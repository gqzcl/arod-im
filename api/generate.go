package api

// go:generate protoc -I./ -I$GOPATH/ --go_out=.  --go_opt=paths=source_relative  protocol/protocol.proto

// go:generate protoc -I./ -I$GOPATH/ --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative logic/logic.proto

// go:generate protoc -I./ -I$GOPATH/ --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative comet/comet.proto

//go:generate kratos proto client api/protocol/protocol.proto
//go:generate kratos proto client api/comet/v1/comet.proto
//go:generate kratos proto client api/logic/v1/logic.proto
