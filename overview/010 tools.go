/*
Go 自带一套完整的开发工具链，无需安装额外的构建工具。以下是最常用的 go 命令速查：
命令	功能	常用场景
go build	编译源代码	生成可执行文件，用于部署
go run	编译并运行	快速测试代码，开发调试
go test	执行测试	运行 _test.go 文件中的测试函数
go mod init	初始化模块	创建新项目时生成 go.mod 文件
go mod tidy	整理依赖	下载缺失的依赖，移除未使用的依赖
go fmt	格式化代码	统一代码风格（Go 社区强制使用）
go vet	静态分析	检查代码中的可疑构造
go get	安装依赖包	下载并安装指定的包
go doc	查看文档	在终端查看包或函数的文档
go env	查看环境变量	查看 GOPATH、GOROOT 等配置
go clean	清理构建缓存	释放磁盘空间
go fmt 是 Go 社区的一项强制性约定——所有 Go 代码都应通过 go fmt 格式化。这消除了关于代码风格的争论，让代码库保持高度一致性。大多数编辑器（VS Code、GoLand 等）都支持保存时自动运行 go fmt。
*/