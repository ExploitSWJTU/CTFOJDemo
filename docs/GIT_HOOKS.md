# Git Hooks 配置说明

## 目录结构

```
/
├── scripts/
│   └── pre-commit.mjs      # 智能 pre-commit 脚本
├── package.json             # 根目录配置（管理 git hooks）
├── lint-staged.config.js    # lint-staged 配置
├── frontend/
│   └── package.json         # 前端配置
└── backend/
    ├── .golangci.yml        # Go linter 配置
    └── go.mod               # Go 模块配置
```

## 工作原理

### Pre-commit Hook 流程

```
git commit
    ↓
scripts/pre-commit.mjs
    ↓
分析暂存文件
    ├── 前端文件？ → 运行 lint-staged (eslint + type-check)
    └── 后端文件？ → 运行 go fmt + go vet
    ↓
全部通过？ → 提交成功
    ↓
有失败？ → 中断提交
```

### 智能检测

| 变更文件 | 运行检查 |
|---------|---------|
| `frontend/**/*.vue` | ESLint + Vue TSC |
| `frontend/**/*.ts` | ESLint + Vue TSC |
| `backend/**/*.go` | go fmt + go vet |
| 仅文档/配置 | 跳过检查 |

## 可用命令

```bash
# 开发
pnpm dev              # 启动前端开发服务器

# 构建
pnpm build            # 构建前端生产版本

# 代码检查
pnpm lint             # 前端 lint
pnpm lint:fix         # 前端 lint + 自动修复
pnpm lint:backend     # 后端 go fmt

pnpm type-check       # 前端类型检查
pnpm type-check:backend  # 后端 go vet

pnpm check            # 前端 lint + type-check
pnpm check:all        # 前后端完整检查
```

## 跳过 Git Hook

紧急情况可跳过 hook：

```bash
git commit --no-verify -m "紧急修复"
```

**注意：** 仅在特殊情况下使用，正常情况下请修复代码问题。

## 安装/更新 Hooks

```bash
# 首次安装
pnpm install

# 手动更新 hooks
npx simple-git-hooks
```

## 故障排查

### Hook 不执行

1. 检查 `.git/hooks/pre-commit` 是否存在
2. 运行 `npx simple-git-hooks` 重新安装
3. 检查 Node.js 版本（推荐 18+）

### 类型检查失败

前端：
```bash
cd frontend && pnpm type-check
```

后端：
```bash
cd backend && go vet ./...
```

### ESLint 报错

```bash
cd frontend && pnpm lint:fix
```

## 最佳实践

1. **小步提交** - 每次提交只改动少量文件，加快检查速度
2. **本地先跑检查** - `pnpm check` 通过后再 commit
3. **及时修复** - hook 报错立即修复，不要累积
4. **团队同步** - 确保团队成员使用相同版本的工具
