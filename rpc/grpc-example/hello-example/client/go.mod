module github.com/vagaryer/go-by-example/rpc/grpc-example/hello-example/client

go 1.16

// 导入本地包
replace github.com/vagaryer/go-by-example/rpc/grpc-example/proto/hello => ../../proto/hello

require (
	github.com/vagaryer/go-by-example/rpc/grpc-example/proto/hello v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.42.0
)
