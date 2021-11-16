### gRPC使用的步骤
1. 编写proto文件，使用protoc-gen-go内置的gRPC插件生成gRPC代码，命令如下
```
protoc --go_out=plugins=grpc:. hello.proto
```
2. 在服务端重新实现gRPC生成的服务接口，创建gRPC，注册服务，启动服务
3. 在客户端使用gRPC拨号，使用gRPC生成的代码创建客户端，调用方法