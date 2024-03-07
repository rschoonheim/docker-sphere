# Build the proto files for go
#
protoc --go_out=./packages \
       --go-grpc_out=./packages \
       ./api/*.proto

# Build the proto files for gRPC-web
#
protoc --js_out=import_style=commonjs:./packages \
       --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./packages \
       ./api/*.proto

