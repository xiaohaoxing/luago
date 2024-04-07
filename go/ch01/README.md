## golang 环境说明

本书当前没有 go module 的配置，在新版本的 golang 中需要设置 module 才能正确运行。

1. 在 `$LUAGO/go/ch01` 路径下执行 `go mod init luago` 初始化模块。
2. 执行 `go mod tidy` 下载依赖性（暂时为空）。
3. 执行 `go build src/luago/main.go` 编译生成可执行文件。
4. 执行 `./main` 运行可执行文件，得到运行结果。

