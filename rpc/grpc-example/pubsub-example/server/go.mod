module github.com/vagaryer/go-by-example/rpc/grpc-example/pubsub-example/server

go 1.16

replace github.com/vagaryer/go-by-example/rpc/grpc-example/proto/pubsub => ../../proto/pubsub

require (
	github.com/docker/docker v20.10.10+incompatible
	github.com/vagaryer/go-by-example/rpc/grpc-example/proto/pubsub v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.42.0
)
