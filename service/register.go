/**
注册服务
 */
package service

import (
  "fmt"
  "google.golang.org/grpc"
  . "github.com/r00tjimmy/ColorfulRabbit"
  "jadeRabit/consul"
  "net"
  "os"
  "time"
)

var lgx = &Logx{
  DebugFlag:   true,
}

// 放在main函数在面，程序装载运行之后， 就不会改变了


var conf = &Conf{
 os.Getenv("GOPATH") + "/src/jadeRabit/conf/env.conf",
 os.Getenv("GOPATH") + "/src/jadeRabit/conf/config.conf",
}

const (
  ConsulPort = 8500
)

func Register(srvName string, srvPort int) (srv *grpc.Server, listener net.Listener, err error) {
  //srvHost := "127.0.0.1"
  srvHost := conf.GetVal("srvHost")
  srvAddr := srvHost + ":" + fmt.Sprint(srvPort)
  //consulHost := "127.0.0.1"
  consulHost := conf.GetVal("consulHost")
  consulAddr := consulHost + ":" + fmt.Sprint(ConsulPort)
  consulToken := ""
  ttl := 15       // consul重新注册服务的时间

  // 注册服务到consul
  srv = grpc.NewServer()
  listener, err = net.Listen("tcp", srvAddr)
  CheckFatal(err, "grpc listen fatal")

  err = consul.Register(srvName, srvHost, srvPort, consulAddr, time.Second * 10, ttl, consulToken)
  CheckFatal(err, "consul register fatal")
  lgx.DebugPrint("srv " + srvName + " is running on: " + fmt.Sprint(os.Getpid()))
  return
}

func RegisterWithId(srvName string, id string, srvPort int) (srv *grpc.Server, listener net.Listener, err error) {
  //srvHost := "127.0.0.1"
  srvHost := conf.GetVal("srvHost")
  srvAddr := srvHost + ":" + fmt.Sprint(srvPort)
  //consulHost := "127.0.0.1"
  consulHost := conf.GetVal("consulHost")
  consulAddr := consulHost + ":" + fmt.Sprint(ConsulPort)
  consulToken := ""
  ttl := 15       // consul重新注册服务的时间

  // 注册服务到consul
  srv = grpc.NewServer()
  listener, err = net.Listen("tcp", srvAddr)
  CheckFatal(err, "grpc listen fatal")

  err = consul.RegisterWithId(srvName, id, srvHost, srvPort, consulAddr, time.Second * 10, ttl, consulToken)
  CheckFatal(err, "consul register fatal")
  lgx.DebugPrint("srv " + srvName + ", id " + id + " is running on: " + fmt.Sprint(os.Getpid()))
  return
}
