Readme
======

```go build ./cmd/gomicro/service.go```

```go build ./cmd/gomicro/client.go```

Dependencies
------------
```go get github.com/micro/protoc-gen-micro```

```go get github.com/golang/protobuf/proto-gen-go```

Compile Proto
-------------
```protoc --proto_path=./proto --micro_out=./proto --go_out=./proto greeter.proto```