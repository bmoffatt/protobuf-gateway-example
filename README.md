# protobuf-gateway-example

```
brew install golang protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# after modifying pb/proto-v2.proto
go generate -x ./...

# run the servers
go run cmd/gateway-server/main.go &
go run cmd/target-server/main.go &

# run the client
go run cmd/client/main.go
```

# ???

`pb/proto-v1.proto` and `pb/proto-v2.proto` are nearly the same file, with only additional request fields added to `proto-v2.proto`. Model code generated for `v2` is placed in a separate import path. `cmd/client` and `cmd/target-server` use the `/v2` import path, and `cmd/gateway-server` uses the bare import path.

Project demonstrates that when sharing nested protobuf messages, an intermediate server does not need to be re-compiled in-order to receive forward types from an `vN+1` client and `vN+1` server dependency
