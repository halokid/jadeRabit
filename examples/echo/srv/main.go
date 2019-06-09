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
  SrvPort =  8089
  SrvName =  "echo"
)

/**
func (e *Echo) Pong(ctx context.Context, req *proto.EchoRequest) (rsp *proto.EchoResponse, err error) {
  //rsp.RspMsg = req.ReqMsg
  fmt.Println(req.ReqMsg)
  return rsp, nil
}
**/

func (e *Echo) Pong(ctx context.Context, req *proto.EchoRequest) (rsp *proto.EchoResponse, err error) {
  logx.DebugPrint("req.ReqMsg -----------", req.ReqMsg)
  rsp = &proto.EchoResponse{RspMsg: "get " + req.ReqMsg}
  return rsp, nil
}


func main() {
  srv, listener, err := service.Register(SrvName, SrvPort)
  CheckFatal(err,"srv register error")
  proto.RegisterEchoServer(srv, &Echo{})
  err = srv.Serve(listener)
  CheckFatal(err, "srv listen error")
}
