![gopher-schema-as-code](https://s3.eu-central-1.amazonaws.com/entgo.io/assets/gopher-schema-as-code.png)

# ent

ent 是一个简单而又强大的 Go 实体框架，它可以轻松地构建和维护大型数据模型的应用程序，并坚持以下原则：

- 轻松地将数据库模式建模为图结构。
- 将模式定义为一个程序化的 Go 代码。
- 基于代码生成的静态类型。
- 数据库查询和图遍历易于编写。
- 使用 Go 模板进行简单的扩展和定制。

# Hello World

```sh
$ go get entgo.io/ent/cmd/ent
$ ent init User
```

上面的命令将在 `<project>/ent/schema/` 目录下生成 User 的签名：

```go
// <project>/ent/schema/user.go

package schema

import "entgo.io/ent"

// User holds the schema definition for the User entity.
type User struct {
    ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
    return nil
}

// Edges of the User.
func (User) Edges() []ent.Edge {
    return nil
}
```

然后新增自定义的属性：

```go
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)


// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.Int("age").
            Positive(),
        field.String("name").
            Default("unknown"),
    }
}
```

从项目的根目录下运行 `go generate`，如下：

```go
ent
├── client.go
├── config.go
├── context.go
├── ent.go
├── migrate
│   ├── migrate.go
│   └── schema.go
├── predicate
│   └── predicate.go
├── schema
│   └── user.go
├── tx.go
├── user
│   ├── user.go
│   └── where.go
├── user.go
├── user_create.go
├── user_delete.go
├── user_query.go
└── user_update.go
```

在代码使用中：

```go
package main

import (
    "context"
    "log"

    "<project>/ent"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
    if err != nil {
        log.Fatalf("failed opening connection to sqlite: %v", err)
    }
    defer client.Close()
    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}
```

现在，我们准备好创建我们的用户了。为了举例，我们把这个函数叫做 CreateUser。

```go
func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
    u, err := client.User.
        Create().
        SetAge(30).
        SetName("a8m").
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed creating user: %v", err)
    }
    log.Println("user was created: ", u)
    return u, nil
}
```

ent 为每个实体模式生成一个包，其中包含了它的谓词、默认值、验证器和关于存储元素（列名、主键等）的附加信息。

```go
package main

import (
    "log"

    "<project>/ent"
    "<project>/ent/user"
)

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
    u, err := client.User.
        Query().
        Where(user.NameEQ("a8m")).
        // `Only` fails if no user found,
        // or more than 1 user returned.
        Only(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed querying user: %v", err)
    }
    log.Println("user returned: ", u)
    return u, nil
}
```

## 关系

```go
go run entgo.io/ent/cmd/ent init Car Group
```

然后手动新增如下代码：

```go
import (
    "regexp"

    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)

// Fields of the Car.
func (Car) Fields() []ent.Field {
    return []ent.Field{
        field.String("model"),
        field.Time("registered_at"),
    }
}


// Fields of the Group.
func (Group) Fields() []ent.Field {
    return []ent.Field{
        field.String("name").
            // Regexp validation for group name.
            Match(regexp.MustCompile("[a-zA-Z_]+$")),
    }
}
```

让我们定义第一个关系。从 User 到 Car 的一条边，定义了一个用户可以拥有 1 辆或多辆汽车，但一辆汽车只有一个车主（一对多关系）。

![关系示意图](https://s3.ax1x.com/2021/02/04/ylLVqP.png)

```go
import (
   "log"

   "entgo.io/ent"
   "entgo.io/ent/schema/edge"
)

// Edges of the User.
func (User) Edges() []ent.Edge {
   return []ent.Edge{
       edge.To("cars", Car.Type),
   }
}
```

我们继续我们的例子，创建 2 辆汽车，并将它们添加到一个用户中。

```go
func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
    // Create a new car with model "Tesla".
    tesla, err := client.Car.
        Create().
        SetModel("Tesla").
        SetRegisteredAt(time.Now()).
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed creating car: %v", err)
    }

    // Create a new car with model "Ford".
    ford, err := client.Car.
        Create().
        SetModel("Ford").
        SetRegisteredAt(time.Now()).
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed creating car: %v", err)
    }
    log.Println("car was created: ", ford)

    // Create a new user, and add it the 2 cars.
    a8m, err := client.User.
        Create().
        SetAge(30).
        SetName("a8m").
        AddCars(tesla, ford).
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed creating user: %v", err)
    }
    log.Println("user was created: ", a8m)
    return a8m, nil
}
```

然后在查询的时候：

```go
import (
    "log"

    "<project>/ent"
    "<project>/ent/car"
)

func QueryCars(ctx context.Context, a8m *ent.User) error {
    cars, err := a8m.QueryCars().All(ctx)
    if err != nil {
        return fmt.Errorf("failed querying user cars: %v", err)
    }
    log.Println("returned cars:", cars)

    // What about filtering specific cars.
    ford, err := a8m.QueryCars().
        Where(car.ModelEQ("Ford")).
        Only(ctx)
    if err != nil {
        return fmt.Errorf("failed querying user cars: %v", err)
    }
    log.Println(ford)
    return nil
}
```
