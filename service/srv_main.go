package main

import (
	"context"
	"github.com/micro/go-micro"
	srcHello "lemon_service/proto/hello"
)

type Hello struct {}

func (h *Hello) SayHi(ctx context.Context,req *srcHello.Request,rsp *srcHello.Response) error {
	rsp.Ret = "你好,"+req.Name+",欢迎到 "+req.Address
	return nil
}

func (h *Hello) Add(ctx context.Context,params *srcHello.Params,result *srcHello.Result) error  {
	result.Res = params.Num1+params.Num2
	return nil
}


// start cmd: go run srv_main.go --registry consul --registry_address 192.168.0.50:8500
// list services
/**
microv113 --registry  consul --registry_address  192.168.0.52:8500  list services
consul
etcd-192.168.0.72
go.micro.srv.Hello
 */
func main()  {

	//create a service
	service := micro.NewService(
		micro.Name("go.micro.srv.Hello"))

	//service params init
	service.Init()

	//registry method
	err := srcHello.RegisterHelloHandler(service.Server(),new(Hello))
	if err != nil {
		panic(err)
	}

	//start service
	if err := service.Run(); err != nil {
		panic(err)
	}
}
