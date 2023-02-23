# grpc-template

This repo is a simple template for using a handleful of tools to provide an API that is typed using Protobuf, and exposes a JSON/REST interface.

The protobuf code for this project is stored in `api/proto/v1/api`. The main files of interest for this project include:
- main.go
- api/proto/v1/api/* (the proto definitions)
- api/buf.gen.yaml (the codegen config)

## Requirements

Protobuf Tools:
- Buf

Go Libraries:
```sh
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

Misc:
- Swagger (Optional)

## Building & Running
```sh
# run codegen
go generate ./...

# run the project
go run main.go
```

## Useful Links
- https://github.com/grpc-ecosystem/grpc-gateway
- https://grpc-ecosystem.github.io/grpc-gateway/docs/operations/inject_router/
- https://web.archive.org/web/20201112010739/https://coreos.com/blog/grpc-protobufs-swagger.html
- https://github.com/googleapis/googleapis/blob/master/google/api/http.proto (docs on path option syntax)