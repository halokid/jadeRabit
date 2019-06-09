package main

import (
  "context"
  "fmt"
  "github.com/gin-gonic/gin"
  . "github.com/r00tjimmy/ColorfulRabbit"
  consulApi "github.com/hashicorp/consul/api"
  "google.golang.org/grpc"
  "os"
  proto "../srv/proto"
)


var logx = &Logx{
  DebugFlag:  true,
}


func SrvRet(c *gin.Context) {
  srvRet := ""
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

  if len(services) > 0 {
    conn, err := grpc.Dial(fmt.Sprintf("%s:%d", services[0].ServiceAddress, services[0].ServicePort), grpc.WithInsecure())
    CheckError(err)

    defer conn.Close()

    c := proto.NewEchoClient(conn)
    req := &proto.EchoRequest{ReqMsg: "Hello JadeRabit"}
    rsp, err := c.Pong(context.Background(), req)
    fmt.Println(rsp.RspMsg)
    //return rsp.RspMsg
    srvRet = rsp.RspMsg
  }
  //return ""
  c.JSON(200, gin.H{
      "code": 1,
      "msg":  srvRet,
  })
}



func Perman(c *gin.Context) {
  c.JSON(200, gin.H{
      "code": 1,
      "msg":  "success",
  })
}


func main()  {
  r := gin.Default()

  r.GET("/echo", SrvRet)
  r.GET("/perman", Perman)

  _ = r.Run("0.0.0.0:8090")
}






