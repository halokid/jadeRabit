package main

import (
  proto "./proto"
  "context"
  . "github.com/r00tjimmy/ColorfulRabbit"
  "jadeRabit/service"
)

var logx = &Logx{
  DebugFlag:  true,
}

type Echo struct {}


const (
  SrvPort2 =  8090
  SrvName2 =  "echo"
  SrvId2 = "echo2"
)

func (e *Echo) Pong(ctx context.Context, req *proto.EchoRequest) (rsp *proto.EchoResponse, err error) {
  logx.DebugPrint("req.ReqMsg -----------", req.ReqMsg)
  rsp = &proto.EchoResponse{RspMsg: "srv2 get " + req.ReqMsg}
  return rsp, nil
}


func main() {
  srv2, listener2, err := service.RegisterWithId(SrvName2, SrvId2, SrvPort2)
  CheckFatal(err,"srv register error")
  proto.RegisterEchoServer(srv2, &Echo{})
  err = srv2.Serve(listener2)
  CheckFatal(err, "srv listen error")
}




