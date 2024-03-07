# Build the proto files for go
#
#protoc --go_out=./internal \
#       --go-grpc_out=./internal \
#       ./api/**/*.proto

buf generate --template ./config/gateway.gen.yaml --path ./api



# Build the proto files for gRPC-web. Respect
# the go package name for the generated files
npx buf generate --template ./config/dashboard.gen.yaml

