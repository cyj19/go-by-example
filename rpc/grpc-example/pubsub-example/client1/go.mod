module github.com/vagaryer/go-by-example/rpc/grpc-example/pubsub-example/client1

go 1.16

replace github.com/vagaryer/go-by-example/rpc/grpc-example/proto/pubsub => ../../proto/pubsub

require (
	github.com/vagaryer/go-by-example/rpc/grpc-example/proto/pubsub v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.42.0
)
