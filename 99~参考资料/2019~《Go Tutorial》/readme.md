# Learning notes for golang

如果发现了本项目里的问题或者想了解本项目里还没涉及到的 go 语言技术栈，欢迎提交[issue](https://github.com/jincheng9/go-tutorial/issues/new)。

如果觉得本项目不错，欢迎给个**Star**。

## 基础篇

- [lesson0: GitHub 最流行的 golang-cheat-sheet 中文版本](https://github.com/jincheng9/golang-cheat-sheet-cn)

- [lesson1: Go 程序结构](./workspace/lesson1)

- [lesson2: 数据类型：数字，字符串，bool](./workspace/lesson2)

- [lesson3: 变量类型定义：全局变量，局部变量，多变量，零值](./workspace/lesson3)

- [lesson4: 常量和枚举](./workspace/lesson4)

- [lesson5: 运算操作符](./workspace/lesson5)

- [lesson6: 控制语句 if/switch](./workspace/lesson6)

- [lesson7: 循环语句 for/goto/break/continue](./workspace/lesson7)

- [lesson8: 函数，闭包和方法](./workspace/lesson8)

- [lesson9: 变量作用域](./workspace/lesson9)

- [lesson10: 数组：一维数组和多维数组](./workspace/lesson10)

- [lesson11: 指针 pointer](./workspace/lesson11)

- [lesson12: 结构体 struct](./workspace/lesson12)

- [lesson13: 切片 Slice](./workspace/lesson13)

- [lesson14: range 迭代](./workspace/lesson14)

- [lesson15: map 集合](./workspace/lesson15)

- [lesson16: 递归函数](./workspace/lesson16)

- [lesson17: 类型转换](./workspace/lesson17)

- [lesson18: 接口 interface](./workspace/lesson18)

- [lesson19: 协程 goroutine 和管道 channel](./workspace/lesson19)

- [lesson20: defer 语义](./workspace/lesson20)

- [lesson21: 并发编程之 sync.WaitGroup](./workspace/lesson21)

- [lesson22: 并发编程之 sync.Once](./workspace/lesson22)

- [lesson23: 并发编程之 sync.Mutex 和 sync.RWMutex](./workspace/lesson23)

- [lesson24: 并发编程之 sync.Cond](./workspace/lesson24)

- [lesson25: 并发编程之 sync.Map](./workspace/lesson25)

- [lesson26: 并发编程之原子操作 sync/atomic](./workspace/lesson26)

- [lesson27: 包 Package 和模块 Module](./workspace/lesson27)

- [lesson28: panic, recover 运行期错误处理](./workspace/lesson28)

- [lesson29: select 语义](./workspace/lesson29)

- [lesson30: go 单元测试](./workspace/lesson30)

- [lesson31: go 性能测试](./workspace/lesson31)

- [lesson32: go 模糊测试](./workspace/senior/p22)

## 进阶篇

- 常用关键字

  - [被 defer 的函数一定会执行么？](./workspace/senior/p2)
  - [new 和 make 的使用区别和最佳实践是什么？](./workspace/senior/p4)
  - [receive-only channel 和 send-only channel 的争议](./workspace/senior/p30)

- 语言基础
  - [Go 有引用变量和引用传递么？map,channel 和 slice 作为函数参数是引用传递么？](./workspace/senior/p3)
  - [一文读懂 Go 匿名结构体的使用场景](./workspace/senior/p5)
  - [Go 语言中 fmt.Println(true)的结果一定是 true 么？](./workspace/senior/p19)
  - [Go 语言中命名函数参数和命名函数返回值的注意事项](./workspace/senior/p21)
- 并发编程
  - [Go 语言中 Context 并发模式](./workspace/senior/p31/01-go-context.md)
- 泛型
  - [泛型：Go 泛型入门官方教程](./workspace/senior/p6)
  - [泛型：一文读懂 Go 泛型设计和使用场景](./workspace/senior/p7)
  - [泛型：Go 1.18 正式版本将从标准库中移除 constraints 包](./workspace/senior/p17)
  - [泛型：什么场景应该使用泛型](./workspace/official-blog/when-to-use-generics.md)
- Fuzzing
  - [Fuzzing: Go Fuzzing 入门官方教程](./workspace/senior/p22)
  - [Fuzzing: 一文读懂 Go Fuzzing 使用和原理](./workspace/senior/p23)
- Workspace mode 工作区模式
  - [Go 1.18：工作区模式 workspace mode 简介](./workspace/senior/p25)
  - [Go 1.18：工作区模式最佳实践](./workspace/official-blog/go1.18-workspace-best-practice.md)
- 语言规范
  - [Practical Go：如何写出更好维护的 Go 代码](https://github.com/jincheng9/practical-go-cn)
  - [Google 的 Go 语言编码规范](./workspace/style/google.md)
- Go 开发中的常见错误
  - [第 1 篇：go 未知枚举值](./workspace/senior/p28/01-unknown-enum-value.md)
  - [第 2 篇：go benchmark 性能测试的坑](./workspace/senior/p28/02-go-benchmark-inline.md)
  - [第 3 篇：go 指针的性能问题和内存逃逸](./workspace/senior/p28/03-go-pointer.md)
  - [第 4 篇：for/switch 和 for/select 做 break 操作退出的注意事项](./workspace/senior/p28/04-break-for-switch-select.md)
  - [第 5 篇：go 语言 Error 管理](./workspace/senior/p28/05-go-error-management.md)
  - [第 6 篇：slice 初始化常犯的错误](./workspace/senior/p28/06-go-slice-init.md)
  - [第 7 篇：不使用-race 选项做并发竞争检测](./workspace/senior/p28/07-go-race-detector.md)
  - [第 8 篇：并发编程中 Context 使用常见错误](./workspace/senior/p28/08-go-context-management.md)
  - [第 9 篇：使用文件名称作为函数输入](./workspace/senior/p28/09-go-use-filename-as-input.md)
  - [第 10 篇：Goroutine 和循环变量一起使用的坑](./workspace/senior/p28/10-go-goroutine-loop-var.md)
  - [第 11 篇：意外的变量遮蔽(variable shadowing)](./workspace/senior/p28/11-go-unintended-variable-shadowing.md)
  - [第 12 篇：冗余的嵌套代码](./workspace/senior/p28/12-go-unnecessary-nested-code.md)
  - [第 13 篇：init 函数的常见错误和最佳实践](./workspace/senior/p28/13-go-package-init-function.md)
  - [第 14 篇：过度使用 getter 和 setter 方法](./workspace/senior/p28/14-go-overuse-getter-setter.md)
  - [第 15 篇：interface 使用的常见错误和最佳实践](./workspace/senior/p28/15-go-interface-pollution.md)
  - [第 16 篇：any 的常见错误和最佳实践](./workspace/senior/p28/16-any-keyword.md)
- 高性能 Go
  - [一文读懂 Go 1.20 引入的 PGO 性能优化](./workspace/senior/p35)
  - [high performance go workshop](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html)
- Go 安全
  - [Go 的全新安全漏洞检测工具 govulncheck 来了](./workspace/senior/p32/01-go-vulnerability-management.md)
- Go 版本演进
  - [Go 1.19 要来了，看看都有哪些变化：第 1 篇](./workspace/senior/p29/readme.md)
  - [Go 1.19 要来了，看看都有哪些变化：第 2 篇](./workspace/senior/p29/readme2.md)
  - [Go 1.19 要来了，看看都有哪些变化：第 3 篇](./workspace/senior/p29/readme3.md)
  - [Go 1.19 要来了，看看都有哪些变化：第 4 篇](./workspace/senior/p29/readme4.md)
  - [Go 1.20 要来了，看看都有哪些变化：第 1 篇](./workspace/senior/p33/readme.md)
  - [Go 1.20 要来了，看看都有哪些变化：第 2 篇](./workspace/senior/p33/readme2.md)
  - [Go 1.20 要来了，看看都有哪些变化：第 3 篇](./workspace/senior/p33/readme3.md)
  - [Go 1.20 要来了，看看都有哪些变化：第 4 篇](./workspace/senior/p33/readme4.md)
  - [Go 1.21 的 2 个语言变化](./workspace/senior/p34)

## Go Quiz

1. [Go Quiz: 从 Go 面试题看 slice 的底层原理和注意事项](./workspace/senior/p8)

2. [Go Quiz: 从 Go 面试题搞懂 slice range 遍历的坑](./workspace/senior/p13)

3. [Go Quiz: 从 Go 面试题看 channel 的注意事项](./workspace/senior/p9)

4. [Go Quiz: 从 Go 面试题看 channel 在 select 场景下的注意事项](./workspace/senior/p14)

5. [Go Quiz: 从 Go 面试题看分号规则和 switch 的注意事项](./workspace/senior/p10)

6. [Go Quiz: 从 Go 面试题看 defer 语义的底层原理和注意事项第 1 篇](./workspace/senior/p11)

7. [Go Quiz: 从 Go 面试题看 defer 的注意事项第 2 篇](./workspace/senior/p12)

8. [Go Quiz: 从 Go 面试题看 defer 的注意事项第 3 篇](./workspace/senior/p15)

9. [Go Quiz: Google 工程师的 Go 语言题目](./workspace/senior/p16)

10. [Go Quiz: 从 Go 面试题看 panic 注意事项第 1 篇](./workspace/senior/p18)

11. [Go Quiz: 从 Go 面试题看 recover 注意事项第 1 篇](./workspace/senior/p18/readme2.md)

12. [Go Quiz: 从 Go 面试题看函数命名返回值的注意事项](./workspace/senior/p20)

13. [Go Quiz: 从 Go 面试题看锁的注意事项](./workspace/senior/p24)

14. [Go Quiz: 从 Go 面试题看变量的零值和初始化赋值的注意事项](./workspace/senior/p26)

15. [Go Quiz: 从 Go 面试题看数值类型的自动推导](./workspace/senior/p27)

16. [Go questions-golang.design](https://golang.design/go-questions/)

## Go 标准库

- [Go 标准库脑图](./workspace/img/go-std-lib-mindmap.png)

- [Go 标准库之 log 使用详解](./workspace/std/01)

- [Go 标准库之 cmd 命令使用详解](./workspace/std/02)

## 实战篇

### 代码规范

- [Google 的 Go 语言编码规范](./workspace/style/google.md)
- [Practical Go 中文版本](https://github.com/jincheng9/practical-go-cn)
- [Go 谚语 by Rob Pike](https://go-proverbs.github.io/)

### Web 框架

#### Gin

- [当前流行的 Go web 框架比较以及 Gin 介绍](./workspace/gin/01/)
- [Gin 源码结构解析](./workspace/gin/02)

### RPC

#### gRPC

- [gRPC 入门指引](./workspace/rpc/01)
- [gRPC-Go 入门教程](./workspace/rpc/02)

### Databases

#### MySQL

- [Tutorial of go-sql-driver/mysql](./workspace/mysql/01/)

#### Redis

- [Tutorial of go-redis/redis](./workspace/redis/01/)

### Docker/K8s

- [Docker 入门教程 101: 用途，架构，安装和使用](https://github.com/jincheng9/disributed-system-notes/tree/main/docker/01)
- [Docker 入门教程 101: 基于 Docker 部署 Go 项目](https://github.com/jincheng9/disributed-system-notes/tree/main/docker/02)

### Document Tools

#### Swagger

- [gin-swagger 常见问题](./workspace/swagger)

## CI/CD

- [Jenkins 教程 01：安装部署](./workspace/devops/jenkins01.md)

## 外文翻译

1. [GitHub 最流行的 golang-cheat-sheet 中文版本](https://github.com/jincheng9/golang-cheat-sheet-cn)

1. [官方博文：Go 开源 13 周年](./workspace/official-blog/13-years-of-go.md)

1. [官方博文：Go 开发者调研方式改变了](./workspace/official-blog/survey-change.md)

1. [官方博文：什么场景应该使用泛型](./workspace/official-blog/when-to-use-generics.md)

1. [官方博文：Go 工作区模式最佳实践](./workspace/official-blog/go1.18-workspace-best-practice.md)

1. [官方博文：Go 1.18 发布啦！](./workspace/official-blog/go118_release.md)

1. [官方教程：Go fuzzing 模糊测试](./workspace/senior/p22)

1. [官方教程：Go 泛型入门](./workspace/senior/p6)

1. [官方博文：Go 1.18 Beta 2 发布](./workspace/official-blog/go118_beta2.md)

1. [官方博文：Go 官方推出了 Go 1.18 的 2 个新教程](./workspace/official-blog/go118_two_new_tutorial.md)

1. [官方博文：支持泛型的 Go 1.18 Beta 1 版本正式发布](./workspace/official-blog/go118beta1.md)

1. [官方博文：Go 开源 12 周年](./workspace/official-blog/twelve-years-of-go.md)

## Go 环境和工具

1. [GitHub 上的项目 go get 连不上怎么办？](./workspace/senior/p1/)

2. [GoLand 常用快捷键](./workspace/senior/p1/readme2.md)

3. [Mac 的 shell 切换、环境变量设置以及软件安装问题](./workspace/senior/p1/readme3.md)

4. [Go testing 缓存导致测试没执行的问题](./workspace/senior/p1/readme4.md)

5. [go install 安装的不同 Go 版本的可执行程序和源码存放在哪里](./workspace/senior/p1/readme5.md)

6. [Mac 系统查看 Go 开发相关的系统设置](./workspace/senior/p1/readme6.md)

## Go Book

- [The Go Programming Language-Go 语言圣经](http://www.gopl.io/)
- [Go 语言高级编程-chai2010.gitbooks.io](https://chai2010.gitbooks.io/advanced-go-programming-book/content/)
- [Go 语言设计与实现-draveness.me](https://draveness.me/golang/)
- [Go 设计模式-Tamer Tas@google](https://github.com/tmrts/go-patterns)
- [深入解析 Go-tiancaiamao.gitbooks.io](https://tiancaiamao.gitbooks.io/go-internals/content/zh/)
- [码农桃花源-qcrao91.gitbook.io](https://qcrao91.gitbook.io/go/)
- [Go 语言高性能编程-geektutu](https://geektutu.com/post/hpg-benchmark.html)
- [Go Under The Hood-golang.design](https://golang.design/under-the-hood/)
- [英文 Go 书籍 list](https://github.com/dariubs/GoBooks)

## Go Blog

- [Jincheng's Blog](https://jincheng9.github.io/)
- [Russ Cox-Go 团队负责人](https://research.swtch.com/)
- [Go Documentation](https://go.dev/doc/)
- [Golang GitHub Wiki](https://github.com/golang/go/wiki)
- [Go By Example](https://gobyexample.com/)
- [Golang By Example](https://golangbyexample.com/)
- [CS Professor Nilsson from KTH](https://yourbasic.org/golang/)
- [John Arundel](https://bitfieldconsulting.com/)
- [Dave Cheney](https://dave.cheney.net/)
- [Jaana Dogan-Pricipal Engineer at AWS](https://rakyll.org/about/)
- [go101.org](https://go101.org/article/101.html)
- [Valentin Deleplace-Google Engineer](https://medium.com/@val_deleplace)
- [Jay Conrod-Ex Google Go Team Member](https://jayconrod.com/posts)
- [Medium: A Journey with Go](https://medium.com/a-journey-with-go)
- [Teiva Harsanyi-100 Go Mistakes author](https://teivah.github.io/)
- [Carl M. Johnson.net-Tech Director of Spotlight PA](https://carlmjohnson.net/)
- [Alex Edwards-A full stack Web Developer](https://www.alexedwards.net/blog)
- [golang.design](https://golang.design/)
- [Amit Saha-Atlassian Engineer](https://echorand.me/)
- [Paschalis Tsilias-Grafana Engineer](https://tpaschalis.github.io/)
- [liwenzhou-李文周](https://www.liwenzhou.com/)
- [TalkGo 发起人-杨文](https://maiyang.me/)
- [smallest-rpcx 作者](https://colobu.com/)
- [ChangkunOu-欧长坤](https://changkun.de/blog/)
- [chai2010-柴树杉](https://chai2010.cn/about/)
- [cch123-曹春晖](https://xargin.com/)
- [halfrost-Dezhi Yu](https://halfrost.com/)
- [draveness-左书祺](https://draveness.me/)
- [unknwon-无闻](https://unknwon.io/)
- [strikefreedom-Andy Pan](https://strikefreedom.top/)
- [qcrao-码农桃花源](https://qcrao.com/)
- [geektutu-极客兔兔](https://geektutu.com/)
- [topgoer.com](https://www.topgoer.com/)
- [topgoer.cn](https://topgoer.cn/)
- [涛叔-taoshu.in](https://taoshu.in/)
- [jianyu chen](https://eddycjy.com/)
- [Zhiyun Luo-Tencent IEG Developer](https://www.luozhiyun.com/)

## Go Video

### YouTube

- [Gopher Academy](https://www.youtube.com/@GopherAcademy)

- [GopherCon Talks Since 2014](https://www.youtube.com/c/GopherAcademy/videos)

- [GoLab Conference Since 2018](https://www.youtube.com/channel/UCMEvzoHTIdZI7IM8LoRbLsQ/featured)

- [Basics, Intermediate, Advanced Go Tutorials-Bitfield Consulting](https://www.youtube.com/c/BitfieldConsulting)

- [TutorialEdge Golang Development](https://www.youtube.com/watch?v=W5b64DXeP0o&list=PLzUGFf4GhXBL4GHXVcMMvzgtO8-WEJIoY)

## Go Community

- GoCN： https://github.com/gocn/opentalk
- Go 夜读：https://github.com/talkgo/night

## Go News

- https://go.dev/blog/
- https://twitter.com/_rsc
- https://twitter.com/rob_pike
- https://twitter.com/golang
- https://twitter.com/golangweekly
- https://twitter.com/GolangTrends

## 微信公众号

- coding 进阶：分享 Go 语言入门、中级到高级教程，以及微服务、云原生架构

  ![coding进阶](./workspace/img/wechat.png)

## 微信赞助

![img](./workspace/img/wechat-payment.png)

## 微信群交流

加我微信，入群交流

![](./workspace/img/wechat-group.png)
