# 命令行应用

命令行应用是最基础，也是最常见的入门应用；我们可以基于这些模板快速地开发一些实用的小工具。

# flag

我们编写一个简单的示例，用于了解标准库 flag 的基本使用，代码如下：

```go
func main() {
	var name string
	flag.StringVar(&name, "name", "Go Series", "帮助信息")
	flag.StringVar(&name, "n", "Go Series", "帮助信息")
	flag.Parse()

	log.Printf("name: %s", name)
}
```

通过上述代码，我们调用标准库 flag 的 StringVar 方法实现了对命令行参数 name 的解析和绑定，其各个形参的含义分别为命令行标识位的名称、默认值、帮助信息。针对命令行参数，其支持如下三种命令行标志语法，分别如下：

- -flag：仅支持布尔类型。
- -flag x ：仅支持非布尔类型。
- -flag=x：均支持

同时 flag 标准库还提供了多种类型参数绑定的方式，根据各自的应用程序使用情况选用即可，接下来我们运行该程序，检查输出结果与预想的是否一致，如下：

```shell
$ go run main.go -name=eddycjy -n=Test
name: Test
```

## 子命令的实现

在我们日常使用的 CLI 应用中，另外一个最常见的功能就是子命令的使用，一个工具它可能包含了大量相关联的功能命令以此形成工具集，可以说是刚需，那么这个功能在标准库 flag 中可以如何实现呢，如下述示例：

```go
var name string

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) <= 0 {
		return
	}

	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "Go 语言", "帮助信息")
		_ = goCmd.Parse(args[1:])
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.StringVar(&name, "n", "PHP 语言", "帮助信息")
		_ = phpCmd.Parse(args[1:])
	}

	log.Printf("name: %s", name)
```

在上述代码中，我们首先调用了 flag.Parse 方法，将命令行解析为定义的标志，便于我们后续的参数使用。另外由于我们需要处理子命令的情况，因此我们调用了 flag.NewFlagSet 方法，该方法会返回带有指定名称和错误处理属性的空命令集给我们去使用，相当于就是创建一个新的命令集了去支持子命令了。这里需要特别注意的是 flag.NewFlagSet 方法的第二个参数是 ErrorHandling，用于指定处理异常错误的情况处理，其内置提供以下三种模式：

```go
const (
	// 返回错误描述
	ContinueOnError ErrorHandling = iota
	// 调用 os.Exit(2) 退出程序
	ExitOnError
	// 调用 panic 语句抛出错误异常
	PanicOnError
)
```

接下来我们运行针对子命令的示例程序，对正确和异常场景进行检查，如下：

```go
$ go run main.go go -name=Customer
name: Customer
```
