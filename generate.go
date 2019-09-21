// run `go generate` to generate files for go
//
// make .pb.go file from .proto
//go:generate protoc schema/schema.proto --go_out=plugins=grpc:.
//
// fix: move .pb.go file into parent dir
//go:generate mv schema/schema.pb.go .

package gapi
