package grpcclient

import (
	"go-boilerplate/internals/config"
	api_v1 "go-boilerplate/pkg/api/v1"
	"log"

	"google.golang.org/grpc"
)

// HTTPGRPCClient ...
type HTTPGRPCClient struct {
	Config   config.Configuration
	PingPong api_v1.PingPongServiceClient
}

// Connect ...
func (client *HTTPGRPCClient) Connect() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("0.0.0.0:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("Connect to service on", "0.0.0.0:3000")

	client.PingPong = api_v1.NewPingPongServiceClient(conn)
}

// NewHTTPGRPCClient ...
func NewHTTPGRPCClient(config config.Configuration) *HTTPGRPCClient {
	// fmt.Printf("new grpc client: %+v\n", config)
	grpcClient := HTTPGRPCClient{
		Config: config,
	}
	grpcClient.Connect()
	return &grpcClient
}
