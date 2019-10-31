package main

import (
	"context"
	"fmt"

	examplegrpc "github.com/pieterclaerhout/example-grpc"
	"github.com/pieterclaerhout/example-grpc/versioninfo"
	"github.com/pieterclaerhout/go-log"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Project: " + versioninfo.ProjectName)
	fmt.Println("Description: " + versioninfo.ProjectDescription)
	fmt.Println("Copyright: " + versioninfo.ProjectCopyright)
	fmt.Println("Version: " + versioninfo.Version)
	fmt.Println("Revision: " + versioninfo.Revision)
	fmt.Println("Branch: " + versioninfo.Branch)

	go func() {

		log.Info("Starting gRPC server")

		g := examplegrpc.New()
		err := g.Start()
		log.CheckError(err)

	}()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	log.CheckError(err)
	defer conn.Close()

	client := examplegrpc.NewGreeterClient(conn)

	result, err := client.SayHello(context.Background(), &examplegrpc.HelloRequest{
		Name: "Pieter",
	})
	log.CheckError(err)

	log.WarnDump(result.GetMessage(), "result")

}
