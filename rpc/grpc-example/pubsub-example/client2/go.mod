module github.com/cyj19/go-by-example/rpc/grpc-example/pubsub-example/client2

go 1.16

replace github.com/cyj19/go-by-example/rpc/grpc-example/proto/pubsub => ../../proto/pubsub

require (
	github.com/cyj19/go-by-example/rpc/grpc-example/proto/pubsub v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.42.0
)
