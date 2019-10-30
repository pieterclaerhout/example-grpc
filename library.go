package examplegrpc

import (
	context "context"
	fmt "fmt"
	"net"

	grpc "google.golang.org/grpc"
)

// Greeter ...
type Greeter struct {
}

// New creates new server greeter
func New() *Greeter {
	return &Greeter{}
}

// Start starts server
func (g *Greeter) Start() error {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	RegisterGreeterServer(grpcServer, g)
	grpcServer.Serve(lis)
	return nil
}

// SayHello says hello
func (g *Greeter) SayHello(ctx context.Context, r *HelloRequest) (*HelloReply, error) {
	return &HelloReply{
		Message: fmt.Sprintf("Hello, %s!", r.Name),
	}, nil
}
