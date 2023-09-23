# How to start a Go project in 2023

## Install / Setup

The first thing to do is download and install Go. I would suggest always installing from the Go website itself https://golang.org/ and following the instructions for your OS of choice. Other than the Go 1.18 release (which included generics) I have never had any issue always installing the newest version of Go and compiling away. The backwards compatibility promise is real. So much so even if a project’s go.mod file says 1.20, if it does not use any 1.20 functionality you can probably still compile it using an earlier release.

Older guides will mention setting up your $GOPATH. This is something you can comfortably ignore in 2023. Check my [previous post](https://boyter.org/posts/how-to-start-go-project-2018/) if you are curious. However everything has or is moving to modules, so just consider this something you don’t have to learn.

One thing I do recommend is update your machines path to point to the `bin` directory of the default $GOPATH, `export PATH=$PATH:$(go env GOPATH)/bin` so that you can install anything you are working on quickly and have it available everywhere. For example, on my current machine it contains the following.

```shell
# boyter @ Bens-MacBook-Air in ~/go/bin [10:04:57]
$ tree
.
├── boltbrowser
├── cs
├── dcd
├── goreleaser
├── gow
├── hashit
├── lc
├── scc
```

My dotfile contains `export PATH=/Users/boyter/go/bin:$PATH` to achieve the above.

With the exported path I can run `scc`, `dcd` and `lc` anywhere I want, after first running `go install` for the project I am working on. When on Windows I go a little further than most and have my $GOPATH shared between Windows and Linux using the WSL so I can work on the same code-base in either, and as such you would see .exe files in the above. To do that you just need to do a symlink inside the WSL to the Windows directory.

### Editor

Based on the most recent [Go survey results](https://go.dev/blog/survey2023-q1-results) most people code Go in Visual Studio Code or Jetbrains Goland, which are the only editors I am going to include.

Goland works pretty much out of the box, and is the IDE I use day to day. So long as you have Go installed it should find it and start working. The biggest issue is when new releases of Go come out. This can cause Goland to be confused, and break things like your debugger. The solution is to upgrade after a week or so once the Jetbrains team has updated to match.

I have found these updates sometimes result result in an infinite loop of The “clang” command requires the command line developer tools on macOS. The fix for this I [wrote about](https://boyter.org/posts/golang-the-clang-command-requires-command-line-developer-tools/) but boils down to running `xcodebuild -runFirstLaunch` to resolve the issue.

Visual Studio Code has come a long way and is a lot better than it was when I first started with Go. Install the [latest version](https://code.visualstudio.com/), and then install the [Go in Visual Studio Code](https://code.visualstudio.com/docs/languages/go) extension. You get intellisense, reformatting, auto-import/remove import functionality you want. Debugging supposedly works now as well, but I cannot report on this.

I still prefer to pay and use for Goland because I find it to be like pairing with a brilliant engineer who never sleeps and is almost never wrong. Its ability to generate table tests, and run individual ones saves a lot of time and the refactoring tools are great. However for this post I tried using Visual Studio Code for a few hours and I was very impressed, and have no problems recommending it now.

## Starting a Project

Starting a project is as easy as starting a new directory and running `go mod init NAMEHERE` where `NAMEHERE` is the name of the package you want for your project. It used to be that you used a name that matched the location of your repository so for example `github.com/boyter/scc` but you can use whatever you want now. Using the full repo URL isn’t a bad idea though and I still prefer it for most projects.

### Getting Packages / Dependencies

Getting a package is almost as simple as knowing its path and using `go get URL` to download it to your local system. I usually vendor dependencies so I have a copy stored with my project allowing for reproducible builds. It also allows me to patch bugs in the dependencies easily while waiting on upstream fixes. To do so run `go mod vendor` which will pull everything into the `vendor` directory. If you do this I suggest setting up a `.ignore` file with `vendor` in it.

Where getting packages can be confusing is if the package maintainer has moved on from semantic version 1 to 2 or further. In this case you will need to add the version you want at the end to pull in the version you want.

For example my project `scc` is on version 3.1.0. If I were to import it without specifying the version,

```shell
$ go get github.com/boyter/scc/
go: added github.com/boyter/scc v2.12.0+incompatible
```

I would get a version 2.12 package which can be confusing to those new to Go. When adding the latest version,

```shell
$ go get github.com/boyter/scc/v3
go: added github.com/boyter/scc/v3 v3.1.0
```

Which as you an see has pulled down the correct version which is what I would expect.

The following guide [Just tell me how to use Go Modules](https://engineering.kablamo.com.au/posts/2018/just-tell-me-how-to-use-go-modules/) and it’s [Hacker News Conversation](https://news.ycombinator.com/item?id=18653225) covers this fairly well.

### Clean / Tidy

One thing that will come up occasionally when you try to run `go` after working with packages is it reporting you need to run `go mod tidy`. Don’t worry too much about this, just run `go mod tidy` and whatever you were trying to do till you are able to progress again. You can read about what its doing at [Go Modules Reference](https://go.dev/ref/mod).

Cached artifacts from the Go build system can be stored on your local system and take up a fair amount of space (my local at time of writing is ~1GB in size). To clean this up run `go clean -cache`.

## Learning Go

You can get to grips with Go pretty easily using the [go.dev learning tutorials](https://go.dev/learn/). This will get you up to speed with how to write code, and the syntax you need to be aware of. However for learning how to structure your own HTTP application, which is what most people are doing I strongly suggest the following book https://lets-go.alexedwards.net/ It does cost money, but it will short-cut your learning process by a few hours.

I have a sample that I personally use for setting up new HTTP projects which you can find on github https://github.com/boyter/go-http-template

One thing I strongly suggest reading however is the [50 Shades of Go](https://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html) post. It covers a lot of the Go pitfalls you are likely to run into. Checking for this is something a lot of companies screen for when hiring, as exposure to these issues is a good indication of experience using Go.

## Searching

To search for anything about Go in your search engine of choice use the word `golang` rather than `go` when searching. For example to search for how to open a file I would search for `golang open file`.

Note its best to not refer to the language as golang in casual conversation, as this will annoy a lot of pedantic people. Everyone knows what you are talking about but expect someone to say something eventually.

## Building / Installing

For commands which have `package main`

```shell
go build   builds the command and leaves the result in the current working directory.
go install builds the command in a temporary directory then moves it to $GOPATH/bin.
```

For packages

```shell
go build   builds your package then discards the results.
go install builds then installs the package in your $GOPATH/pkg directory.
```

If you want to cross compile, that is build on Linux for Windows or vice versa you can set what architecture your want to target and the OS through environment variables. You can view your defaults in `go env` but to change them you would do something like,

```shell
GOOS=darwin GOARCH=amd64 go build
GOOS=darwin GOARCH=arm64 go build
GOOS=windows GOARCH=amd64 go build
GOOS=windows GOARCH=arm64 go build
GOOS=linux GOARCH=amd64 go build
GOOS=linux GOARCH=arm64 go build
```

### Trimming Builds

Go binaries are by default “fat” and larger than you might expect. There is an easy way to reduce the size,

```shell
go build -ldflags="-s -w"
```

Which strips out the debug information. For smaller binaries where startup time does not matter you can also use https://upx.github.io/ but I have found issues with using this when cross compiling. See this [other post](https://boyter.org/posts/trimming-golang-binary-fat/) I wrote about using both.

### Packaging / Deploying

While you can use the above mentioned GOOS and GOARCH to build your own packages, I strongly suggest using [goreleaser](https://goreleaser.com/). It makes deployments considerably easier and its guide ensures you are tagging correctly.

## Linting / Static Analysis / Security Scanning

While you can use sonar and various other tools for this, I prefer to have something you can run locally, and easily integrate into your CI/CD system. Using the below tools will get you those all important audit ticks.

For linting and static analysis https://github.com/golangci/golangci-lint

For security checks I like to use gitleaks https://github.com/gitleaks/gitleaks and run it with the following checks.

```shell
gitleaks detect -v -c gitleaks.toml
gitleaks protect -v -c gitleaks.tom
```

Note that you need to include a gitleaks toml file. [Here is the one I use](https://gist.github.com/boyter/48942e5bcf9eed3cd3ed5f8ad413920f) as a base where I have included the vendor directory to be ignored as things like the AWS SDK causes gitleaks to freak out.

## Profiling

Profiling in Go has first class support. For CPU profiling you want your profiler to run either over a time period for things like HTTP services or when the program exits for short lived applications.

For short lived applications add the following in your main function,

```go
f, _ := os.Create("profile.pprof")
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()
```

This will start profiling when you run the command, and save the results to profile.pprof when the program exits.

For HTTP something like the following works,

```go
f, _ := os.Create("profile.pprof")
_ = pprof.StartCPUProfile(f)
go func() {
	time.Sleep(30 * time.Second)
	pprof.StopCPUProfile()
}()
```

Where it starts collecting CPU profile information for 30 seconds before saving it to disk. You can put this either inside your main function or behind a route, or even some sort of background task to collect profile information over time.

Memory profiles take a snapshot of the heap. I tend to use them mostly to get an idea of what is going on in long lived HTTP services.

```go
f, _ := os.Create("memprofile.pprof")
_ = pprof.WriteHeapProfile(f)
```

Putting the above behind a simple route dumps a snapshot of the heap to disk which I can then analyze.

In either case the analysis of the profile is the same,

```shell
go tool pprof -http=localhost:8090 profile.pprof
```

The above will open a http server on port 8090 which you can then inspect. This is usually how I inspect profile outputs since I find the HTTP interface easy to read and I really like using flame graphs. You can find more details on the [go.dev website for pprof](https://go.dev/blog/pprof).

## Unit Testing

To run all the unit tests for your code (with caching there is no reason to not run them all anymore) you should run the following which will run all the unit tests

```shell
go test ./...
```

To run benchmarks run the below inside the directory where the benchmark is. Say you have `./processor/` inside your project with a benchmark file inside there go to that directory and run,

```shell
go test --bench .
```

To run the built in fuzz tests,

```shell
go test -fuzz .
```

To create a test file you need only create a file with `_test` as a suffix in the name. For example to create a test for a file called `file.go` you might want to call the file `file_test.go`.

If you want to run an individual test you can do so,

```shell
go test ./... -run NameOfTest
```

Which will attempt to any test in all packages that have the name `NameOfTest`. Keep in mind that the argument `NameOfTest` supports regular expressions so its possible to target groups of tests assuming you name them well. For general running you can use `.` which matches everything.

If you find yourself wanting or needing to run tests ignoring the cache you can do the following,

```shell
GOCACHE=off go test ./...
```

The standard practice with Go tests is to put them next to the file you are testing. However this is not actually required. So long as you can import the code (that is it is made exposed with an uppercase prefix) you can put the tests anywhere you like. This of course means you cannot test the private code which some consider an anti-pattern anyway.

For fuzz testing I suggest reading this guide by [bitfield consulting](https://bitfieldconsulting.com/golang/fuzz-tests) which covers the use of the inbuilt fuzz detector well. Note that if you search for how to fuzz test in Go you will probably run into articles about the previous first choice https://github.com/dvyukov/go-fuzz so look for guides written after mid 2022.

### Mocks

Generally mocking in Go is as simple as defining an interface over the things you want to mock away. However some dislike the manual approach and use tools like [testify](<https://boyter.org/posts/how-to-start-go-project-2023/(https://github.com/stretchr/testify)>) and [mockery](https://github.com/vektra/mockery) to achieve this.

If you are coming from a Java background, don’t bother looking for a Mockito replacement. There isn’t anything even close to it in Go. If you feel like creating one please let me know though.

I fall into the manual approach generally so I have no strong feelings either way on the above. In short though, stick to “Accept interfaces, return structs” as your approach to code and you should be fine. You can read about this at the following links https://medium.com/swlh/golangs-interfaces-explained-with-mocks-886f69eca6f0 https://bryanftan.medium.com/accept-interfaces-return-structs-in-go-d4cab29a301b https://tutorialedge.net/golang/accept-interfaces-return-structs/

## Integration Testing

If you end up adding integration tests inside your Go code its common practice to split them via tags. This is where you put the following at the top of your test file

```go
//go:build integration

package mypackage
```

You can then run them

```shell
go test --tags=integration ./...
```

This will still run the untagged tests. You can also use this to split tests into separate groups. However you do need to be careful, because by default when each group runs they run in their own context, so methods in one test group will not be available to others and cause a compile error.

### Test Caching

Test results are cached by default which might not be ideal for integration tests. Where you want to override this `-count=1` can be added to your run command to run the test 1 time ignoring the cached results. You can replace 1 with a higher value if required.

```shell
go test -count=1 --tags=integration ./...
```

## Community

Your best bet to hang out with other “Gophers” is either the [subreddit](https://www.reddit.com/r/golang/) or [slack](https://gophers.slack.com/). Of the two I find the slack to be more accommodating and nicer to deal with.

Twitter accounts I find useful, although some might have moved to the fediverse, but you can confirm via their profile.

- https://twitter.com/go_perf
- https://twitter.com/golangnews
- https://twitter.com/golang_news
- https://twitter.com/golang
- https://twitter.com/golangweekly
- https://twitter.com/goinggodotnet
- https://twitter.com/_rsc
- https://twitter.com/bradfitz

The following newsletter is worth subscribing to as well https://golangweekly.com/ and is a great way to keep an eye on the latest developments.

The following websites/blogs tend to have quality Go content worth paying attention to

- https://bitfieldconsulting.com/golang/
- https://dave.cheney.net/
- https://www.ardanlabs.com/categories/go-programing/

## Multiple Main Entry Points

There are times where you want to potentially have multiple entry points into an application by having multiple `main.go` files in the main package. One way to achieve this is to have shared code in one repository, and then import it into others. However this can be cumbersome when you want to use vendor imports.

One common pattern for this is to have a directory inside the root of the application and place your main.go files in there. For example,

```shell
SRC
├── cmd
│   ├── commandline
│   │   └── main.go
│   ├── webhttp
│   │   └── main.go
│   ├── convert1.0-2.0
│   │   └── main.go
```

Then each entry point can import from the root package and you can compile and run multiple entry points into your application. Assuming your application lives in `http://github.com/name/mycode` you would need to import like so in each application,

```go
package main

import (
	"github.com/name/mycode"
)
```

With the above you can now call into code exposed by the repository package in the root.

## OS Specific Code

Occasionally you will require code in your application that will not compile or run on different operating systems. The most common way to deal with this is to have the following structure in your application,

```shell
main_darwin.go
main_linux.go
main_windows.go
```

Assuming that the above just contained definitions for line breaks on multiple operating systems EG `const LineBreak = "\n\r"` or `const LineBreak = "\n"` the you can import and refer to `LineBreak` however you wish. The same technique will work for functions or anything else you wish to include.

## Docker

Using the above techniques you can run inside Docker using multiple entry points easily. A sample dockerfile to achieve this is below using code from our hypothetical repository at `https://username@bitbucket.code.company-name.com.au/scm/code/random-code.git`

The below would build and run the main application,

```
FROM golang:1.20

COPY ./ /go/src/bitbucket.code.company-name.com.au/scm/code/
WORKDIR /go/src/bitbucket.code.company-name.com.au/scm/code/

RUN go build main.go

CMD ["./main"]
```

The below would build and run from the one of the alternate entry point’s for the application,

```shell
FROM golang:1.20

COPY ./ /go/src/bitbucket.code.company-name.com.au/scm/code/
WORKDIR /go/src/bitbucket.code.company-name.com.au/scm/code/cmd/webhttp/

RUN go build main.go

CMD ["./main"]
```

A few people who have read this post suggested using multi stage docker builds https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds which works well with Docker 17.05 or higher. More details here https://medium.com/travis-on-docker/multi-stage-docker-builds-for-creating-tiny-go-images-e0e1867efe5a An example would be,

```shell
FROM golang:1.20
COPY . /go/src/bitbucket.code.company-name.com.au/scm/code
WORKDIR /go/src/bitbucket.code.company-name.com.au/scm/code/
RUN CGO_ENABLED=0 go build main.go

FROM alpine:3.7
RUN apk add --no-cache ca-certificates
COPY --from=0 /go/src/bitbucket.code.company-name.com.au/scm/code/main .
CMD ["./main"]
```

The result is much smaller images to run your code which is always nice.

## Useful Tools/Packages

A brief list of useful tools I like related to Go development, and packages that I like to use. Note that some are not written in Go.

### Tools

- gow https://github.com/mitranim/gow - A watch mode command. Run it with your arguments, and it will hot recompile under the hood for you. Very useful for HTTP development. For example `gow -e=go,html,css run .` will watch for file changes to any Go, HTML or CSS file, and if found rerun the `go run .` command giving you a hot reload.
- hyperfine https://github.com/sharkdp/hyperfine - A command line benchmarking tool. Think of it as a replacement for running `time` multiple times and averaging the results.
- dcd https://github.com/boyter/dcd - Duplicate code detector. My own project (so I hesitate to add it) but you can run it to find examples of duplicated code in a project. Especially useful when looking to refactor.
- gotestsum https://github.com/gotestyourself/gotestsum - Alternate test runner. Gives different test outputs with format options that you might prefer. Can produce junit output format to work with CI/CD systems.
- https://mholt.github.io/json-to-go/ - JSON to Go generator. Goland can do this for you too, but this tool works pretty well for pasting in JSON and getting back a struct that can hold it.
- gofumpt https://github.com/mvdan/gofumpt - A stricter formatter than gofmt. I personally have not used this, but had it suggested to me.
- https://github.com/golangci/golangci-lint - Static type checker and lint enforcer. Apply this from day one of your project and it will save you a lot of cleanup. The suggestions it provides are always good and it helps when asked the usual sort of audit questions. Hook it into your CI/CD pipeline as a deployment gate for best results.
- gitleaks https://github.com/gitleaks/gitleaks - SAST tool to find and identify checked in secrets, passwords and such. Again works well to help pass audit questions.
- BFG Repo-Cleaner https://rtyley.github.io/bfg-repo-cleaner/ - Easiest way to remove large binaries or checked in secrets from a git repository. Very useful for fixing issues gitleaks finds.

### Packages

- https://github.com/tidwall/gjson - A quick way to get JSON values from a JSON document. Rather than deserialize into a struct you can just get the value you want. Especially useful for integration tests when running against your own endpoints. Since importing the struct you used to marshal out won’t catch regressions, using this is a decent way to write non brittle tests. `g := gjson.Get(resp.Body, "response.someValue")`.
- https://github.com/rs/zerolog/ - Structured JSON logs. A fast way to get structured logs that make sense. My preference is to use this with a unique code `date | md5 | cut -c 1-8` allowing you to track down errors to the exact line `log.Error().Str(common.UniqueCode, "9822f401").Err(err).Msg("")`. Add in context information to get the invoke details too giving you some level of observability through your logs.
- https://github.com/gorilla/mux - A replacement for the standard Go router which is a bit janky. Annoyingly this code is now archived. While the code has not rusted and it still works finding something that is maintained should be a priority. The following blog post has a decent over view of the potential replacements https://mariocarrion.com/2022/12/19/gorilla-mux-archived-migration-path.html
- https://github.com/google/go-cmp - Better than `reflect.DeepEqual` for equality checks.
- https://github.com/google/uuid/ - Probably the defacto package for creating the various versions of UUID’s.
