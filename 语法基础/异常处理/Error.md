# Error

Go 语言中并不存在 try-catch 等异常处理的关键字，对于那些可能返回异常的函数，只需要在函数返回值中添加额外的 Error 类型的返回值：

```go
type error interface {
    Error() string
}
```

错误信息不要以大写字母开头或是以句点结尾。因为他们通常会在一个上下文中被打印出来。某个可能返回异常的函数调用方式如下：

```go
import (
    "fmt"
    "errors"
)

func main() {
    result, err:= Divide(2,0)

    if err != nil {
            fmt.Println(err)
    }else {
            fmt.Println(result)
    }
}

func Divide(value1 int, value2 int)(int, error) {
    if(value2 == 0){
        return 0, errors.New("value2 mustn't be zero")
    }
    return value1/value2 , nil
}
```

# 获取错误的上下文

有时候为了方便上层用户理解；底层实现者会将底层的错误重新包装为新的错误类型返回给用户：

```go
if _, err := html.Parse(resp.Body); err != nil {
	return nil, fmt.Errorf("parsing %s as HTML: %v", url,err)
}
```

为了记录这种错误类型在包装的变迁过程中的信息，我们一般会定义一个辅助的 WrapError 函数，用于包装原始的错误，同时保留完整的原始错误类型。为了问题定位的方便，同时也为了能记录错误发生时的函数调用状态，我们很多时候希望在出现致命错误的时候保存完整的函数调用信息。同时，为了支持 RPC 等跨网络的传输，我们可能要需要将错误序列化为类似 JSON 格式的数据，然后再从这些数据中将错误解码恢出来。

```go
type Error interface {
	Caller() []CallerInfo
	Wraped() []error
	Code() int
	error

	private()
}

type CallerInfo struct {
	FuncName string
	FileName string
	FileLine int
}
```

其中`Error`为接口类型，是`error`接口类型的扩展，用于给错误增加调用栈信息，同时支持错误的多级嵌套包装，支持错误码格式。为了使用方便，我们可以定义以下的辅助函数：

```go
func New(msg string) error
func NewWithCode(code int, msg string) error

func Wrap(err error, msg string) error
func WrapWithCode(code int, err error, msg string) error

func FromJson(json string) (Error, error)
func ToJson(err error) string
```

`New`用于构建新的错误类型，和标准库中`errors.New`功能类似，但是增加了出错时的函数调用栈信息。`FromJson`用于从 JSON 字符串编码的错误中恢复错误对象。`NewWithCode`则是构造一个带错误码的错误，同时也包含出错时的函数调用栈信息。`Wrap`和`WrapWithCode`则是错误二次包装函数，用于将底层的错误包装为新的错误，但是保留的原始的底层错误信息。这里返回的错误对象都可以直接调用`json.Marshal`将错误编码为 JSON 字符串。

我们可以这样使用包装函数:

```go
import (
	"dev.wx/errors"
)

func loadConfig() error {
	_, err := ioutil.ReadFile("/path/to/file")
	if err != nil {
		return errors.Wrap(err, "read failed")
	}

	// ...
}

func setup() error {
	err := loadConfig()
	if err != nil {
		return errors.Wrap(err, "invalid config")
	}

	// ...
}

func main() {
	if err := setup(); err != nil {
		log.Fatal(err)
	}

	// ...
}
```

上面的例子中，错误被进行了 2 层包装。我们可以这样遍历原始错误经历了哪些包装流程：

```go
for i, e := range err.(errors.Error).Wraped() {
    fmt.Printf("wraped(%d): %v\n", i, e)
}
```

同时也可以获取每个包装错误的函数调用堆栈信息：

```go
for i, x := range err.(errors.Error).Caller() {
    fmt.Printf("caller:%d: %s\n", i, x.FuncName)
}
```

如果需要将错误通过网络传输，可以用 `errors.ToJson(err)` 编码为 JSON 字符串：

```go
// 以JSON字符串方式发送错误
func sendError(ch chan<- string, err error) {
	ch <- errors.ToJson(err)
}

// 接收JSON字符串格式的错误
func recvError(ch <-chan string) error {
	p, err := errors.FromJson(<-ch)
	if err != nil {
		log.Fatal(err)
	}
	return p
}
```

对于基于 HTTP 协议的网络服务，我们还可以给错误绑定一个对应的 HTTP 状态码：

```go
err := errors.NewWithCode(404, "http error code")

fmt.Println(err)
fmt.Println(err.(errors.Error).Code())
```

在 Go 语言中，错误处理也有一套独特的编码风格。检查某个子函数是否失败后，我们通常将处理失败的逻辑代码放在处理成功的代码之前。如果某个错误会导致函数返回，那么成功时的逻辑代码不应放在`else`语句块中，而应直接放在函数体中。

```go
f, err := os.Open("filename.ext")
if err != nil {
	// 失败的情形, 马上返回错误
}

// 正常的处理流程
```

Go 语言中大部分函数的代码结构几乎相同，首先是一系列的初始检查，用于防止错误发生，之后是函数的实际逻辑。

# 延时处理

让我们演示一个文件复制的例子：函数需要打开两个文件，然后将其中一个文件的内容复制到另一个文件：

```go
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	written, err = io.Copy(dst, src)
	dst.Close()
	src.Close()
	return
}
```

上面的代码虽然能够工作，但是隐藏一个 bug。如果第一个`os.Open`调用成功，但是第二个`os.Create`调用失败，那么会在没有释放`src`文件资源的情况下返回。虽然我们可以通过在第二个返回语句前添加`src.Close()`调用来修复这个 BUG；但是当代码变得复杂时，类似的问题将很难被发现和修复。我们可以通过`defer`语句来确保每个被正常打开的文件都能被正常关闭：

```go
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
```

defer 语句可以让我们在打开文件时马上思考如何关闭文件。不管函数如何返回，文件关闭语句始终会被执行。同时 defer 语句可以保证，即使 io.Copy 发生了异常，文件依然可以安全地关闭。
