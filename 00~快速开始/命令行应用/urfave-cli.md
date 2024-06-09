# urfave/cli

```go
package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli/v2"
)

func main() {
  app := &cli.App{
    Name: "greet",
    Usage: "fight the loneliness!",
    Action: func(c *cli.Context) error {
      fmt.Println("Hello friend!")
      return nil
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
```

## Flags

```go
package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli/v2"
)

func main() {
  app := &cli.App{
    Flags: []cli.Flag {
      &cli.StringFlag{
        Name: "lang",
        Value: "english",
        Usage: "language for the greeting",
      },
    },
    Action: func(c *cli.Context) error {
      name := "Nefertiti"
      if c.NArg() > 0 {
        name = c.Args().Get(0)
      }
      if c.String("lang") == "spanish" {
        fmt.Println("Hola", name)
      } else {
        fmt.Println("Hello", name)
      }
      return nil
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
```

您还可以为标志设置目标变量，将内容扫描到该目标变量。

```go
package main

import (
  "log"
  "os"
  "fmt"

  "github.com/urfave/cli/v2"
)

func main() {
  var language string

  app := &cli.App{
    Flags: []cli.Flag {
      &cli.StringFlag{
        Name:        "lang",
        Value:       "english",
        Usage:       "language for the greeting",
        Destination: &language,
      },
    },
    Action: func(c *cli.Context) error {
      name := "someone"
      if c.NArg() > 0 {
        name = c.Args().Get(0)
      }
      if language == "spanish" {
        fmt.Println("Hola", name)
      } else {
        fmt.Println("Hello", name)
      }
      return nil
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
```

## Subcommands

```go
package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli/v2"
)

func main() {
  app := &cli.App{
    Commands: []*cli.Command{
      {
        Name:    "add",
        Aliases: []string{"a"},
        Usage:   "add a task to the list",
        Action:  func(c *cli.Context) error {
          fmt.Println("added task: ", c.Args().First())
          return nil
        },
      },
      {
        Name:    "complete",
        Aliases: []string{"c"},
        Usage:   "complete a task on the list",
        Action:  func(c *cli.Context) error {
          fmt.Println("completed task: ", c.Args().First())
          return nil
        },
      },
      {
        Name:        "template",
        Aliases:     []string{"t"},
        Usage:       "options for task templates",
        Subcommands: []*cli.Command{
          {
            Name:  "add",
            Usage: "add a new template",
            Action: func(c *cli.Context) error {
              fmt.Println("new task template: ", c.Args().First())
              return nil
            },
          },
          {
            Name:  "remove",
            Usage: "remove an existing template",
            Action: func(c *cli.Context) error {
              fmt.Println("removed task template: ", c.Args().First())
              return nil
            },
          },
        },
      },
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
```

对于具有许多子命令的应用程序中的其他组织，您可以为每个命令关联一个类别，以将它们分组到帮助输出中。

```go
package main

import (
  "log"
  "os"

  "github.com/urfave/cli/v2"
)

func main() {
  app := &cli.App{
    Commands: []*cli.Command{
      {
        Name: "noop",
      },
      {
        Name:     "add",
        Category: "template",
      },
      {
        Name:     "remove",
        Category: "template",
      },
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}

/*
    COMMANDS:
    noop

    Template actions:
        add
        remove
*/
```

## 完整示例

