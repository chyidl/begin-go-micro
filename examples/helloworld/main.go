package main

import (
	"context"
	pb "github.com/asim/go-micro/examples/v3/helloworld/proto"
	"github.com/asim/go-micro/v3"
	"github.com/micro/micro/v3/cmd/protoc-gen-micro/plugin/micro"
	"log"
)

type Greeter struct {
}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("helloworld"),
		)
	service.Init()

	pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}