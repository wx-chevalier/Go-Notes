# Go-Blueprint 使用教程

## 1. 安装

首先需要安装 go-blueprint CLI 工具：

```bash
go install github.com/melkeydev/go-blueprint@latest
```

确保你的 `$GOPATH/bin` 已添加到系统的 PATH 环境变量中。

## 2. 基本使用

### 2.1 创建新项目（交互式）

最简单的方式是运行：

```bash
go-blueprint create
```

这将启动交互式命令行界面，引导你完成项目创建过程。

### 2.2 使用命令行参数创建项目

```bash
# 基本用法
go-blueprint create --name my-project --framework gin --driver postgres --git commit

# 参数说明：
# --name: 项目名称
# --framework: 选择框架
# --driver: 数据库驱动
# --git: 初始化 git 仓库
```

## 3. 支持的框架

Go-Blueprint 支持以下主流 Web 框架：

- Chi
- Gin
- Fiber
- HttpRouter
- Gorilla/mux
- Echo

## 4. 数据库支持

支持的数据库：

- MySQL
- PostgreSQL
- SQLite
- MongoDB
- Redis
- ScyllaDB (GoCQL)

## 5. 高级特性使用

### 5.1 启用高级特性

使用 `--advanced` 标志来访问高级特性：

```bash
go-blueprint create --advanced
```

### 5.2 具体高级特性示例

1. **添加 HTMX 支持**：

```bash
go-blueprint create --advanced --feature htmx
```

2. **添加 GitHub Actions**：

```bash
go-blueprint create --advanced --feature githubaction
```

3. **添加 WebSocket 支持**：

```bash
go-blueprint create --advanced --feature websocket
```

4. **添加 Tailwind CSS**：

```bash
go-blueprint create --advanced --feature tailwind
```

5. **添加 Docker 配置**：

```bash
go-blueprint create --advanced --feature docker
```

6. **添加 React 前端**：

```bash
go-blueprint create --advanced --feature react
```

### 5.3 组合多个特性

可以同时使用多个特性：

```bash
go-blueprint create --name my-project \
    --framework chi \
    --driver mysql \
    --advanced \
    --feature htmx \
    --feature githubaction \
    --feature websocket \
    --feature tailwind \
    --feature docker \
    --feature react \
    --git commit
```

## 6. 项目结构示例

使用 go-blueprint 创建的项目通常会包含以下结构：

```
my-project/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   └── database/
├── pkg/
├── configs/
├── scripts/
├── .github/        # 如果选择了 GitHub Actions
├── Dockerfile      # 如果选择了 Docker
├── go.mod
└── README.md
```

## 7. 最佳实践建议

1. **选择合适的框架**

   - 如果需要高性能，考虑使用 Fiber
   - 如果需要简单易用，考虑使用 Gin
   - 如果需要轻量级，考虑使用 Chi

2. **数据库选择**

   - 关系型数据需求选择 PostgreSQL 或 MySQL
   - 需要高性能缓存选择 Redis
   - 需要文档数据库选择 MongoDB

3. **开发流程**
   - 创建项目后先阅读生成的 README.md
   - 检查配置文件确保符合需求
   - 根据需要调整项目结构

## 8. 常见问题解决

1. 如果遇到权限问题：

```bash
sudo chmod +x $(go env GOPATH)/bin/go-blueprint
```

2. 如果需要更新到最新版本：

```bash
go install github.com/melkeydev/go-blueprint@latest
```

## 参考链接

