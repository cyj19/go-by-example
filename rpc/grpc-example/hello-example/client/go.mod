module github.com/vagaryer/go-by-example/hello-client

go 1.16

replace hello => ../../proto/hello

require (
	google.golang.org/grpc v1.40.0
	hello v0.0.0-00010101000000-000000000000
)
