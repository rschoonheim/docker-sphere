# Build the proto files for go
#
protoc --go_out=./internal \
       --go-grpc_out=./internal \
       ./api/**/*.proto




# Build the proto files for gRPC-web. Respect
# the go package name for the generated files
protoc -I=./api \
       --grpc-web_out=import_style=typescript,mode=grpcweb:./internal \
       ./api/**/*.proto


