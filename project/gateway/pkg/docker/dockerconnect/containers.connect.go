// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/gateway/docker/v1/containers.proto

package gateway_dockerconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	docker "docker_sphere/project/gateway/pkg/docker"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// DockerServiceName is the fully-qualified name of the DockerService service.
	DockerServiceName = "proto.gateway.docker.v1.DockerService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// DockerServiceListContainersProcedure is the fully-qualified name of the DockerService's
	// ListContainers RPC.
	DockerServiceListContainersProcedure = "/proto.gateway.docker.v1.DockerService/ListContainers"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	dockerServiceServiceDescriptor              = docker.File_proto_gateway_docker_v1_containers_proto.Services().ByName("DockerService")
	dockerServiceListContainersMethodDescriptor = dockerServiceServiceDescriptor.Methods().ByName("ListContainers")
)

// DockerServiceClient is a client for the proto.gateway.docker.v1.DockerService service.
type DockerServiceClient interface {
	ListContainers(context.Context, *connect.Request[docker.ListContainersRequest]) (*connect.Response[docker.ListContainersResponse], error)
}

// NewDockerServiceClient constructs a client for the proto.gateway.docker.v1.DockerService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDockerServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) DockerServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &dockerServiceClient{
		listContainers: connect.NewClient[docker.ListContainersRequest, docker.ListContainersResponse](
			httpClient,
			baseURL+DockerServiceListContainersProcedure,
			connect.WithSchema(dockerServiceListContainersMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// dockerServiceClient implements DockerServiceClient.
type dockerServiceClient struct {
	listContainers *connect.Client[docker.ListContainersRequest, docker.ListContainersResponse]
}

// ListContainers calls proto.gateway.docker.v1.DockerService.ListContainers.
func (c *dockerServiceClient) ListContainers(ctx context.Context, req *connect.Request[docker.ListContainersRequest]) (*connect.Response[docker.ListContainersResponse], error) {
	return c.listContainers.CallUnary(ctx, req)
}

// DockerServiceHandler is an implementation of the proto.gateway.docker.v1.DockerService service.
type DockerServiceHandler interface {
	ListContainers(context.Context, *connect.Request[docker.ListContainersRequest]) (*connect.Response[docker.ListContainersResponse], error)
}

// NewDockerServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDockerServiceHandler(svc DockerServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	dockerServiceListContainersHandler := connect.NewUnaryHandler(
		DockerServiceListContainersProcedure,
		svc.ListContainers,
		connect.WithSchema(dockerServiceListContainersMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/proto.gateway.docker.v1.DockerService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case DockerServiceListContainersProcedure:
			dockerServiceListContainersHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedDockerServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedDockerServiceHandler struct{}

func (UnimplementedDockerServiceHandler) ListContainers(context.Context, *connect.Request[docker.ListContainersRequest]) (*connect.Response[docker.ListContainersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.gateway.docker.v1.DockerService.ListContainers is not implemented"))
}