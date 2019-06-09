/**
发现服务
 */
package service

import (
  "fmt"
  consulApi "github.com/hashicorp/consul/api"
  "google.golang.org/grpc"
  . "github.com/r00tjimmy/ColorfulRabbit"
)

type SrvDisver struct {
  RegCfg      consulApi.Config
}

// 自动发现服务（获取rpc服务的连接句柄）
func (sd *SrvDisver) Get(srvName string) *grpc.ClientConn{
  cli, err := consulApi.NewClient(&sd.RegCfg)
  CheckFatal(err, "srv discover regconfig error")
  services, _, err := cli.Catalog().Service(srvName, "", nil)
  CheckFatal(err, "can not get service " + srvName + " from srv reg center")
  // 选择的服务节点默认为第一个
  srvIdx := 0
  // 服务不止一个节点的时候， 调用选择节点规则算法
  if len(services) > 0 {
    srvIdx = GetSrvsSigle(len(services))
  }
  conn, err := grpc.Dial(fmt.Sprintf("%s:%d", services[srvIdx].ServiceAddress, services[srvIdx].ServicePort), grpc.WithInsecure())
  return conn
}


// 自动发现服务的算法规则统一入口
func GetSrvsSigle(sLen int) int {
  srvIdx := RandInt(0, sLen)
  return srvIdx
}





