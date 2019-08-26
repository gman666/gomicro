Readme
======

```go build ./cmd/gomicro/service.go```

```go build ./cmd/gomicro/client.go```

Go Version
----------
Tested with: go version go1.12.9 linux/amd64


Dependencies
------------
```go get github.com/golang/protobuf/{proto,protoc-gen-go}```

```go get github.com/micro/protoc-gen-micro```

```sudo snap install protobuf --classic```

```go get -u github.com/micro/micro```

If you are unable to successfully download all dependencies, such as:

```go: github.com/nats-io/nats.go@v1.8.2-0.20190607221125-9f4d16fe7c2d: unknown revision 9f4d16fe7c2d```

you may need to utilise goproxy and try again:

```export GOPROXY=https://goproxy.io```

Compile Proto
-------------
```protoc --proto_path=./proto --micro_out=./proto --go_out=./proto greeter.proto```