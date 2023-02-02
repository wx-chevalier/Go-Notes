# Web 应用

Go 语言里面提供了一个完善的 net/http 包，通过 http 包可以很方便的就搭建起来一个可以运行的 Web 服务。同时使用这个包能很简单地对 Web 的路由，静态文件，模版，cookie 等进行设置和操作。

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("path:", r.URL.Path)
	fmt.Fprintf(w, "hello go")
}

func main() {
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
```

# HTTP 包运行机制

Go 实现 Web 服务流程如下

1. 创建 `Listen Socket`, 监听指定的端口, 等待客户端请求到来。

2. `Listen Socket` 接受客户端的请求, 得到 `Client Socket`, 接下来通过 `Client Socket` 与客户端通信。

3. 处理客户端的请求, 首先从 `Client Socket` 读取 HTTP 请求, 然后交给相应的 `handler` 处理请求, 最后将 `handler` 处理完毕的数据, 通过 `Client Socket` 写给客户端。

其中涉及服务器端的概念:

- Request：用户请求的信息，用来解析用户的请求信息，包括 post、get、cookie、url 等信息
- Conn：用户的每次请求链接
- Handler：处理请求和生成返回信息的处理逻辑
- Response：服务器需要反馈给客户端的信息

# 服务监听与请求处理过程

Go 是通过一个 ListenAndServe 监听服务，底层处理：初始化一个 server 对象，然后调用 `net.Listen("tcp", addr)`，监控我们设置的端口。监控端口之后，调用 `srv.Serve(net.Listener)` 函数，处理接收客户端的请求信息。首先通过 Listener 接收请求，其次创建一个 Conn，最后单独开了一个 goroutine，把这个请求的数据当做参数扔给这个 conn 去服务。go c.serve() 用户的每一次请求都是在一个新的 goroutine 去服务，相互不影响。

分配相应的函数处理请求: conn 首先会解析 request:c.readRequest(), 然后获取相应的 handler:handler := c.server.Handler，这个是调用函数 ListenAndServe 时候的第二个参数，例子传递的是 nil，也就是为空，那么默认获取 `handler = DefaultServeMux。DefaultServeMux` 是一个路由器，它用来匹配 url 跳转到其相应的 handle 函数。调用 `http.HandleFunc("/", sayhelloName)` 作用是注册了请求/的路由规则，将 url 和 handle 函数注册到 DefaultServeMux 变量，最后调用 DefaultServeMux 的 ServeHTTP 方法，这个方法内部调用 handle 函数。

![服务监听与请求处理流程图](https://s2.ax1x.com/2019/12/02/Qnvu3F.png)

# 自定义路由实现

定义的类型实现 ServeHTTP 方法，即可实现自定义路由

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

type MyMux struct {}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        sayhelloName(w, r)
        return
    }

    http.NotFound(w, r)
    return
}


func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println("path:", r.URL.Path)
    fmt.Fprintf(w, "hello go")
}

func main() {
    mux := &MyMux{}
    err := http.ListenAndServe(":9090", mux)
    if err != nil {
        log.Fatal("ListenAndServer: ", err)
    }
}
```
