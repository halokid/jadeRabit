package main

import (
  "context"
  . "github.com/r00tjimmy/ColorfulRabbit"
  proto "./proto"
  "jadeRabit/service"
)

var logx = &Logx{
  DebugFlag:  true,
}

type User struct {}

const (
  SrvPort = 8090
  SrvName = "shopUser"
)

func (u *User) Get(ctx context.Context, req *proto.Request) (rsp *proto.Response, err error) {
  userInfo := `{"uid": 1, "name": "rabbit"}`
  rsp = &proto.Response{UserInfo: userInfo}
  return rsp, nil
}

func main() {
  srv, listener, err := service.RegisterWithId()
}
