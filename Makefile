all: generate

# generate files for go
generate:
	# make .pb.go file from .proto
	protoc schema/schema.proto --go_out=plugins=grpc:.

	# fix: move .pb.go file into parent dir
	mv schema/schema.pb.go .
