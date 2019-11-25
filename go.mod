module github.com/eviot/grpc-api-go

go 1.13

require (
	github.com/eviot/log v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.3.2
	google.golang.org/grpc v1.23.1
)

replace github.com/eviot/log => ../log
