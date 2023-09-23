# cobra

首先需要安装本项目所依赖的基础库 Cobra，便于我们后续快速搭建 CLI 应用程序，在项目根目录执行命令如下：

```sh
$ go get -u github.com/spf13/cobra@v1.0.0
```

```shell
tour
├── main.go
├── go.mod
├── go.sum
├── cmd
├── internal
└── pkg
```

在本项目中，我们创建了入口文件 main.go，并新增了三个目录，分别是 cmd、internal 以及 pkg，并在 `cmd` 目录下新建 word.go 文件，用于单词格式转换的子命令 word 的设置，写入如下代码：

```go
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  "支持多种单词格式转换",
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {}
```

接下来还是在 cmd 目录下，增加 root.go 文件，作为根命令，写入如下代码：

```go
var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
}
```

最后在启动 main.go 文件中，写入如下运行代码：

```go
func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
```

这里的 word 包负责实际的转换事宜，参阅[数据类型/字符串](/)一节。在完成了单词的各个转换方法后，我们开始编写 word 子命令，将其对应的方法集成到我们的 Command 中，打开项目下的 cmd/word.go 文件，定义目前单词所支持的转换模式枚举值，新增代码如下：

```go
const (
	ModeUpper                      = iota + 1 // 全部转大写
	ModeLower                                 // 全部转小写
	ModeUnderscoreToUpperCamelCase            // 下划线转大写驼峰
	ModeUnderscoreToLowerCamelCase            // 下线线转小写驼峰
	ModeCamelCaseToUnderscore                 // 驼峰转下划线
)
```

接下来进行具体的单词子命令的设置和集成，继续新增如下代码：

```go
var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部转大写",
	"2：全部转小写",
	"3：下划线转大写驼峰",
	"4：下划线转小写驼峰",
	"5：驼峰转下划线",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
		}

		log.Printf("输出结果: %s", content)
	},
}
```

在上述代码中，核心在于子命令 word 的 `cobra.Command` 调用和设置，其一共包含如下四个常用选项，分别是：

- Use：子命令的命令标识。
- Short：简短说明，在 help 输出的帮助信息中展示。
- Long：完整说明，在 help 输出的帮助信息中展示。

接下来我们根据单词转换所需的参数，分别是单词内容和转换的模式进行命令行参数的设置和初始化，继续写入如下代码：

```go
var str string
var mode int8

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
```

在 VarP 系列的方法中，第一个参数为需绑定的变量、第二个参数为接收该参数的完整的命令标志，第三个参数为对应的短标识，第四个参数为默认值，第五个参数为使用说明。
在完成了单词格式转换的功能后，已经初步的拥有了一个工具了，现在我们来验证一下功能是否正常，一般我们拿到一个 CLI 应用程序，我们会先执行 help 来先查看其帮助信息，如下：

```shell
$ go run main.go help word
该子命令支持各种单词格式转换，模式如下：
1：全部转大写
2：全部转小写
3：下划线转大写驼峰
4：下划线转小写驼峰
5：驼峰转下划线

Usage:
   word [flags]

Flags:
  -h, --help         help for word
  -m, --mode int8    请输入单词转换的模式
  -s, --str string   请输入单词内容
```

手工验证四种单词的转换模式的功能点是否正常，如下：

```shell
$ go run main.go word -s=test -m=1
输出结果: TEST
```
