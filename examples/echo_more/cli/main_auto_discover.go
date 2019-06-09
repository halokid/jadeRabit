package main

import (
  "context"
  "fmt"
  consulApi "github.com/hashicorp/consul/api"
  . "github.com/r00tjimmy/ColorfulRabbit"
  "jadeRabit/service"
  proto "../srv1/proto"
)

//var logx2 = &Logx{
//  DebugFlag:   true,
//}

func main() {

  var srvDicver service.SrvDisver

  config := consulApi.Config{
    Address:"http://127.0.0.1:8500",
  }

  srvDicver.RegCfg = config
  conn := srvDicver.Get("echo")
  c := proto.NewEchoClient(conn)
  req := &proto.EchoRequest{ReqMsg: "Hello Jade"}
  rsp, err := c.Pong(context.Background(), req)
  CheckFatal(err)
  fmt.Println(rsp.RspMsg)
}







