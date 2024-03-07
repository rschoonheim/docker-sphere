if ! buf lint; then
  echo "buf lint failed"
  exit 1
fi


# Build the server protos
#
protoc --go_out=./internal/pkg \
       --go-grpc_out=./internal/pkg \
       ./proto/v1/user.proto
