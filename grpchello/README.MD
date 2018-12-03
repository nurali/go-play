Reference: https://grpc.io/docs/quickstart/go.html

Generate `helloworld.pb.go` file:
```
protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
```
