# BENZHI_README

## 项目说明

- 项目：zhanglei10281852-gif/portcoord-backend-qa-01
- 项目用途：PortCoord is a pure Go backend for coordinating arrival declarations, berthing windows, cargo work orders, pilot and tug tasks, capacity quotas, and responsibility handovers between port operators. It is designed as a production-style baseline for later Go coding-agent tasks. This baseline contains no intentionally seeded defect or private answer material.
- Go 工具链：`golang:1.26`
- 前端工具链：无

## 标准构建、运行和测试命令

进入容器后执行：

```bash
# 编译
cd '/app' && GOTOOLCHAIN=local go build ./...

# 启动
cd '/app' && GOTOOLCHAIN=local go run ./cmd/executor
cd '/app' && GOTOOLCHAIN=local go run ./cmd/migrate
cd '/app' && GOTOOLCHAIN=local go run ./cmd/scheduler

# 测试
cd '/app' && GOTOOLCHAIN=local go test ./...
```

## Docker 构建和进入容器

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh benzhi-task-31-amd64 linux/amd64
./build_benzhi_docker.sh benzhi-task-31-arm64 linux/arm64
docker run -it benzhi-task-31-amd64:latest
docker run -it --platform linux/arm64 benzhi-task-31-arm64:latest
```

## 题目验证命令

1. 预期退出码 0：`go test ./internal/quota -run TestQuota_StorageFailureIsReturnedToCaller -count=1`
2. 预期退出码 0：`go test -buildvcs=false -count=1 ./...`
3. 预期退出码 0：`GOTOOLCHAIN=local go build -buildvcs=false ./... && GOTOOLCHAIN=local go vet ./...`

## Bug 复现

Bug 现象、触发步骤和完整错误信息见 `BUG_REPRO.md`。
