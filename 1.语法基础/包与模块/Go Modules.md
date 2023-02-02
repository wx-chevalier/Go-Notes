# Go Modules

Go Modules 将包名与路径分离，可以存放于文件系统上的任何为止，而不用管 GOPATH 路径到底是什么，我们可以创建任意的项目目录:

```sh
$ mkdir -p /tmp/scratchpad/hello
$ cd /tmp/scratchpad/hello
```

然后可以用 `go mod init example.com/m` 生成 go.mod 模板。模块根目录和其子目录的所有包构成模块，在根目录下存在 go.mod 文件，子目录会向着父目录、爷目录一直找到 go.mod 文件。模块路径指模块根目录的导入路径，也是其他子目录导入路径的前缀。go.mod 文件第一行定义了模块路径，有了这一行才算作是一个模块。go.mod 文件接下来的篇幅用来定义当前模块的依赖和依赖版本，也可以排除依赖和替换依赖。

```sh
module example.com/m

require (
    golang.org/x/text v0.3.0
    gopkg.in/yaml.v2 v2.1.0
    rsc.io/quote v1.5.2
)

replace (
    golang.org/x/text => github.com/golang/text v0.3.0
)
```

然后照常编写 Go 模块代码:

```go
// hello.go
package main

import (
    "fmt"
    "rsc.io/quote"
)

func main() {
    fmt.Println(quote.Hello())
}
```

在执行 `go build` 命令之后，即可以在 go.mod 文件中查看模块定义与显式的声明，它会自动地将未声明的依赖添加到 go.mod 文件中。

# 模块结构

模块是包含了 Go 源文件的目录树，并在根目录中添加了名为 go.mod 的文件，go.mod 包含模块导入名称，声明了要求的依赖项，排除的依赖项和替换的依赖项。

```sh
module my/thing

require (
        one/thing v1.3.2
        other/thing v2.5.0 // indirect
        ...
)

exclude (
        bad/thing v0.7.3
)

replace (
        src/thing 1.0.2 => dst/thing v1.1.0
)
```

需要注意的是，该文件中声明的依赖，并不会在模块的源代码中使用 import 自动导入，还是需要我们人工添加 import 语句来导入的。模块可以包含其他模块，在这种情况下，它们的内容将从父模块中排除。除了 go.mod 文件外，跟目录下还可以存在一个名为 go.sum 的文件，用于保存所有的依赖项的哈希摘要校验之，用于验证缓存的依赖项是否满足模块要求。

## 目录结构

一般来说，我们在 go.mod 中指定的名称是项目名，每个 package 中的名称需要和目录名保持一致。

```go
// go.mod
module myprojectname
// or
module github.com/myname/myproject
```

然后用如下方式导入其他模块：

```go
import myprojectname/stuff
import github.com/myname/myproject/stuff
```

# 外部依赖

模块依赖项会被下载并存储到 `GOPATH/src/mod` 目录中，直接后果就是废除了模块的组织名称。假设我们正在开发的项目依赖于 github.com/me/lib 且版本号 1.0.0 的模块，对于这种情况，我们会发现在 GOPATH/src/mod 中文件结构如下：

![Go Modules 缓存路径](https://s1.ax1x.com/2019/11/19/M2IIhD.png)

Go 的模块版本号必须以 v 开头，在发布版本时可以通过 Tag 方式来指定不同的版本。我们可以使用 `go list -m all` 来查看全部的依赖，使用 `go mod tidy` 来移除未被使用的依赖，使用 `go mod vendor` 可以生成独立的 vendor 目录。

# 模块代理

```sh
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
```

# Go Module 问题

go mod 是 rsc 主导设计的 Go 版本管理工具，借鉴了 Google 内部的高大上版本管理方式，摒弃了开源社区的版本管理成功经验，借助 MVS 算法，希望能够走出一条不一样的路，然而从发布以来给广大 Gopher 带来了各种各样的麻烦。

## Go 命令的副作用

Go list，Go test，Go build，所有命令都会去拉取依赖，有些库是用被墙的服务做了重定向，只是执行一下 go test，然后就被卡一年是家常便饭。

按照 "By design" 的说法，Google 内部的依赖库版本都会尽量使用能够兼容的最新版本。对于墙内的我们来说，我不管执行什么 Go 命令怎么都卡。逐渐患上 go test PTSD。

解法：配置 GOPROXY 代理，虽然拉取依赖还是慢。

## 形同虚设的 semver 规范

社区里不遵守 semver 规范的库很多，有的开源库在 1.7.4 ~ 1.7.5 中进行了 breaking change，而按照 semver 的定义，这是不应该发生的。go mod 过度高估了开源社区的节操。

## 无法应对删库

Go 号称分布式，但大多 Go 的依赖库都是存在 Github 上，如果 Github 上的原作者删除了该库，那么也会导致大多数的依赖用户 build 失败。

即使看起来我们可以靠 go.mod 和 go.sum 来实现 reproducible build，实际的情况是，像 K8s 这样的项目，依然会把庞大的依赖库放在自己 repo 的 vendor 里。

## Github release/tag 水土不服

在 Github 上发布 lib 的 release，或者给某个 commit 打 tag 之后，我们依然可以对这些 tag 和 release 进行编辑。我们经常看到，有些库的作者在发布一个 release 之后，又删除了这个 release，或对这个 release 进行了编辑。对于用户来说，这样就会依赖一个已经“消失”了的版本，在不存储 vendor 的情况下，reproducible build 沦为笑谈。

## goproxy 的实现并不统一

不知道是否是因为 goproxy 并无规范，在使用不同的代理帮助我们加速下载依赖时，会出现各种不同的错误。例如作者 A 开发的 goproxy，在某个库不存在时，会返回 404。而作者 B 开发的 goproxy，在某个库不存在时，会返回 500。着实令人困惑。

而 goproxy 本身的实现基本都是惰性下载，所以新发布的库，我们要走 goproxy 来测试时，就需要手动 go get 触发。而大多 goproxy 的实现并没有查询功能，goproxy 服务内部到底什么时候同步好了，可以 go get 了，还是 go get 的过程中发生失败了。作为用户是不可查的。

## go get 到的 lib 版本在 go build 时被修改

在 go get 时，可以 go get lib@ver 来获取指定版本的依赖，但是在 go build 时可能发现又被修改成了别的版本(比如被升级了)，非常反直觉。

## 版本信息扩散

由于 go mod 的设计，版本信息被包含在了 import 路径中。当依赖库从 v1 升级至 v2 时，几乎一定意味着我们代码中大量的 import 路径需要修改。

## go.sum 合并冲突

因为上面讲到的一系列问题，go.sum 在多人维护的大项目上，经常会发生变动，也就经常会有冲突。对于中心化版本管理系统来说，这个问题根本就不存在。对于 go mod 来说，go.sum 合并本来是个纯追加逻辑。

# Links

- https://mp.weixin.qq.com/s/Sxv5qb-v6OIhPptLWAHYUw Go Module 来了，企业私有代理你准备好了吗？
- https://colobu.com/2021/07/04/dive-into-go-module-3/?hmsr=toutiao.io 深入 Go Module 之未说的秘密