- [Go-Blueprint GitHub 仓库](https://github.com/Melkeydev/go-blueprint)
- [官方文档](https://docs.go-blueprint.dev/)

这个教程涵盖了 go-blueprint 的主要功能和使用方法。你可以根据具体需求选择合适的特性来创建你的 Go 项目。需要了解更详细的某个部分吗？好的，我将为您创建一个基于 go-blueprint 项目的详细教程。这个项目是一个非常受欢迎的 Go 项目脚手架工具（目前在 GitHub 上有 5.3k+ stars）。

# Go-Blueprint 使用教程

## 1. 安装

首先需要安装 go-blueprint CLI 工具：

```bash
go install github.com/melkeydev/go-blueprint@latest
```

确保你的 `$GOPATH/bin` 已添加到系统的 PATH 环境变量中。

## 2. 基本使用

### 2.1 创建新项目（交互式）

最简单的方式是运行：

```bash
go-blueprint create
```

这将启动交互式命令行界面，引导你完成项目创建过程。

### 2.2 使用命令行参数创建项目

```bash
# 基本用法
go-blueprint create --name my-project --framework gin --driver postgres --git commit

# 参数说明：
# --name: 项目名称
# --framework: 选择框架
# --driver: 数据库驱动
# --git: 初始化 git 仓库
```

## 3. 支持的框架

Go-Blueprint 支持以下主流 Web 框架：

- Chi
- Gin
- Fiber
- HttpRouter
- Gorilla/mux
- Echo

## 4. 数据库支持

支持的数据库：

- MySQL
- PostgreSQL
- SQLite
- MongoDB
- Redis
- ScyllaDB (GoCQL)

## 5. 高级特性使用

### 5.1 启用高级特性

使用 `--advanced` 标志来访问高级特性：

```bash
go-blueprint create --advanced
```

### 5.2 具体高级特性示例

1. **添加 HTMX 支持**：

```bash
go-blueprint create --advanced --feature htmx
```

2. **添加 GitHub Actions**：

```bash
go-blueprint create --advanced --feature githubaction
```

3. **添加 WebSocket 支持**：

```bash
go-blueprint create --advanced --feature websocket
```

4. **添加 Tailwind CSS**：

```bash
go-blueprint create --advanced --feature tailwind
```

5. **添加 Docker 配置**：

```bash
go-blueprint create --advanced --feature docker
```

6. **添加 React 前端**：

```bash
go-blueprint create --advanced --feature react
```

### 5.3 组合多个特性

可以同时使用多个特性：

```bash
go-blueprint create --name my-project \
    --framework chi \
    --driver mysql \
    --advanced \
    --feature htmx \
    --feature githubaction \
    --feature websocket \
    --feature tailwind \
    --feature docker \
    --feature react \
    --git commit
```

## 6. 项目结构示例

使用 go-blueprint 创建的项目通常会包含以下结构：

```
my-project/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   └── database/
├── pkg/
├── configs/
├── scripts/
├── .github/        # 如果选择了 GitHub Actions
├── Dockerfile      # 如果选择了 Docker
├── go.mod
└── README.md
```

## 7. 最佳实践建议

1. **选择合适的框架**

   - 如果需要高性能，考虑使用 Fiber
   - 如果需要简单易用，考虑使用 Gin
   - 如果需要轻量级，考虑使用 Chi

2. **数据库选择**

   - 关系型数据需求选择 PostgreSQL 或 MySQL
   - 需要高性能缓存选择 Redis
   - 需要文档数据库选择 MongoDB

3. **开发流程**
   - 创建项目后先阅读生成的 README.md
   - 检查配置文件确保符合需求
   - 根据需要调整项目结构

## 8. 常见问题解决

1. 如果遇到权限问题：

```bash
sudo chmod +x $(go env GOPATH)/bin/go-blueprint
```

2. 如果需要更新到最新版本：

```bash
go install github.com/melkeydev/go-blueprint@latest
```

## 参考链接

- [Go-Blueprint GitHub 仓库](https://github.com/Melkeydev/go-blueprint)
- [官方文档](https://docs.go-blueprint.dev/)

这个教程涵盖了 go-blueprint 的主要功能和使用方法。你可以根据具体需求选择合适的特性来创建你的 Go 项目。需要了解更详细的某个部分吗？
