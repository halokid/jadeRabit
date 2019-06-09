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
  SrvPort1 =  8089
  SrvName1 =  "echo"
  SrvId1 =  "echo1"
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
  rsp = &proto.EchoResponse{RspMsg: "srv1 get " + req.ReqMsg}
  return rsp, nil
}


func main() {
  srv1, listener1, err := service.RegisterWithId(SrvName1, SrvId1, SrvPort1)
  CheckFatal(err,"srv register error")
  proto.RegisterEchoServer(srv1, &Echo{})
  err = srv1.Serve(listener1)
  CheckFatal(err, "srv listen error")
}




