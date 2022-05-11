package pb

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto-v1.proto
//go:generate protoc --go_out=./v2 --go_opt=paths=source_relative --go-grpc_out=./v2 --go-grpc_opt=paths=source_relative proto-v2.proto
