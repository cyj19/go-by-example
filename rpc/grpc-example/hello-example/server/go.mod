module github.com/cyj19/go-by-example/rpc/grpc-example/hello-example/server

go 1.16

// 导入本地包
replace github.com/cyj19/go-by-example/rpc/grpc-example/proto/hello => ../../proto/hello

require (
	github.com/cyj19/go-by-example/rpc/grpc-example/proto/hello v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.42.0
)
