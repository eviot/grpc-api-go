all: generate

# generate files for go
generate:
	# make .pb.go file from .proto
	protoc -I./schema --go_out=plugins=grpc:. schema/*.proto
