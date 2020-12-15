package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	srvHello "lemon_service/proto/hello"
	"log"
	"strconv"
	"strings"
)

type Hello struct {
    client srvHello.HelloService
}

func (h *Hello) SayHi(ctx context.Context,req *api.Request,rsp *api.Response) error {
	log.Println("接收到 Hello.SayHi API的请求")
	name,ok := req.Get["name"]
	if !ok || len(name.Values) == 0{
		return errors.BadRequest("go.micro.api.Hello","名字不能为空")
	}

	address,ok := req.Get["address"]
	if !ok || len(address.Values) == 0{
		return errors.BadRequest("go.micro.api.Hello","城市名不能为空")
	}

	//将参数交给底层处理
	response,err := h.client.SayHi(ctx,&srvHello.Request{
		Name:    strings.Join(name.Values," "),
		Address: strings.Join(address.Values," "),
	})

	if err != nil {
		return err
	}
	//处理成功，则返回
	rsp.StatusCode = 200
	b,_:= json.Marshal(map[string]string{
		"message":response.Ret,
	})
	rsp.Body = string(b)
	return nil
}

func (h *Hello) Add(ctx context.Context,req *api.Request,rsp *api.Response) error {
	log.Println("接收到 Hello.Add API的请求")
	num1,ok := req.Get["num1"]
	if !ok || len(num1.Values) == 0{
		return errors.BadRequest("go.micro.api.Hello","num1不能为空也")
	}

	num2,ok := req.Get["num2"]
	if !ok || len(num2.Values) == 0{
		return errors.BadRequest("go.micro.api.Hello","num2不能为空")
	}

	n1,err := parseToInt(num1.Values[0])
	if err != nil {
		return errors.BadRequest("go.micro.api.Hello",
			"num1(%s)必须为一个数值",num1.Values[0])
	}
	n2,err := parseToInt(num2.Values[0])
	if err != nil {
		return errors.BadRequest("go.micro.api.Hello",
			"num2(%s)必须为一个数值",num2.Values[0])
	}

	//将参数交给底层处理
	response,err := h.client.Add(ctx,&srvHello.Params{
		Num1: n1,
		Num2: n2,
	})

	if err != nil {
		return err
	}
	//处理成功，则返回
	rsp.StatusCode = 200
	b,_:= json.Marshal(map[string]string{
		"message":fmt.Sprintf("%d",response.Res),
	})
	rsp.Body = string(b)
	return nil
}

func parseToInt(v string)(int32,error)  {
	var i int
	var err error
	 i,err  = strconv.Atoi(v)
	 return int32(i),err
}

//http://127.0.0.1:8080/Hello/Hello/add?num1=11&num2=100
//http://127.0.0.1:8080/Hello/Hello/SayHi?name=jackm&adrress=西安
func main()  {
	//先自己注册到注册中心上
	service := micro.NewService(micro.Name("go.micro.api.Hello"))
	//初始化参数
	service.Init()

	//将请求转发给go.micro.srv.Hello服务进行处理
	_ = service.Server().Handle(
		service.Server().NewHandler(
			&Hello{client: srvHello.NewHelloService("go.micro.srv.Hello",service.Client())}))
	//运行服务
	if err := service.Run(); err != nil {
		panic(err)
	}
}