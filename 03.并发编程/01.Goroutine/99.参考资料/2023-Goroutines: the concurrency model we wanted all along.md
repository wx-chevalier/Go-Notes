> [原文地址](https://jayconrod.com/posts/128/goroutines-the-concurrency-model-we-wanted-all-along)

# Goroutines: the concurrency model we wanted all along

I often hear criticisms of Go that seem superficial: the variable names are too short, the functions are not overloadable, the exporting-symbols-by-capitalizing thing is weird, and until recently, generics are non-existent. Go is an idiosyncratic language, and of course I have my own opinions and my own set of gripes too, but today I want to highlight one of the features I like most about the language: _goroutines_.

我经常听到对 Go 的批评似乎是肤浅的：变量名太短，函数不可重载，按大写导出符号的事情很奇怪，直到最近，泛型还不存在。Go 是一种特殊的语言，当然我也有我自己的观点和自己的抱怨，但今天我想强调我最喜欢的语言之一：goroutines。

Goroutines are Go's main concurrency primitives. They look very much like threads, but they are cheap to create and manage. Go's runtime schedules goroutines onto real threads efficiently to avoid wasting resources, so you can easily create _lots_ of goroutines (like one goroutine per request), and you can write simple, imperative, blocking code. Consequently, Go networking code tends to be straightforward and easier to reason about than the equivalent code in other languages.

Goroutines 是 Go 的主要并发原语。它们看起来非常像线程，但它们的创建和管理成本很低。Go 的运行时将 goroutines 有效地调度到真正的线程上以避免浪费资源，因此您可以轻松创建大量 goroutine（例如每个请求一个 goroutine），并且可以编写简单的命令式阻塞代码。因此，Go 网络代码往往比其他语言中的等效代码更直接且更容易推理。

For me, goroutines are the single feature that distinguishes Go from other languages. They are why I prefer Go for writing code that requires any amount of parallelism.

对我来说，goroutines 是将 Go 与其他语言区分开来的唯一特性。这就是为什么我更喜欢 Go 来编写需要任意并行度的代码。

Before we talk more about goroutines, let's cover some history so it makes sense why you'd want them.

在我们更多地讨论 goroutines 之前，让我们介绍一些历史，以便理解为什么你想要它们。

## Forked and threaded servers 分叉和线程服务器

High performance servers need to handle requests from many clients at the same time. There are many ways to design a server to handle that.
高性能服务器需要同时处理来自多个客户端的请求。有很多方法可以设计服务器来处理这个问题。

The simplest naïve design is to have a main process that calls [`accept`](https://manpages.debian.org/bookworm/manpages-dev/accept.2.en.html) in a loop, then calls [`fork`](https://manpages.debian.org/bookworm/manpages-dev/fork.2.en.html) to create a child process that handles the request.
最简单的朴素设计是让一个主进程在循环中调用，然后调用 `accept` `fork` 以创建一个处理请求的子进程。

I learned about this pattern from the excellent [Beej's Guide to Network Programming](https://beej.us/guide/bgnet/html/) in the early 2000s when I was a student. `fork` is a nice pattern to use when learning network programming since you can focus on networking and not server architecture. It's hard to write an efficient server following this pattern though, and I don't think anyone uses it in practice anymore.
我在 2000 年代初还是学生的时候从优秀的 Beej's Guide to Network Programming 中了解到这种模式。 `fork` 是学习网络编程时使用的一种很好的模式，因为您可以专注于网络而不是服务器体系结构。但是，很难按照这种模式编写高效的服务器，而且我认为没有人在实践中使用它了。

There are a lot of problems with `fork`. First is cost: a `fork` call on Linux appears to be fast, but it marks all your memory as copy-on-write. Each write to a copy-on-write page causes a minor page fault, a small delay that's hard to measure. Context switching between processes is also expensive. Another problem is scale: it is difficult to coordinate use of shared resources like CPU, memory, database connections, and whatever else among a large number of child processes. If you get a surge of traffic and you create too many processes, they'll contend with each other for the CPU. If you limit the number of processes you create, then a large number of slow clients can block your service for everyone while your CPU sits idle. Careful use of timeouts helps (and is necessary regardless of server architecture).
有很多 `fork` 问题。首先是成本：在 Linux 上 `fork` 调用似乎很快，但它将所有内存标记为写入时复制。每次写入写入时复制页面都会导致轻微的页面错误，这是一个难以测量的小延迟。进程之间的上下文切换也很昂贵。另一个问题是规模：很难在大量子进程中协调共享资源的使用，如 CPU、内存、数据库连接和其他任何资源。如果您获得流量激增并且您创建了太多进程，它们将相互争夺 CPU。如果限制创建的进程数，则当 CPU 处于空闲状态时，大量慢速客户端可能会阻止每个人的服务。谨慎使用超时会有所帮助（无论服务器体系结构如何，都是必需的）。

These problems are somewhat mitigated by using threads instead of processes. A thread is cheaper to create than a process since it shares memory and most other resources. It's also relatively easy to communicate between threads in a shared address space, using semaphores and other constructs to manage shared resources.
通过使用线程而不是进程可以在一定程度上缓解这些问题。创建线程比创建线程更便宜，因为它共享内存和大多数其他资源。使用信号量和其他构造来管理共享资源，在共享地址空间中的线程之间进行通信也相对容易。

Threads do still have a substantial cost though, and you'll run into scaling problems if you create a new thread for every connection. As with processes, you need to limit the number of running threads to avoid heavy CPU contention, and you need to time out slow requests. Creating a new thread still takes time, though you can mitigate that by recycling threads across requests with a thread pool.
不过，线程仍然有相当大的成本，如果您为每个连接创建一个新线程，则会遇到扩展问题。与进程一样，您需要限制正在运行的线程数以避免严重的 CPU 争用，并且需要使慢速请求超时。创建新线程仍然需要时间，但您可以通过使用线程池跨请求回收线程来缓解这种情况。

Whether you're using processes or threads, you still have a question that's difficult to answer: how many should you create? If you allow an unlimited number of threads, clients can use up all your memory and CPU with a small surge in traffic. If you limit your server to a maximum number of threads, then a bunch of slow clients can clog up your server. Timeouts help, but it's still difficult use your hardware resources efficiently.
无论您使用的是进程还是线程，您仍然有一个难以回答的问题：您应该创建多少个？如果允许无限数量的线程，则客户端可能会用完所有内存和 CPU，流量会略有激增。如果将服务器限制为最大线程数，则一堆慢速客户端可能会阻塞服务器。超时会有所帮助，但仍然难以有效利用硬件资源。

## Event-driven servers 事件驱动型服务器

Since we can't easily predict how many threads we'll need, what happens when we try to decouple requests from threads? What if we have just one thread dedicated to application logic (or perhaps a small, fixed number of threads), and we handle all the network traffic in the background using asynchronous system calls? This is an _event-driven server architecture_.
由于我们无法轻松预测需要多少线程，因此当我们尝试将请求与线程分离时会发生什么？如果我们只有一个线程专用于应用程序逻辑（或者可能是少量固定数量的线程），并且我们使用异步系统调用在后台处理所有网络流量，该怎么办？这是一个事件驱动的服务器体系结构。

Event-driven servers were designed around the [`select`](https://manpages.debian.org/unstable/manpages-dev/select.2.en.html) system call. (Later mechanisms like [`poll`](https://manpages.debian.org/bookworm/manpages-dev/poll.2.en.html) have replaced `select`, but `select` is widely known, and they all serve the same conceptual purpose here.) `select` accepts a list of file descriptors (generally sockets) and returns which ones are ready to read or write. If none of the file descriptors are ready, `select` blocks until at least one is.
事件驱动的服务器是围绕 `select` 系统调用设计的。（后来的机制已经 `poll` 取代 `select` 了，但 `select` 广为人知，它们在这里都有相同的概念目的。 `select` 接受文件描述符（通常是套接字）的列表，并返回哪些描述符已准备好读取或写入。如果文件描述符都未准备就绪， `select` 则阻止，直到至少有一个文件描述符准备就绪。

```c
#include <sys/select.h>

int select(int nfds, fd_set *restrict readfds,
           fd_set *restrict writefds, fd_set *restrict exceptfds,
           struct timeval *restrict timeout);
#include <poll.h>

int poll(struct pollfd *fds, nfds_t nfds, int timeout);
```

(I want to elaborate a bit on what it means for a socket to be "ready" because I had some trouble understanding this initially. Each socket has a kernel buffer for receiving and a buffer for sending. When your computer receives a packet, the kernel stores its data in the receive buffer for the appropriate socket until your program gets it with `recv`. Likewise, when your program calls `send`, the kernel stores the data in the socket's send buffer until it can be transmitted. A socket is ready to receive if there's data in the receive buffer. It's ready to send if there's available space in the send buffer. If the socket is not ready, `recv` and `send` block until it is, by default.)
（我想详细说明套接字“准备就绪”的含义，因为我最初在理解这一点时遇到了一些麻烦。每个套接字都有一个用于接收的内核缓冲区和一个用于发送的缓冲区。当您的计算机收到数据包时，内核会将其数据存储在相应套接字的接收缓冲区中，直到您的程序使用 `recv` .同样，当你的程序调用 `send` 时，内核将数据存储在套接字的发送缓冲区中，直到可以传输。如果接收缓冲区中有数据，套接字已准备好接收。如果发送缓冲区中有可用空间，则已准备好发送。默认情况下，如果套接字未就绪， `recv` 则 `send` 阻止直到准备就绪。

To implement an event-driven server, you track a socket and some state for each request that's blocked on the network. The server has a single main event loop where it calls `select` with all those blocked sockets. When `select` returns, the server knows which requests can make progress, so for each request, it calls into the application logic with the stored state. When the application needs to use the network again, it adds the socket back to the "blocked" pool together with new state. The _state_ here can be anything the application needs to resume what it was doing: a closure to be called back, or a promise to be completed.
若要实现事件驱动服务器，请跟踪网络上阻止的每个请求的套接字和一些状态。服务器有一个主事件循环，它使用所有这些阻塞的套接字进行调用 `select` 。返回时 `select` ，服务器知道哪些请求可以取得进展，因此对于每个请求，它会调用具有存储状态的应用程序逻辑。当应用程序需要再次使用网络时，它会将套接字与新状态一起添加回“阻塞”池。此处的状态可以是应用程序恢复其正在执行的操作所需的任何内容：要回调的闭包或要完成的承诺。

This can all technically be implemented with a single thread. I can't speak to the details of any particular implementation, but languages that lack threading like JavaScript follow this model pretty closely. [Node.js describes itself](https://nodejs.org/en/about) as "an event-driven JavaScript runtime ... designed to build scalable network applications." This is exactly what they mean.
从技术上讲，这一切都可以通过单个线程实现。我不能谈论任何特定实现的细节，但是像 JavaScript 这样缺乏线程的语言非常紧密地遵循这个模型。Node.js 将自己描述为“一个事件驱动的 JavaScript 运行时......旨在构建可扩展的网络应用程序。这正是他们的意思。

Event-driven servers can generally make better use of CPU and memory than purely fork- or thread-based servers. You can spawn an application thread per core to handle requests in parallel. The threads don't contend with each other for CPU since the number of threads equals the number of cores. The threads are never idle when there are requests that can make progress. Efficient. So efficient that it's hard to justify writing a server any other way these days.
事件驱动的服务器通常可以比纯粹基于分支或线程的服务器更好地利用 CPU 和内存。您可以为每个内核生成一个应用程序线程以并行处理请求。线程不会相互争用 CPU，因为线程数等于内核数。当有可以取得进展的请求时，线程永远不会空闲。有效。如此高效，以至于如今很难证明以任何其他方式编写服务器的合理性。

It sounds like a good idea on paper anyway, but it's a nightmare to write application code that works like this. The specific way in which it's a nightmare depends on the language and framework you're using. In JavaScript, asynchronous functions typically return a `Promise`, to which you attach callbacks. In Java gRPC, you're dealing with `StreamObserver`. If you're not careful, you end up with lots of deeply nested "arrow code" functions. If you _are_ careful, you split up your functions and classes, obfuscating your control flow. Either way, you're in callback hell.
无论如何，这在纸面上听起来是个好主意，但编写像这样工作的应用程序代码是一场噩梦。它成为噩梦的具体方式取决于您使用的语言和框架。在 JavaScript 中，异步函数通常返回 `Promise` ，您可以向其附加回调。在 Java gRPC 中，您正在处理 `StreamObserver` .如果你不小心，你最终会得到很多深度嵌套的“箭头代码”函数。如果您小心，则会拆分函数和类，从而混淆控制流。无论哪种方式，您都处于回调地狱中。

To show what I mean, below is an example from the [Java gRPC Basics Tutorial](https://grpc.io/docs/languages/java/basics/#bidirectional-streaming-rpc). Does the control flow here make sense?
为了说明我的意思，下面是 Java gRPC 基础教程中的一个示例。这里的控制流有意义吗？

```java
public void routeChat() throws Exception {
  info("*** RoutChat");
  final CountDownLatch finishLatch = new CountDownLatch(1);
  StreamObserver<RouteNote> requestObserver =
      asyncStub.routeChat(new StreamObserver<RouteNote>() {
        @Override
        public void onNext(RouteNote note) {
          info("Got message \"{0}\" at {1}, {2}", note.getMessage(), note.getLocation()
              .getLatitude(), note.getLocation().getLongitude());
        }

        @Override
        public void onError(Throwable t) {
          Status status = Status.fromThrowable(t);
          logger.log(Level.WARNING, "RouteChat Failed: {0}", status);
          finishLatch.countDown();
        }

        @Override
        public void onCompleted() {
          info("Finished RouteChat");
          finishLatch.countDown();
        }
      });

  try {
    RouteNote[] requests =
        {newNote("First message", 0, 0), newNote("Second message", 0, 1),
            newNote("Third message", 1, 0), newNote("Fourth message", 1, 1)};

    for (RouteNote request : requests) {
      info("Sending message \"{0}\" at {1}, {2}", request.getMessage(), request.getLocation()
          .getLatitude(), request.getLocation().getLongitude());
      requestObserver.onNext(request);
    }
  } catch (RuntimeException e) {
    // Cancel RPC
    requestObserver.onError(e);
    throw e;
  }
  // Mark the end of requests
  requestObserver.onCompleted();

  // Receiving happens asynchronously
  finishLatch.await(1, TimeUnit.MINUTES);
}
```

This is from the _beginner_ tutorial, and it's not a complete example either. The sending code is synchronous while the receiving code is asynchronous.
这是来自初学者教程，也不是一个完整的示例。发送代码是同步的，而接收代码是异步的。

In Java, you may be dealing different with asynchronous types for your HTTP server, gRPC, SQL database, cloud SDK and whatever else, and you need adapters between all of them. It gets to be a mess very quickly. Locks are dangerous, too. You need to be careful about holding locks across network calls. It's also easy to make a mistake with locking and callbacks. For example, if a `synchronized` method calls a function that returns a `ListenableFuture` then attaches an inline callback, the callback also needs a `synchronized` block, even though it's nested inside the parent method.
在 Java 中，您可能会处理 HTTP 服务器、gRPC、SQL 数据库、云 SDK 和其他任何异步类型的不同类型，并且您需要在所有这些类型之间使用适配器。它很快就会变得一团糟。锁也很危险。您需要小心跨网络调用保持锁定。锁定和回调也很容易出错。例如，如果方法调用返回 然后 `ListenableFuture` 附加内联回调的函数，则回调也需要一个 `synchronized` 块，即使它嵌套在父 `synchronized` 方法中也是如此。

## goroutines goroutines

What was this post about again? Oh yes, goroutines.
这篇文章又是关于什么的？哦，是的，去例程。

A _goroutine_ is Go's version of a thread. Like threads in other languages, each goroutine has its own stack. Goroutines may execute in parallel, concurrently with other goroutines. Unlike threads, goroutines are very cheap to create: they aren't bound to an OS thread, and their stacks start out very small (2 KiB) but can grow as needed. When you create a goroutine, you're essentially allocating a closure and adding it to a queue in the runtime.
goroutine 是 Go 的线程版本。与其他语言中的线程一样，每个 goroutine 都有自己的堆栈。Goroutines 可以并行执行，与其他 goroutines 同时执行。与线程不同，goroutines 的创建成本非常低：它们不绑定到操作系统线程，它们的堆栈开始时非常小（2 KiB），但可以根据需要进行增长。当你创建一个 goroutine 时，你实际上是在分配一个闭包并将其添加到运行时的队列中。

Internally, Go's runtime has a set of OS threads that execute goroutines (normally one thread per core). When a thread is available and a goroutine is ready to run, the runtime schedules the goroutine onto the thread, executing application logic. If a goroutine blocks on something like a mutex or channel, the runtime adds it to a set of blocked goroutines then schedules the next ready goroutine onto the same OS thread. _This also applies to the network_: when a goroutine sends or receives data on a socket that's not ready, it yields its OS thread to the scheduler.
在内部，Go 的运行时有一组执行 goroutines 的操作系统线程（通常每个内核一个线程）。当线程可用且 goroutine 准备好运行时，运行时会将 goroutine 调度到线程上，执行应用程序逻辑。如果 goroutine 阻塞了互斥锁或通道之类的东西，运行时会将其添加到一组被阻止的 goroutine，然后将下一个准备好的 goroutine 调度到同一个操作系统线程上。这也适用于网络：当 goroutine 在未准备好的套接字上发送或接收数据时，它会将其操作系统线程交给调度程序。

Sound familiar? Go's scheduler acts a lot like the main loop in an event-driven server. Except instead of relying solely on `select` and focusing on file descriptors, the scheduler handles everything in the language that might block. You no longer need to avoid blocking calls because the scheduler makes efficient use of the CPU either way. You're free to spawn lots of goroutines (one per request!) because they're cheap to create, and the threads don't contend for the CPU (minimal context switching). You don't need to worry about thread pools and executor services because the runtime effectively has one big thread pool.
听起来很耳熟？Go 的调度器的行为很像事件驱动服务器中的主循环。除了不完全依赖 `select` 和关注文件描述符之外，调度程序以可能阻止的语言处理所有内容。您不再需要避免阻塞调用，因为无论哪种方式，调度程序都会有效地使用 CPU。你可以自由地生成大量的 goroutines（每个请求一个！），因为它们的创建成本很低，而且线程不会争用 CPU（最小的上下文切换）。您无需担心线程池和执行器服务，因为运行时实际上有一个大线程池。

In short, you can write simple blocking application code in a clean imperative style as if you were writing a thread-based server, but you keep all the efficiency advantages of an event-driven server. Best of both worlds. This kind of code composes well across frameworks. You don't need adapters between your `StreamObservers` and `ListenableFutures`.
简而言之，您可以像编写基于线程的服务器一样，以干净的命令式风格编写简单的阻塞应用程序代码，但您可以保留事件驱动服务器的所有效率优势。两全其美。这种代码可以跨框架很好地组合。您 `StreamObservers` 不需要 和 `ListenableFutures` 之间的适配器。

Let's take a look at the same example from the [Go gRPC Basics Tutorial](https://grpc.io/docs/languages/go/basics/#bidirectional-streaming-rpc-1). I find the control flow here easier to comprehend than the Java example because both the sending and receiving code are synchronous. In both goroutines, we're able to call `stream.Recv` and `stream.Send` in a `for` loop. No need for callbacks, subclasses, or executors.
让我们看一下 Go gRPC 基础知识教程中的同一示例。我发现这里的控制流比 Java 示例更容易理解，因为发送和接收代码都是同步的。在这两个 goroutines 中，我们都能够调用 `stream.Recv` 和 `stream.Send` `for` 循环。不需要回调、子类或执行器。

```go
stream, err := client.RouteChat(context.Background())
waitc := make(chan struct{})
go func() {
  for {
    in, err := stream.Recv()
    if err == io.EOF {
      // read done.
      close(waitc)
      return
    }
    if err != nil {
      log.Fatalf("Failed to receive a note : %v", err)
    }
    log.Printf("Got message %s at point(%d, %d)", in.Message, in.Location.Latitude, in.Location.Longitude)
  }
}()
for _, note := range notes {
  if err := stream.Send(note); err != nil {
    log.Fatalf("Failed to send a note: %v", err)
  }
}
stream.CloseSend()
<-waitc
```

## `async` / `await`, virtual threads `async` / `await` 、虚拟线程

Go is a fairly new language, and it was designed at a time when event-driven servers were already popular and well-understood. Java, C++, Python, and other languages are older and don't have that same luxury. So as a language designer, what features can you add to make writing asynchronous application code less painful?
Go 是一种相当新的语言，它是在事件驱动服务器已经流行和被充分理解的时候设计的。Java，C++，Python 和其他语言较旧，没有同样的奢侈。因此，作为语言设计师，您可以添加哪些功能来减轻编写异步应用程序代码的痛苦？

[`async` / `await`](https://en.wikipedia.org/wiki/Async/await) has become the main solution in most languages. Wikipedia tells me that it was first added to F# in 2007 (around the same time Go was being developed). It's now in C#, C++, JavaScript, Python, Rust, Swift, and a bunch of other languages.
`async` / `await` 已成为大多数语言的主要解决方案。维基百科告诉我，它在 2007 年首次添加到 F# 中（大约在开发 Go 的同时）。它现在有 C#，C++，JavaScript，Python，Rust，Swift 和许多其他语言。

`async` / `await` let you write code that works with asynchronous functions in a style that resembles imperative blocking code. An asynchronous function is marked with the `async` keyword and returns a promise (or whatever the language equivalent is). When you call an asynchronous function, you can "block" and get the value in the promise with the `await` keyword. If a function uses `await`, it must also be marked `async`, and the compiler rewrites it to return a promise. There are different ways to implement `async` and `await`, but the feature generally implies a language has support for cooperatively scheduled coroutines if not full support for threads. When a coroutine `awaits` a promise that's not completed yet, it yields to the language's runtime, which may schedule another ready coroutine on the same thread.
`async` / `await` 允许您编写与异步函数一起使用的代码，其样式类似于命令性阻塞代码。异步函数用 `async` 关键字标记并返回一个 promise（或任何等效语言）。调用异步函数时，可以使用 `await` 关键字“阻止”并获取 promise 中的值。如果一个函数使用 `await` ，它也必须被标记 `async` ，并且编译器重写它以返回一个承诺。有不同的 `await` 实现 `async` 方式和，但该功能通常意味着一种语言支持协作调度协程，如果不是完全支持线程。当协程是一个尚未完成的承诺时，它会屈服于语言的运行时，这可能会在同一线程上安排另一个就绪的协程 `awaits` 。

I'm pleased that so many languages have adopted `async` / `await`. Since I've worked primarily in Go, Java, and C++, I haven't personally had much opportunity to use it, but whenever I need to do something in JavaScript, I'm really glad it's there. `await` control flow is much easier to understand than asynchronous callback code. It is a bit annoying though to annotate functions with `async`. Bob Nystrom's [What Color is Your Function](https://journal.stuffwithstuff.com/2015/02/01/what-color-is-your-function/) is a humorous criticism of that.
我很高兴有这么多语言采用了 `async` / `await` 。由于我主要从事 Go、Java 和 C++ 的工作，所以我个人没有太多机会使用它，但每当我需要用 JavaScript 做一些事情时，我真的很高兴它在那里。 `await` 控制流比异步回调代码更容易理解。虽然用 注释 `async` 函数有点烦人。鲍勃·奈斯特罗姆（Bob Nystrom）的《你的功能是什么颜色》（What Color is Your Function）对此进行了幽默的批评。

Java is conspicuously absent from the list of languages above. Until now, you've had to either spawn an unreasonable number of threads or deal with Java's particular circle of callback hell. Happily, [JEP 444](https://openjdk.org/jeps/444) adds [virtual threads](https://blog.rockthejvm.com/ultimate-guide-to-java-virtual-threads/), which sound a lot like goroutines. Virtual threads are cheap to create. The JVM schedules them onto _platform threads_ (real threads the kernel knows about). There are a fixed number of platform threads, generally one per core. When a virtual thread performs a blocking operation, it releases its platform thread, and the JVM may schedule another virtual thread onto it. Unlike goroutines, virtual thread scheduling is cooperative: a virtual thread doesn't yield to the scheduler until it performs a blocking operation. This means that a tight loop can hold a thread indefinitely. I don't know whether this is an implementation limitation or if there's a deeper issue. Go used to have this problem until fully preemptive scheduling was implemented in 1.14.
Java 在上面的语言列表中明显不存在。到目前为止，你要么生成不合理的线程数，要么处理 Java 的特定回调地狱圈。令人高兴的是，JEP 444 添加了虚拟线程，这听起来很像 goroutines。创建虚拟线程的成本很低。JVM 将它们调度到平台线程（内核知道的真实线程）上。平台线程的数量是固定的，通常每个内核一个。当虚拟线程执行阻塞操作时，它会释放其平台线程，JVM 可能会将另一个虚拟线程调度到该线程上。与 goroutines 不同，虚拟线程调度是协作的：虚拟线程在执行阻塞操作之前不会屈服于调度程序。这意味着紧密的环可以无限期地保持线程。我不知道这是实现限制还是存在更深层次的问题。Go 曾经有这个问题，直到 1.14 中实现了完全抢占式调度。

Virtual threads are available for preview in the current release and are expected to become stable in JDK 21 (due September 2023). I'm looking forward to deleting a lot of `ListenableFutures`, but it's unclear to me how virtual threads will interact with existing frameworks designed around a more purely event-driven model. Whenever a new language or runtime feature is introduced, there's a long cultural migration period, and I think the Java ecosystem is pretty conservative in that regard. Still, I'm encouraged by the enthusiastic adoption of `async` / `await` in other languages. I hope Java will be similar with virtual threads.
虚拟线程在当前版本中提供预览版，预计将在 JDK 21（2023 年 9 月到期）中变得稳定。我期待删除很多 `ListenableFutures` ，但我不清楚虚拟线程将如何与围绕更纯粹的事件驱动模型设计的现有框架进行交互。每当引入新的语言或运行时功能时，都会有很长的文化迁移期，我认为 Java 生态系统在这方面非常保守。尽管如此，我对其他语言 `await` 的 `async` 热情采用感到鼓舞。我希望 Java 与虚拟线程类似。
