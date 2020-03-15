# Web 服务代码实现

调用 `Http.HandleFunc`，按顺序做了几件事：

1. 调用了 DefaultServeMux 的 HandleFunc
2. 调用了 DefaultServeMux 的 Handle
3. 往 DefaultServeMux 的 map[string]muxEntry 中增加对应的 handler 和路由规则

调用 `http.ListenAndServe(":9090", nil)`，按顺序做了几件事情：

1. 实例化 Server
2. 调用 Server 的 ListenAndServe()
3. 调用 net.Listen(“tcp”, addr)监听端口
4. 启动一个 for 循环，在循环体中 Accept 请求
5. 对每个请求实例化一个 Conn，并且开启一个 goroutine 为这个请求进行服务 go c.serve()
6. 读取每个请求的内容 w, err := c.readRequest()
7. 判断 handler 是否为空，如果没有设置 handler（这个例子就没有设置 handler），handler 就设置为 DefaultServeMux
8. 调用 handler 的 ServeHttp
9. 在这个例子中，下面就进入到 DefaultServeMux.ServeHttp
10. 根据 request 选择 handler，并且进入到这个 handler 的 ServeHTTP mux.handler(r).ServeHTTP(w, r)
11. 选择 handler：
    A 判断是否有路由能满足这个 request（循环遍历 ServerMux 的 muxEntry）
    B 如果有路由满足，调用这个路由 handler 的 ServeHttp
    C 如果没有路由满足，调用 NotFoundHandler 的 ServeHttp

# 路由注册代码

调用 `http.HandleFunc(“/”, sayhelloName)` 注册路由：

```go
// /usr/local/go/src/net/http/server.go:2081
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    DefaultServeMux.HandleFunc(pattern, handler) // DefaultServeMux 类型为 *ServeMux
}
```

使用默认 ServeMux：

```go
// :2027
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    mux.Handle(pattern, HandlerFunc(handler))
}
```

注册路由策略 DefaultServeMux：

```go
func (mux *ServeMux) Handle(pattern string, handler Handler) {
    mux.mu.Lock()
    defer mux.mu.Unlock()

    ...

    mux.m[pattern] = muxEntry{explicit: true, h: handler, pattern: pattern}

    if pattern[0] != '/' {
        mux.hosts = true
    }

    ...
}

```

涉及数据结构：

```go
// :1900 ServeMux 默认实例是 DefaultServeMux
type ServeMux struct {
    mu    sync.RWMutex // 锁，由于请求涉及到并发处理，因此这里需要一个锁机制
    m     map[string]muxEntry // 路由规则，一个string对应一个mux实体，这里的string就是注册的路由表达式
    hosts bool // 是否在任意的规则中带有host信息
}

type muxEntry struct {
    explicit bool
    h        Handler // 路由处理器
    pattern  string  // url 匹配正则
}

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

# 服务监听代码

调用 `err := http.ListenAndServe(“:9090”, nil)` 监听端口：

```go
// /usr/local/go/src/net/http/server.go:2349
func ListenAndServe(addr string, handler Handler) error {
    server := &Server{Addr: addr, Handler: handler} // handler 为空
    return server.ListenAndServe()
}

```

创建一个 Server 对象，并调用 Server 的 ListenAndServe()。然后监听 TCP 端口：

```go
// :2210
func (srv *Server) ListenAndServe() error {
    addr := srv.Addr
    if addr == "" {
        addr = ":http"
    }
    ln, err := net.Listen("tcp", addr)
    if err != nil {
        return err
    }
    return srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
}
```

## 接收请求

```go
// :2256
func (srv *Server) Serve(l net.Listener) error {
    defer l.Close()

    ...

    baseCtx := context.Background()
    ctx := context.WithValue(baseCtx, ServerContextKey, srv)
    ctx = context.WithValue(ctx, LocalAddrContextKey, l.Addr())
    for {
        rw, e := l.Accept() // 1. Listener 接收请求
        if e != nil {
            ...
        }
        tempDelay = 0
        c := srv.newConn(rw) // 2. 创建 *conn
        c.setState(c.rwc, StateNew) // before Serve can return
        go c.serve(ctx) // 3. 新启一个goroutine，将请求数据做为参数传给 conn，由这个新的goroutine 来处理这次请求
    }
}

```

goroutine 处理请求：

```go
// Serve a new connection.
func (c *conn) serve(ctx context.Context) {
    ...
    // HTTP/1.x from here on.

    c.r = &connReader{r: c.rwc}
    c.bufr = newBufioReader(c.r)
    c.bufw = newBufioWriterSize(checkConnErrorWriter{c}, 4<<10)

    ctx, cancelCtx := context.WithCancel(ctx)
    defer cancelCtx()

    for {
        w, err := c.readRequest(ctx) // 1. 获取请求数据
        ...
        serverHandler{c.server}.ServeHTTP(w, w.req) // 2. 处理请求 serverHandler, 对应下面第5步
        w.cancelCtx()
        if c.hijacked() {
            return
        }
        w.finishRequest() // 3. 返回响应结果
        if !w.shouldReuseConnection() {
            if w.requestBodyLimitHit || w.closedRequestBodyEarly() {
                c.closeWriteAndWait()
            }
            return
        }
        c.setState(c.rwc, StateIdle)
    }
}

```

## 处理请求

```go
func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
    handler := sh.srv.Handler
    if handler == nil {
        handler = DefaultServeMux // ServeMux
    }
    if req.RequestURI == "*" && req.Method == "OPTIONS" {
        handler = globalOptionsHandler{}
    }
    handler.ServeHTTP(rw, req)
}
```

handler.ServeHTTP(rw, req)：

```go
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
    if r.RequestURI == "*" {
        if r.ProtoAtLeast(1, 1) {
            w.Header().Set("Connection", "close")
        }
        w.WriteHeader(StatusBadRequest)
        return
    }
    h, _ := mux.Handler(r) // HandlerFunc, Handler
    h.ServeHTTP(w, r)
}
```

执行处理：

```go
// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```

涉及的数据类型：

```go
type Server struct {
    Addr         string        // TCP address to listen on, ":http" if empty
    Handler      Handler       // handler to invoke, http.DefaultServeMux if nil
    ReadTimeout  time.Duration // maximum duration before timing out read of the request
    WriteTimeout time.Duration // maximum duration before timing out write of the response
    ...
}

type conn struct {
    server *Server // server is the server on which the connection arrived.
    rwc net.Conn // rwc is the underlying network connection. It is usually of type *net.TCPConn or *tls.Conn.
    remoteAddr string // This is the value of a Handler's (*Request).RemoteAddr.
    mu sync.Mutex // mu guards hijackedv, use of bufr, (*response).closeNotifyCh.
    ...
}

type serverHandler struct {
    srv *Server
}
```
