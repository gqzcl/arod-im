package api

//go:generate protoc -I./ -I$GOPATH/src --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative comet/comet.proto

//go:generate protoc -I./ -I$GOPATH/src --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative logic/logic.proto

//go:generate protoc -I./  -I$GOPATH/src --go_out=. --go_opt=paths=source_relative protocol/protocol.proto
