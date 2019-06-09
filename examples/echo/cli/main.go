package main

import (
  "context"
  "fmt"
  consulApi "github.com/hashicorp/consul/api"
  . "github.com/r00tjimmy/ColorfulRabbit"
  "google.golang.org/grpc"
  proto "../srv/proto"
  "os"
)

func main() {
  config := consulApi.Config{
    Address:"http://127.0.0.1:8500",
  }

  cli, err := consulApi.NewClient(&config) //非默认情况下需要设置实际的参数
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(0)
  }

  services, _, err := cli.Catalog().Service("echo", "", nil)
  CheckFatal(err, "can not get service echo from consul")

  callSrv(services)
}

// 调用srv服务
func callSrv(services []*consulApi.CatalogService) {
  if len(services) > 0 {
    conn, err := grpc.Dial(fmt.Sprintf("%s:%d", services[0].ServiceAddress, services[0].ServicePort), grpc.WithInsecure())
    CheckError(err)
    c := proto.NewEchoClient(conn)
    req := &proto.EchoRequest{ReqMsg: "Hello JadeRabit"}
    rsp, err := c.Pong(context.Background(), req)
    fmt.Println(rsp.RspMsg)
  }
}

