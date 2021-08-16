package main

import (
	"context"
	pb "github.com/chyidl/begain-go-micro/examples/ch1/proto"
	"github.com/asim/go-micro/v3"
	"log"
)

type Greeter struct {}

func (g *Greeter) Hello(_ context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// 创建新的服务
	service := micro.NewService(micro.Name("ch1.begain-go-micro"),)

	// 初始化方法
	service.Init()

	// 注册服务
	pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}