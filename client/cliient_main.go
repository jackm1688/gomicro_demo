package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	srvHello "lemon_service/proto/hello"
)

func main()  {

	//先将自己注册到注册中心去
	service := micro.NewService(micro.Name("go.micro.srv.clent"))

	//初始参数
	service.Init()

	//创建 Hello对象客户端实例
	client :=  srvHello.NewHelloService("go.micro.srv.Hello",service.Client())

	//通过rpc服务调用Hello的SayHi 和Add方法
	rsp,err := client.SayHi(context.Background(),&srvHello.Request{
		Name:    "柠檬酱",
		Address: "西安",
	})

	if err != nil {
		fmt.Printf("call SayHi error:%v\n",err)
	}else {
		fmt.Println("ret:",rsp.Ret)
	}

	rsp2,err := client.Add(context.TODO(),&srvHello.Params{
		Num1: 100,
		Num2: 201,
	})

	if err != nil {
		fmt.Printf("call Add error:%v\n",err)
	}else {
		fmt.Println("ret:",rsp2.Res)
	}



}
