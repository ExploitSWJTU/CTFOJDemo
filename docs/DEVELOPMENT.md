# 开发环境配置指南

## 前置要求

- Node.js >= 18
- Go >= 1.20
- Docker & Docker Compose
- VSCode (推荐)

## 快速开始

### 1. 安装依赖

```bash
# 前端依赖
cd frontend && pnpm install

# 后端依赖
cd backend && go mod download

# 根目录工具
cd .. && pnpm install
```

### 2. 安装开发工具

```bash
# Go 热重载工具
go install github.com/air-verse/air@latest

# Go 调试工具
go install github.com/go-delve/delve/cmd/dlv@latest

# Go linter (可选)
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### 3. 启动基础设施

```bash
docker compose up -d db redis
```

### 4. 开发模式

**方式一：全栈启动（推荐）**
```bash
pnpm dev
```

**方式二：分别启动**
```bash
# 终端 1 - 后端 (热重载)
cd backend && air

# 终端 2 - 前端
cd frontend && pnpm dev
```

**方式三：VSCode 调试**
1. 按 `F5` 或点击调试面板
2. 选择 "Launch Backend" 或 "Launch Frontend"
3. 后端会自动进入调试模式，可设置断点

## VSCode 配置

### 扩展推荐

- Go (官方)
- Vue - Official
- ESLint
- Prettier
- Docker

### 调试配置

项目已包含 `.vscode/launch.json`，支持：

- **Launch Backend**: Go 后端调试（Delve）
- **Launch Frontend**: Chrome 前端调试

### 任务配置

按 `Ctrl+Shift+B` (或 `Cmd+Shift+B` on Mac) 运行任务：

- `air: backend` - 后端热重载
- `npm: dev` - 前端开发服务器
- `docker: up` - 启动数据库

## 环境变量

创建 `.env` 文件（参考 `.env.example`）：

```bash
# Database
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=ctfoj
DB_PORT=5432

# JWT
JWT_SECRET=dev-secret-key-change-in-production

# Backend
PORT=8080
GIN_MODE=debug
```

## 项目结构

```
/
├── frontend/           # Vue 3 前端
│   ├── src/
│   │   ├── api/       # API 封装
│   │   ├── stores/    # Pinia 状态
│   │   ├── views/     # 页面
│   │   └── types/     # TypeScript 类型
│   └── package.json
├── backend/            # Go 后端
│   ├── cmd/
│   │   └── server/    # 主入口
│   ├── internal/
│   │   ├── api/       # HTTP 处理器
│   │   ├── model/     # GORM 模型
│   │   └── service/   # 业务逻辑
│   └── .air.toml      # Air 配置
├── docker-compose.yml  # Docker 编排
└── .vscode/           # VSCode 配置
```

## 常见问题

### Air 不工作

确保已安装 air：
```bash
go install github.com/air-verse/air@latest
```

检查 `~/.air.toml` 或项目中的 `.air.toml`

### 数据库连接失败

1. 检查 Docker 是否运行：`docker ps`
2. 重启数据库：`docker compose restart db`
3. 查看日志：`docker compose logs db`

### 端口冲突

修改 `.env` 中的端口：
```bash
PORT=8081  # 后端
DB_PORT=5433  # PostgreSQL
```

## 代码质量

```bash
# 前端
pnpm lint:fix
pnpm type-check

# 后端
cd backend && go fmt ./...
cd backend && go vet ./...
cd backend && golangci-lint run

# 全部检查
pnpm check:all
```

## Git Hooks

Pre-commit hook 会自动运行：
- 前端：ESLint + TypeScript 检查
- 后端：go fmt + go vet

跳过 hook（紧急情况下）：
```bash
git commit --no-verify -m "紧急修复"
```
