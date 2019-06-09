package main

import (
  "context"
  "fmt"
  consulApi "github.com/hashicorp/consul/api"
  . "github.com/r00tjimmy/ColorfulRabbit"
  "google.golang.org/grpc"
  proto "../srv1/proto"
  "os"
)

var logx = &Logx{
  DebugFlag:   true,
}

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
  logx.DebugPrint(services)
  for _, serItem := range services {
    logx.DebugPrint(serItem.ServiceAddress,  ":",  serItem.ServicePort)
  }

  callSrv(services)
}

// 调用srv服务
func callSrv(services []*consulApi.CatalogService) {
  if len(services) > 0 {
    svcIdx := RandInt(0, len(services))
    // 随机算法获取服务的地址端口
    conn, err := grpc.Dial(fmt.Sprintf("%s:%d", services[svcIdx].ServiceAddress, services[svcIdx].ServicePort), grpc.WithInsecure())
    CheckError(err)
    c := proto.NewEchoClient(conn)
    req := &proto.EchoRequest{ReqMsg: "Hello JadeRabit"}
    rsp, err := c.Pong(context.Background(), req)
    fmt.Println(rsp.RspMsg)
  }
}