```go
package main

import (
  "errors"
  "flag"
  "fmt"
  "io"
  "io/ioutil"
  "os"
  "time"

  "github.com/urfave/cli/v2"
)

func init() {
  cli.AppHelpTemplate += "\nCUSTOMIZED: you bet ur muffins\n"
  cli.CommandHelpTemplate += "\nYMMV\n"
  cli.SubcommandHelpTemplate += "\nor something\n"

  cli.HelpFlag = &cli.BoolFlag{Name: "halp"}
  cli.VersionFlag = &cli.BoolFlag{Name: "print-version", Aliases: []string{"V"}}

  cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
    fmt.Fprintf(w, "best of luck to you\n")
  }
  cli.VersionPrinter = func(c *cli.Context) {
    fmt.Fprintf(c.App.Writer, "version=%s\n", c.App.Version)
  }
  cli.OsExiter = func(c int) {
    fmt.Fprintf(cli.ErrWriter, "refusing to exit %d\n", c)
  }
  cli.ErrWriter = ioutil.Discard
  cli.FlagStringer = func(fl cli.Flag) string {
    return fmt.Sprintf("\t\t%s", fl.Names()[0])
  }
}

type hexWriter struct{}

func (w *hexWriter) Write(p []byte) (int, error) {
  for _, b := range p {
    fmt.Printf("%x", b)
  }
  fmt.Printf("\n")

  return len(p), nil
}

type genericType struct {
  s string
}

func (g *genericType) Set(value string) error {
  g.s = value
  return nil
}

func (g *genericType) String() string {
  return g.s
}

func main() {
  app := &cli.App{
    Name: "kənˈtrīv",
    Version: "v19.99.0",
    Compiled: time.Now(),
    Authors: []*cli.Author{
      &cli.Author{
        Name:  "Example Human",
        Email: "human@example.com",
      },
    },
    Copyright: "(c) 1999 Serious Enterprise",
    HelpName: "contrive",
    Usage: "demonstrate available API",
    UsageText: "contrive - demonstrating the available API",
    ArgsUsage: "[args and such]",
    Commands: []*cli.Command{
      &cli.Command{
        Name:        "doo",
        Aliases:     []string{"do"},
        Category:    "motion",
        Usage:       "do the doo",
        UsageText:   "doo - does the dooing",
        Description: "no really, there is a lot of dooing to be done",
        ArgsUsage:   "[arrgh]",
        Flags: []cli.Flag{
          &cli.BoolFlag{Name: "forever", Aliases: []string{"forevvarr"}},
        },
        Subcommands: []*cli.Command{
          &cli.Command{
            Name:   "wop",
            Action: wopAction,
          },
        },
        SkipFlagParsing: false,
        HideHelp:        false,
        Hidden:          false,
        HelpName:        "doo!",
        BashComplete: func(c *cli.Context) {
          fmt.Fprintf(c.App.Writer, "--better\n")
        },
        Before: func(c *cli.Context) error {
          fmt.Fprintf(c.App.Writer, "brace for impact\n")
          return nil
        },
        After: func(c *cli.Context) error {
          fmt.Fprintf(c.App.Writer, "did we lose anyone?\n")
          return nil
        },
        Action: func(c *cli.Context) error {
          c.Command.FullName()
          c.Command.HasName("wop")
          c.Command.Names()
          c.Command.VisibleFlags()
          fmt.Fprintf(c.App.Writer, "dodododododoodododddooooododododooo\n")
          if c.Bool("forever") {
            c.Command.Run(c)
          }
          return nil
        },
        OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
          fmt.Fprintf(c.App.Writer, "for shame\n")
          return err
        },
      },
    },
    Flags: []cli.Flag{
      &cli.BoolFlag{Name: "fancy"},
      &cli.BoolFlag{Value: true, Name: "fancier"},
      &cli.DurationFlag{Name: "howlong", Aliases: []string{"H"}, Value: time.Second * 3},
      &cli.Float64Flag{Name: "howmuch"},
      &cli.GenericFlag{Name: "wat", Value: &genericType{}},
      &cli.Int64Flag{Name: "longdistance"},
      &cli.Int64SliceFlag{Name: "intervals"},
      &cli.IntFlag{Name: "distance"},
      &cli.IntSliceFlag{Name: "times"},
      &cli.StringFlag{Name: "dance-move", Aliases: []string{"d"}},
      &cli.StringSliceFlag{Name: "names", Aliases: []string{"N"}},
      &cli.UintFlag{Name: "age"},
      &cli.Uint64Flag{Name: "bigage"},
    },
    EnableBashCompletion: true,
    HideHelp: false,
    HideVersion: false,
    BashComplete: func(c *cli.Context) {
      fmt.Fprintf(c.App.Writer, "lipstick\nkiss\nme\nlipstick\nringo\n")
    },
    Before: func(c *cli.Context) error {
      fmt.Fprintf(c.App.Writer, "HEEEERE GOES\n")
      return nil
    },
    After: func(c *cli.Context) error {
      fmt.Fprintf(c.App.Writer, "Phew!\n")
      return nil
    },
    CommandNotFound: func(c *cli.Context, command string) {
      fmt.Fprintf(c.App.Writer, "Thar be no %q here.\n", command)
    },
    OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
      if isSubcommand {
        return err
      }

      fmt.Fprintf(c.App.Writer, "WRONG: %#v\n", err)
      return nil
    },
    Action: func(c *cli.Context) error {
      cli.DefaultAppComplete(c)
      cli.HandleExitCoder(errors.New("not an exit coder, though"))
      cli.ShowAppHelp(c)
      cli.ShowCommandCompletions(c, "nope")
      cli.ShowCommandHelp(c, "also-nope")
      cli.ShowCompletions(c)
      cli.ShowSubcommandHelp(c)
      cli.ShowVersion(c)

      fmt.Printf("%#v\n", c.App.Command("doo"))
      if c.Bool("infinite") {
        c.App.Run([]string{"app", "doo", "wop"})
      }

      if c.Bool("forevar") {
        c.App.RunAsSubcommand(c)
      }
      c.App.Setup()
      fmt.Printf("%#v\n", c.App.VisibleCategories())
      fmt.Printf("%#v\n", c.App.VisibleCommands())
      fmt.Printf("%#v\n", c.App.VisibleFlags())

      fmt.Printf("%#v\n", c.Args().First())
      if c.Args().Len() > 0 {
        fmt.Printf("%#v\n", c.Args().Get(1))
      }
      fmt.Printf("%#v\n", c.Args().Present())
      fmt.Printf("%#v\n", c.Args().Tail())

      set := flag.NewFlagSet("contrive", 0)
      nc := cli.NewContext(c.App, set, c)

      fmt.Printf("%#v\n", nc.Args())
      fmt.Printf("%#v\n", nc.Bool("nope"))
      fmt.Printf("%#v\n", !nc.Bool("nerp"))
      fmt.Printf("%#v\n", nc.Duration("howlong"))
      fmt.Printf("%#v\n", nc.Float64("hay"))
      fmt.Printf("%#v\n", nc.Generic("bloop"))
      fmt.Printf("%#v\n", nc.Int64("bonk"))
      fmt.Printf("%#v\n", nc.Int64Slice("burnks"))
      fmt.Printf("%#v\n", nc.Int("bips"))
      fmt.Printf("%#v\n", nc.IntSlice("blups"))
      fmt.Printf("%#v\n", nc.String("snurt"))
      fmt.Printf("%#v\n", nc.StringSlice("snurkles"))
      fmt.Printf("%#v\n", nc.Uint("flub"))
      fmt.Printf("%#v\n", nc.Uint64("florb"))

      fmt.Printf("%#v\n", nc.FlagNames())
      fmt.Printf("%#v\n", nc.IsSet("wat"))
      fmt.Printf("%#v\n", nc.Set("wat", "nope"))
      fmt.Printf("%#v\n", nc.NArg())
      fmt.Printf("%#v\n", nc.NumFlags())
      fmt.Printf("%#v\n", nc.Lineage()[1])

      nc.Set("wat", "also-nope")

      ec := cli.Exit("ohwell", 86)
      fmt.Fprintf(c.App.Writer, "%d", ec.ExitCode())
      fmt.Printf("made it!\n")
      return ec
    },
    Metadata: map[string]interface{}{
      "layers":          "many",
      "explicable":      false,
      "whatever-values": 19.99,
    },
  }

  if os.Getenv("HEXY") != "" {
    app.Writer = &hexWriter{}
    app.ErrWriter = &hexWriter{}
  }

  app.Run(os.Args)
}

func wopAction(c *cli.Context) error {
  fmt.Fprintf(c.App.Writer, ":wave: over here, eh\n")
  return nil
}
```
