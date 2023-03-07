module github.com/eviot/grpc-api-go

go 1.13

require (
	github.com/eviot/log v0.1.3
	github.com/golang/protobuf v1.3.2
	github.com/vmihailenco/msgpack/v4 v4.2.1
	golang.org/x/net v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20191223191004-3caeed10a8bf // indirect
	google.golang.org/grpc v1.26.0
)

replace github.com/eviot/log => ../log
