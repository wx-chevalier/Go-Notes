# How I write HTTP services in Go after 13 years

Nearly six years ago I wrote a blog post outlining [how I write HTTP services in Go](https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html), and I’m here to tell you, once again, how I write HTTP services.

That original post went a little viral and sparked some great discussions that have influenced how I do things today. And after years of hosting the [Go Time podcast](https://changelog.com/gotime), discussing Go on [X/Twitter](https://twitter.com/matryer), and gaining more experience maintaining code like this over years, I thought it was time for a refresh.

(And for those pedants who notice Go isn’t exactly 13 years old, I started writing HTTP services in Go [version .r59](https://go.dev/doc/devel/pre_go1#r59).)

This post covers a range of topics related to building services in Go, including:

- Structuring servers and handlers for maximum maintainability
- Tips and tricks for optimizing for a quick startup and graceful shutdown
- How to handle common work that applies to many types of requests
- Going deep on properly testing your services

From small projects to large, these practices have stood the test of time for me, and I hope they will for you too.

## Who is this post for?

This post is for you. It’s for everybody who plans to write some kind of HTTP service in Go. You may also find this useful if you’re learning Go, as lots of the examples follow good practices. Experienced gophers might also pick up some nice patterns.

To find this post most useful, you’ll need to know the basics of Go. If you don’t feel like you’re quite there yet, I cannot recommend [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests/) by Chris James enough. And if you’d like to hear more from Chris, you can check out the episode of Go Time we did with Ben Johnson on [The files and folders of Go projects](https://changelog.com/gotime/278).

If you’re familiar with the previous versions of this post, this section contains a quick summary of what’s different now. If you’d like to start from the beginning, skip to the next section.

1. My handlers used to be methods hanging off a server struct, but I no longer do this. If a handler function wants a dependency, it can bloody well ask for it as an argument. No more surprise dependencies when you’re just trying to test a single handler.
2. I used to prefer `http.HandlerFunc` over `http.Handler` — enough third-party libraries think about `http.Handler` first that it makes sense to embrace that. `http.HandlerFunc` is still extremely useful, but now most things are represented as the interface type. It doesn’t make much difference either way.
3. I’ve added more about testing including some Opinions™.
4. I’ve added more sections, so a full read through is recommended for everybody.

## The `NewServer` constructor

Let’s start by looking at the backbone of any Go service: the server. The `NewServer` function makes the main `http.Handler`. Usually I have one per service, and I rely on HTTP routes to divert traffic to the right handlers within each service because:

- `NewServer` is a big constructor that takes in all dependencies as arguments
- It returns an `http.Handler` if possible, which can be a dedicated type for more complex situations
- It usually configures its own muxer and calls out to `routes.go`

For example, your code might look similar to this:

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

Expand code

```go
func NewServer(
	logger *Logger
	config *Config
	commentStore *commentStore
	anotherStore *anotherStore
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(
		mux,
		Logger,
		Config,
		commentStore,
		anotherStore,
	)
	var handler http.Handler = mux
	handler = someMiddleware(handler)
	handler = someMiddleware2(handler)
	handler = someMiddleware3(handler)
	return handler
}
```

In test cases that don’t need all of the dependencies, I pass in `nil` as a signal that it won’t be used.

The `NewServer` constructor is responsible for all the top-level HTTP stuff that applies to all endpoints, like CORS, auth middleware, and logging:

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
var handler http.Handler = mux
handler = logging.NewLoggingMiddleware(logger, handler)
handler = logging.NewGoogleTraceIDMiddleware(logger, handler)
handler = checkAuthHeaders(handler)
return handler
```

Setting up the server is usually a case of exposing it using Go’s built-in `http` package:

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

Expand code

```go
srv := NewServer(
	logger,
	config,
	tenantsStore,
	slackLinkStore,
	msteamsLinkStore,
	proxy,
)
httpServer := &http.Server{
	Addr:    net.JoinHostPort(config.Host, config.Port),
	Handler: srv,
}
go func() {
	log.Printf("listening on %s\n", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
	}
}()
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	<-ctx.Done()
	if err := httpServer.Shutdown(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
	}
}()
wg.Wait()
return nil
```

### Long argument lists

There must be a limit at which point it stops being the right thing to do, but most of the time I am happy adding lists of dependencies as arguments. And while they do sometimes get quite long, I find it’s still worth it.

Yes, it saves me from making a struct, but the real benefit is that I get slightly more type safety from arguments. I can make a struct skipping any fields that I don’t like, but a function forces my hand. I have to look up fields to know how to set them in a struct, whereas I can’t call a function if I don’t pass the right arguments.

It’s not so bad if you format it as a vertical list, like I’ve seen in modern frontend code:

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
srv := NewServer(
	logger,
	config,
	tenantsStore,
	commentsStore,
	conversationService,
	chatGPTService,
)
```

## Map the entire API surface in `routes.go`

This file is the one place in your service where all routes are listed.

Sometimes you can’t help but have things spread around a bit, but it’s very helpful to be able to go to one file in every project to see its API surface.

Because of the big dependency argument lists in the `NewServer` constructor, you usually end up with the same list in your routes function. But again, it’s not so bad. And again, you soon know if you forgot something or got the order wrong thanks to Go’s type checking.

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
func addRoutes(
	mux                 *http.ServeMux,
	logger              *logging.Logger,
	config              Config,
	tenantsStore        *TenantsStore,
	commentsStore       *CommentsStore,
	conversationService *ConversationService,
	chatGPTService      *ChatGPTService,
	authProxy           *authProxy
) {
	mux.Handle("/api/v1/", handleTenantsGet(logger, tenantsStore))
	mux.Handle("/oauth2/", handleOAuth2Proxy(logger, authProxy))
	mux.HandleFunc("/healthz", handleHealthzPlease(logger))
	mux.Handle("/", http.NotFoundHandler())
}
```

In my example, `addRoutes` doesn’t return an error. Anything that can throw an error is moved to the `run` function and sorted out before it gets to this point leaving this function free to remain simple and flat. Of course, if any of your handlers do return errors for whatever reason, then fine, this can return an error too.

## `func main()` only calls `run()`

The `run` function is like the `main` function, except that it takes in operating system fundamentals as arguments, and returns, you guessed it, an error.

I wish `func main()` was `func main() error`. Or like in C where you can return the exit code: `func main() int`. By having an ultra simple main function, you too can have your dreams come true:

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
func run(ctx context.Context, w io.Writer, args []string) error {
	// ...
}

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
```

The code above creates a context, which is cancelled by `Ctrl+C` or equivalent, and calls down to the `run` function. If `run` returns `nil`, the function exits normally. If it returns an error, we write it to stderr and exit with a non-zero code. If I’m writing a command line tool where exit codes matter, I would return an int as well so I could write tests to assert the right one was returned.

Operating system fundamentals are passed into run as arguments. For example, you might pass in `os.Args` if it has flag support, and even `os.Stdin`, `os.Stdout`, `os.Stderr` dependencies. This makes your programs much easier to test because test code can call run to execute your program, controlling arguments, and all streams, just by passing different arguments.

The following table shows examples of input arguments to the run function:

| **Value**   | **Type**                 | **Description**                                                                        |
| :---------- | :----------------------- | :------------------------------------------------------------------------------------- |
| `os.Args`   | `[]string`               | The arguments passed in when executing your program. It’s also used for parsing flags. |
| `os.Stdin`  | `io.Reader`              | For reading input                                                                      |
| `os.Stdout` | `io.Writer`              | For writing output                                                                     |
| `os.Stderr` | `io.Writer`              | For writing error logs                                                                 |
| `os.Getenv` | `func(string) string`    | For reading environment variables                                                      |
| `os.Getwd`  | `func() (string, error)` | Get the working directory                                                              |

If you keep away from any global scope data, you can usually use `t.Parallel()` in more places, to speed up your test suites. Everything is self-contained, so multiple calls to `run` don’t interfere with each other.

I often end up with `run` function signatures that look like this:

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
func run(
	ctx    context.Context,
	args   []string,
	getenv func(string) string,
	stdin  io.Reader,
	stdout, stderr io.Writer,
) error
```

Now that we’re inside the `run` function, we can go back to writing normal Go code where we can return errors like it’s nobody’s business. We gophers just love returning errors, and the sooner we admit that to ourselves, the sooner those people on the internet can win and go away.

### Gracefully shutting down

If you’re running lots of tests, it’s important for your program to stop when each one is finished. (Alternatively, you might decide to keep one instance running for all tests, but that’s up to you.)

The context is passed through. It gets cancelled if a termination signal comes into the program, so it’s important to respect it at every level. At the very least, pass it to your dependencies. At best, check the `Err()` method in any long-running or loopy code, and if it returns an error, stop what you’re doing and return it up. This will help the server to gracefully shut down. If you kick off other goroutines, you can also use the context to decide if it’s time to stop them or not.

### Controlling the environment

The `args` and `getenv` parameters give us a couple of ways to control how our program behaves through flags and environment variables. Flags are processed using the args (as long as you don’t use the global space version of flags, and instead use `flags.NewFlagSet` inside `run`) so we can call run with different values:

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
args := []string{
	"myapp",
	"--out", outFile,
	"--fmt", "markdown",
}
go run(ctx, args, etc.)
```

If your program uses environment variables over flags (or even both) then the `getenv` function allows you to plug in different values without changing the actual environment.

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
getenv := func(key string) string {
	switch key {
	case "MYAPP_FORMAT":
		return "markdown"
	case "MYAPP_TIMEOUT":
		return "5s"
	default:
		return ""
}
go run(ctx, args, getenv)
```

For me, using this `getenv` technique beats using `t.SetEnv` for controlling environment variables because you can continue to run your tests in parallel by calling `t.Parallel()`, which `t.SetEnv` doesn’t allow.

This technique is even more useful when writing command line tools, because you often want to run the program with different settings to test all of its behavior.

In the `main` function, we can pass in the real things:

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	if err := run(ctx, os.Getenv, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
```

## Maker funcs return the handler

My handler functions don’t implement `http.Handler` or `http.HandlerFunc` directly, they return them. Specifically, they return `http.Handler` types.

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
// handleSomething handles one of those web requests
// that you hear so much about.
func handleSomething(logger *Logger) http.Handler {
	thing := prepareThing()
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// use thing to handle request
			logger.Info(r.Context(), "msg", "handleSomething")
		}
	)
}
```

This pattern gives each handler its own closure environment. You can do initialization work in this space, and the data will be available to the handlers when they are called.

Be sure to only read the shared data. If handlers modify anything, you’ll need a mutex or something to protect it.

Storing program state here is not usually what you want. In most cloud environments, you can’t trust that code will continue running over long periods of time. Depending on your production environment, servers will often shut down to save resources, or even just crash for other reasons. There may also be many instances of your service running with requests load balanced across them in unpredictable ways. In this case, an instance would only have access to its own local data. So it’s better to use a database or some other storage API to persist data in real projects.

## Handle decoding/encoding in one place

Every service will need to decode the request bodies and encode response bodies. This is a sensible abstraction that stands the test of time.

I usually have a pair of helper functions called encode and decode. An example version using generics shows you that you really are just wrapping a few basic lines, which I wouldn’t usually do, however this becomes useful when you need to make changes here for all of your APIs. (For example, say you get a new boss stuck in the 1990s and they want to add XML support.)

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
func encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}
```

Interestingly, the compiler is able to infer the type from the argument, so you don’t need to pass it when calling encode:

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
err := encode(w, r, http.StatusOK, obj)
```

But since it is a return argument in decode, you will need to specify the type you expect:

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
decoded, err := decode[CreateSomethingRequest](r)
```

I try not to overload these functions, but in the past I was quite pleased with a simple validation interface that fit nicely into the decode function.

## Validating data

I like a simple interface. Love them, actually. Single method interfaces are so easy to implement. So when it comes to validating objects, I like to do this:

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
// Validator is an object that can be validated.
type Validator interface {
	// Valid checks the object and returns any
	// problems. If len(problems) == 0 then
	// the object is valid.
	Valid(ctx context.Context) (problems map[string]string)
}
```

The `Valid` method takes a context (which is optional but has been useful for me in the past) and returns a map. If there is a problem with a field, its name is used as the key, and a human-readable explanation of the issue is set as the value.

The method can do whatever it needs to validate the fields of the struct. For example, it can check to make sure:

- Required fields are not empty
- Strings with a specific format (like email) are correct
- Numbers are within an acceptable range

If you need to do anything more complicated, like check the field in a database, that should happen elsewhere; it’s probably too important to be considered a quick validation check, and you wouldn’t expect to find that kind of thing in a function like this, so it could easily end up being hidden away.

I then use a type assertion to see if the object implements the interface. Or, in the generic world, I might choose to be more explicit about what’s going on by changing the decode method to insist on that interface being implemented.

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
func decodeValid[T Validator](r *http.Request) (T, map[string]string, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, nil, fmt.Errorf("decode json: %w", err)
	}
	if problems := v.Valid(r.Context()); len(problems) > 0 {
		return v, problems, fmt.Errorf("invalid %T: %d problems", v, len(problems))
	}
	return v, nil, nil
}
```

In this code, `T` has to implement the `Validator` interface, and the `Valid` method must return zero problems in order for the object to be considered successfully decoded.

It’s safe to return `nil` for problems because we are going to check `len(problems)`, which will be `0` for a `nil` map, but which won’t panic.

## The adapter pattern for middleware

Middleware functions take an `http.Handler` and return a new one that can run code before and/or after calling the original handler — or it can decide not to call the original handler at all.

An example is a check to make sure the user is an administrator:

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
func (s *server) adminOnly(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !currentUser(r).IsAdmin {
			http.NotFound(w, r)
			return
		}
		h(w, r)
	})
}
```

The logic inside the handler can optionally decide whether to call the original handler or not. In the example above, if `IsAdmin` is false, the handler will return an `HTTP 404 Not Found` and return (or abort); notice that the `h` handler is not called. If `IsAdmin` is true, the user is allowed to access the route, and so execution is passed to the h handler.

Usually I have middleware listed in the `routes.go` file:

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
package app

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/", s.handleAPI())
	mux.HandleFunc("/about", s.handleAbout())
	mux.HandleFunc("/", s.handleIndex())
	mux.HandleFunc("/admin", s.adminOnly(s.handleAdminIndex()))
}
```

