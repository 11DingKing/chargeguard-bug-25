# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

上游取消整改提醒请求后，短信确实停止发送，service 却继续写入“发送成功”审计，通知游标也向前推进，台账因此出现未送达记录。目标仓库代码保持不变，请定位取消发生后仍继续执行的控制流，说明涉及文件和符号，并用运行证据还原这条错误审计如何产生。

## 含 Bug 版本

- 仓库：11DingKing/chargeguard-bug-25
- 仓库地址：https://github.com/11DingKing/chargeguard-bug-25.git
- parent SHA：78b097b307d54e66cc531f9f99da3758ca278afb

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/chargeguard-bug-25.git bug-repro
cd bug-repro
git checkout --detach 78b097b307d54e66cc531f9f99da3758ca278afb
go test ./internal/charging -run TestTaskBehavior -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/charging -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.00s)
    task_behavior_test.go:23: err=<nil>
FAIL
FAIL	chargeguard/internal/charging	0.036s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/charging -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.00s)
    task_behavior_test.go:23: err=<nil>
FAIL
FAIL	chargeguard/internal/charging	0.002s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

根因结论必须准确写明出问题的 Go 文件、具体符号和完整失效机制，并由实际复现、源码调查和验证证据支撑；调查结束时目标仓库代码、测试和配置零改动。
