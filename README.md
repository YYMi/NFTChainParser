# My-Gin-App

My-Gin-App 是一个使用 Gin 框架构建的 Go 语言 Web 应用程序。该项目遵循模块化的目录结构，易于管理、扩展和维护。应用程序支持多版本 API，并包含 JWT 认证、CORS 中间件和数据库连接等功能。

## 目录
- [项目结构](#项目结构)
- [快速开始](#快速开始)
- [配置](#配置)
- [运行应用程序](#运行应用程序)
- [API 文档](#api-文档)
- [脚本](#脚本)
- [贡献](#贡献)
- [许可证](#许可证)

## 项目结构

项目的目录结构如下：

```
my-gin-app/
nftExchangeAdmi-gin/
│
├── config/                    # 配置文件及加载逻辑
├── controller/                # 控制器层，处理请求与响应
├── db/                        # 数据库连接与初始化
├── docs/                      # 项目文档及 Swagger 文档文件
├── errors/                    # 统一错误定义与处理
├── etc/                       # 额外配置或资源文件夹
├── log/                       # 日志文件存放目录
├── middleware/                # 中间件（JWT、CORS、日志等）
├── repository/                # 数据库操作层，封装 CRUD
│   └── bcRepository/          # 特定业务模块的数据访问层示例
├── router/                    # 路由配置
├── service/                   # 服务层，封装业务逻辑
├── task/                      # 后台任务（如定时任务、XXL-JOB集成）
│   └── xxljob/
├── tools/                     # 工具初始化与通用函数
│   └── rest/                  # REST 启动与初始化（如 Gin Engine）
├── types/                     # 通用数据类型定义 (PO, VO, DTO)
├── util/                      # 工具包（如加解密、哈希、JWT、统一响应封装）
├── main.go                    # 程序入口文件
├── go.mod                     # Go Modules依赖
├── go.sum                     # 依赖校验文件
└── README.md                  # 项目说明文档
```

## 快速开始

### 前置条件

- Go 1.18+
- MySQL 或其他支持的数据库
- Redis（可选，用于缓存和会话管理）

### 安装步骤

1. 克隆仓库：
   ```bash
   git clone https://github.com/yourusername/my-gin-app.git
   cd my-rest-app
   ```
2. 安装依赖：
   ```bash
   go mod tidy
   ```
3. 创建 `.env` 文件并复制示例：
   ```bash
   cp .env.example .env
   ```
   更新 `.env` 文件中的配置变量以匹配您的环境。

### 数据库设置

1. 更新 `config/config.yaml` 文件以配置数据库连接。
2. 运行数据库迁移脚本：
   ```bash
   ./script/migrate.sh
   ```

## 运行应用程序

通过运行以下命令启动应用程序：

```bash
go run main.go
```

服务器将启动在配置文件中指定的端口（默认：`8080`）。

## 配置

应用程序使用 `config.yaml` 文件进行配置设置，包括服务器端口、数据库设置和 JWT 密钥。请确保根据您的环境更新这些设置。

## API 文档

该项目包含使用 Swagger 自动生成的 API 文档。您可以通过访问以下地址查看 Swagger 文档：
Swagger-相关注释等已经开发了一个插件swag-note  按照标准模板写代码 可以自己生成相关注释
gormt  插件在原本的基础上也添加了新的功能 ， 可以指定读取配置文件在生成相关代码 

```
http://localhost:8080/swagger/index.html
```

## 脚本

- **migrate.sh**: 此脚本用于数据库迁移和初始化表的设置。

运行脚本：
```bash
./script/migrate.sh
```

## 贡献

欢迎贡献代码！请 fork 本仓库并提交 pull request。

1. Fork 仓库！
2. 创建您的功能分支：`git checkout -b my-new-feature`
3. 提交您的修改：`git commit -am 'Add some feature'`
4. 推送到分支：`git push origin my-new-feature`
5. 提交 pull request :D

## 许可证

版本后续 跟进中 有相关问题 请联系邮箱yuaizifeng@gmail.com

