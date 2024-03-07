# Build protos
#
protoc --go_out=./internal \
       --go-grpc_out=./internal \
       ./protos/*.proto
