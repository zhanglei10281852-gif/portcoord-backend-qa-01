# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

当额度存储不可用时，额度预留接口返回了“额度耗尽”的正常业务结果，调用方无法区分容量不足和服务故障。请修复错误处理，确保存储故障返回明确的服务错误，同时保留真正额度不足时的正常拒绝行为。

## 含 Bug 版本

- 仓库：zhanglei10281852-gif/portcoord-backend-qa-01
- 仓库地址：https://github.com/zhanglei10281852-gif/portcoord-backend-qa-01.git
- parent SHA：3d172896950c5e6ae67c8d2b2e093ddb0720d981

## 复现步骤

```bash
git clone -- https://github.com/zhanglei10281852-gif/portcoord-backend-qa-01.git bug-repro
cd bug-repro
git checkout --detach 3d172896950c5e6ae67c8d2b2e093ddb0720d981
go test ./internal/quota -run TestQuota_StorageFailureIsReturnedToCaller -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/quota -run TestQuota_StorageFailureIsReturnedToCaller -count=1
--- FAIL: TestQuota_StorageFailureIsReturnedToCaller (0.01s)
    quota_test.go:310: expected storage error, got result &{QuotaID: QuotaType:cabin Available:0 Reserved:0 Status:exhausted WaitEstimate:0s Rejected:true Message:quota cabin exhausted: available 0, requested 10}
FAIL
FAIL	portcoord/internal/quota	0.011s
FAIL

```

stderr：

```text
warning: internal/quota/quota_test.go has type 100755, expected 100644
warning: internal/quota/quota_test.go has type 100755, expected 100644

```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/quota -run TestQuota_StorageFailureIsReturnedToCaller -count=1
--- FAIL: TestQuota_StorageFailureIsReturnedToCaller (0.24s)
    quota_test.go:310: expected storage error, got result &{QuotaID: QuotaType:cabin Available:0 Reserved:0 Status:exhausted WaitEstimate:0s Rejected:true Message:quota cabin exhausted: available 0, requested 10}
FAIL
FAIL	portcoord/internal/quota	0.479s
FAIL

```

stderr：

```text
warning: internal/quota/quota_test.go has type 100755, expected 100644
warning: internal/quota/quota_test.go has type 100755, expected 100644

```

## 通过条件

在额度存储不可用的触发条件下，接口必须返回服务错误且不返回额度耗尽业务结果；真正额度不足仍返回正常拒绝。定向测试、相关包测试、全量测试、race、vet 和 build 均通过；回退 gold 唯一生产修改后定向测试重新失败。
