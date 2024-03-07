# Build the proto files for go
#
protoc --go_out=./internal \
       --go-grpc_out=./internal \
       ./api/**/*.proto




# Build the proto files for gRPC-web. Respect
# the go package name for the generated files
protoc -I=./api \
       --js_out=import_style=commonjs,binary:./web/src/proto \
       --grpc-web_out=import_style=typescript,mode=grpcwebtext:./web/src/proto \
       ./api/**/*.proto