This makes it very clear, just by looking at the map of endpoints, which middleware is applied to which routes. If the lists start getting bigger, try splitting them across many lines — I know, I know, but you get used to it.

## Sometimes I return the middleware

The above approach is great for simple cases, but if the middleware needs lots of dependencies (a logger, a database, some API clients, a byte array containing the data for “Never Gonna Give You Up” for a later prank), then I have been known to have a function that returns the middleware function.

The problem is, you end up with code that looks like this:

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
mux.Handle("/route1", middleware(logger, db, slackClient, rroll []byte, handleSomething(handlerSpecificDeps))
mux.Handle("/route2", middleware(logger, db, slackClient, rroll []byte, handleSomething2(handlerSpecificDeps))
mux.Handle("/route3", middleware(logger, db, slackClient, rroll []byte, handleSomething3(handlerSpecificDeps))
mux.Handle("/route4", middleware(logger, db, slackClient, rroll []byte, handleSomething4(handlerSpecificDeps))
```

This bloats out the code and doesn’t really provide anything useful. Instead, I would have the middleware function take the dependencies, but return a function that takes only the next handler.

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
func newMiddleware(
	logger Logger,
	db *DB,
	slackClient *slack.Client,
	rroll []byte,
) func(h http.Handler) http.Handler
```

The return type `func(h http.Handler) http.Handler` is the function that we will call when setting up our routes.

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
middleware := newMiddleware(logger, db, slackClient, rroll)
mux.Handle("/route1", middleware(handleSomething(handlerSpecificDeps))
mux.Handle("/route2", middleware(handleSomething2(handlerSpecificDeps))
mux.Handle("/route3", middleware(handleSomething3(handlerSpecificDeps))
mux.Handle("/route4", middleware(handleSomething4(handlerSpecificDeps))
```

Some people, but not I, like to formalize that function type like this:

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
// middleware is a function that wraps http.Handlers
// proving functionality before and after execution
// of the h handler.
type middleware func(h http.Handler) http.Handler
```

This is fine. Do it if you like. I’m not going to come around to your work, wait outside for you, and then walk alongside you with my arm around your shoulders in an intimidating way, asking if you’re pleased with yourself.

The reason I don’t do it is because it gives an extra level of indirection. When you look at the `newMiddleware` function’s signature above, it’s very clear what’s going on. If the return type is middleware, you have a little extra work to do. Essentially, I optimize for reading code, not writing it.

### An opportunity to hide the request/response types away

If an endpoint has its own request and response types, usually they’re only useful for that particular handler.

If that’s the case, you can define them inside the function.

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
func handleSomething() http.HandlerFunc {
	type request struct {
		Name string
	}
	type response struct {
		Greeting string `json:"greeting"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		...
	}
}
```

This keeps your global space clear and also prevents other handlers from relying on data you may not consider stable.

You sometimes encounter friction with this approach when your test code needs to use the same types. And to be fair, this is a good argument for breaking them out if that’s what you want to do.

## Use inline request/response types for additional storytelling in tests

If your request/response types are hidden inside the handler, you can just declare new types in your test code.

This is an opportunity to do a bit of storytelling to future generations who will need to understand your code.

For example, let’s say we have a `Person` type in our code, and we reuse it on many endpoints. If we had a `/greet` endpoint, we might only care about their name, so we can express this in test code:

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
func TestGreet(t *testing.T) {
	is := is.New(t)
	person := struct {
		Name string `json:"name"`
	}{
		Name: "Mat Ryer",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(person)
	is.NoErr(err) // json.NewEncoder
	req, err := http.NewRequest(http.MethodPost, "/greet", &buf)
	is.NoErr(err)
	//... more test code here
```

It’s clear from this test that the only field we care about is the `Name` field.

## `sync.Once` to defer setup

If I have to do anything expensive when preparing the handler, I defer it until that handler is first called.

This improves application startup time.

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
func (s *server) handleTemplate(files string...) http.HandlerFunc {
	var (
		init    sync.Once
		tpl     *template.Template
		tplerr  error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func(){
			tpl, tplerr = template.ParseFiles(files...)
		})
		if tplerr != nil {
		http.Error(w, tplerr.Error(), http.StatusInternalServerError)
			return
		}
		// use tpl
	}
}
```

`sync.Once` ensures the code is only executed one time, and other calls (other people making the same request) will block until it’s finished.

- The error check is outside of the `init` function, so if something does go wrong we still surface the error and won’t lose it in the logs
- If the handler is not called, the expensive work is never done — this can have big benefits depending on how your code is deployed

Remember that by doing this, you are moving the initialization time from startup to runtime (when the endpoint is first accessed). I use Google App Engine a lot, so this makes sense for me, but your case might be different, so it’s worth thinking about where and when to use `sync.Once` in this way.

## Designing for testability

These patterns evolved partly because of how easy they are to test the code. The `run` function is a simple way to run your program right from test code.

You have lots of options when it comes to testing in Go, and it’s less about right and wrong, and more about:

- How easy is it to understand what your program does by looking at the tests?
- How easy is it to change your code without worrying about breaking things?
- If all your tests pass, can you push to production, or does it need to cover more things?

### What is the unit when unit testing?

Following these patterns, the handlers themselves are also independently testable, but I tend not to do this, and I’ll explain why below. You have to consider what the best approach is for your project.

To test the handler only, you can:

1. Call the function to get the `http.Handler` — you have to pass in all the required dependencies (this is a feature).
2. Call the `ServeHTTP` method on the `http.Handler` you get back using a real `http.Request` and a `ResponseRecorder` from the `httptest` package (see https://pkg.go.dev/net/http/httptest#ResponseRecorder)
3. Make assertions about the response (check the status code, decode the body and make sure it’s right, check any important headers, etc.)

If you do this, you cut out any middleware like auth, and go straight to the handler code. This is nice if there is some specific complexity you want to build some test support around. However, there’s an advantage when your test code calls APIs in the same way your users will. I err on the side of end-to-end testing at this level, rather than unit testing all the pieces inside.

I would rather call the `run` function to execute the whole program as close to how it will run in production as possible. This will parse any arguments, connect to any dependencies, migrate the database, whatever else it will do in the wild, and eventually start up the server. Then when I hit the API from my test code, I am going through all the layers and even interacting with a real database. I am also testing `routes.go` at the same time.

I find I catch more issues earlier with this approach and I can avoid specifically testing boilerplate things. It also reduces the repetition in my tests. If I diligently test every layer, I can end up saying the same things multiple times in slightly different ways. You have to maintain all of this, so if you want to change something, updating one function and three tests doesn’t feel very productive. With end-to-end tests, you just have one set of main tests that describe the interactions between your users and your system.

I still use unit tests within this where appropriate. If I used TDD (which I often do) then I usually have a lot of tests done anyway, which I’m happy to maintain. But I will go back and delete tests if they’re repeating the same thing as an end-to-end test.

This decision will depend on lots of things, from the opinions of those around you to the complexity of your project, so like all the advice in this post, don’t fight to do this if it just doesn’t work for you.

### Testing with the run function

I like to call the `run` function from each test. Each test gets its own self-contained instance of the program. For each test, I can pass different arguments, flag values, standard-in and -out pipes, and even environment variables.

Since the `run` function takes a `context.Context`, and since all our code respects the context (right, everyone? It respects the context, right?) We can get a cancellation function by calling `context.WithCancel`. By deferring the `cancel` function, when the test function returns (i.e., when the tests have finished running) the context will be cancelled and the program will gracefully shut down. In Go 1.14 they added the `t.Cleanup` method which is a replacement for using the `defer` keyword yourself, and if you’d like to learn more about why that is, check out this issue: https://github.com/golang/go/issues/37333.

This is all achieved in surprisingly little code. Of course, you have to keep checking `ctx.Err` or `ctx.Done` all over the place too:

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

```go
func Test(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(cancel)
	go run(ctx)
	// test code goes here
```

### Waiting for readiness

Since the `run` function executes in a goroutine, we don’t really know exactly when it’s going to start up. If we’re going to start hitting the API like real users, we are going to need to know when it’s ready.

We could set up some way of signalling readiness, like a channel or something — but I prefer to have a `/healthz` or `/readyz` endpoint running on the server. As my old grandma used to say, the proof of the pudding is in the actual HTTP requests (she was way ahead of her time).

This is an example where our efforts to make the code more testable gives us an insight into what our users will need. They probably want to know if the service is ready or not as well, so why not have an official way to find this out?

To wait for a service to be ready, you can just write a loop:

Go![Copy code to clipboard](https://grafana.com/media/images/icons/icon-copy-small-2.svg)Copy

Expand code

```go
// waitForReady calls the specified endpoint until it gets a 200
// response or until the context is cancelled or the timeout is
// reached.
func waitForReady(
	ctx context.Context,
	timeout time.Duration,
	endpoint string,
) error {
	client := http.Client{}
	startTime := time.Now()
	for {
		req, err := http.NewRequestWithContext(
			ctx,
			http.MethodGet,
			endpoint,
			nil,
		)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error making request: %s\n", err.Error())
			continue
		}
		if resp.StatusCode == http.StatusOK {
			fmt.Println("Endpoint is ready!")
			resp.Body.Close()
			return nil
		}
		resp.Body.Close()

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if time.Since(startTime) >= timeout {
				return fmt.Errorf("timeout reached while waiting for endpoint")
			}
			// wait a little while between checks
			time.Sleep(250 * time.Millisecond)
		}
	}
}
```

## Putting this all into practice

Rolling simple APIs using these techniques remains my favorite way to go. It suits my aims of achieving maintainability excellence with code that’s easy to read, easy to extend by copying patterns, easy for new people to work with, easy to change without worrying, and explicitly done without any magic. This remains true even in cases where I use a code generation framework like our own [Oto package](https://github.com/pacedotdev/oto) to write the boilerplate for me based on templates that I customize.

On bigger projects or in larger organizations, especially one like Grafana Labs, you’ll often come across specific technology choices that impact these decisions. gRPC is a good example. In cases where there are established patterns and experience, or other tools or abstractions that are widely used, you will often find yourself making the pragmatic choice of going with the flow, as they say, although I suspect (or is it hope?) that there is still something useful in this post for you.

My day job is building out the new [Grafana IRM](https://grafana.com/products/cloud/irm/) suite with a talented group within Grafana Labs. The patterns discussed in this post help us deliver tools that people can rely on. “Tell me more about these great tools!,” I hear you scream at your monitor.

Most people use Grafana to visualize the operation of their systems, and with Grafana Alerting they are pinged when metrics fall outside of acceptable boundaries. With Grafana OnCall, your schedules and escalation rules automate the process of reaching out to the right people when things go wrong.

Grafana Incident lets you manage those unavoidable all-hands-on-deck moments that most of us are all too familiar with. It creates the Zoom room for you to discuss the issue, a dedicated Slack channel, and tracks the timeline of events while you focus on putting out the fire. In Slack, anything you mark with the robot face emoji as a reaction in the channel will be added to the timeline. This makes it very easy to collect key events as you go along, making debrief or post-incident review discussions much easier.

Try it out in Grafana Cloud today, or get in touch with your Grafana contact if you’re lucky enough to have one, and ask them about it.
