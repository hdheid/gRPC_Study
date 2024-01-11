# gRPC的学习

1. 需要下载protobuf的编译器，在官网下载好之后，需要将 bin 目录配置到 path 环境变量中，命令行中输入命令 `protoc --version` 出现版本号即表示安装成功。

2. gRPC需要使用 `go get google.golang.org/grpc` 命令来下载 gRPC 的核心库。

3. 还需要下载代码生成工具，使用
   - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
   - ` go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

4. 上面两个 go install 命令会安装两个 exe 文件到 ￥GOPATH/bin 目录下。

## proto 文件的编写
proto文件，是分别写在每一个需要连接的程序里面，类似于结构文档，是一种约束，表示所有的程序都需要依据这个文件进行通信。

例如在 hello-server