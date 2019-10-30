package main

import (
	"fmt"

	examplegrpc "github.com/pieterclaerhout/example-grpc"
	"github.com/pieterclaerhout/example-grpc/versioninfo"
	"github.com/pieterclaerhout/go-log"
)

func main() {
	fmt.Println("Project: " + versioninfo.ProjectName)
	fmt.Println("Description: " + versioninfo.ProjectDescription)
	fmt.Println("Copyright: " + versioninfo.ProjectCopyright)
	fmt.Println("Version: " + versioninfo.Version)
	fmt.Println("Revision: " + versioninfo.Revision)
	fmt.Println("Branch: " + versioninfo.Branch)

	g := examplegrpc.New()
	err := g.Start()
	log.CheckError(err)

}
