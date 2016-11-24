package server

import (
  "fmt"
  "github.com/labstack/gommon/log"
  "github.com/labstack/echo"
  "github.com/Synapse791/census/config"
  "github.com/coreos/etcd/client"
)

var server *echo.Echo
var etcdApi client.KeysAPI
var appConfig *config.CensusConfig

func Run(c *config.CensusConfig) {
  appConfig = c
  server = echo.New()
  server.Logger.SetLevel(log.INFO)

  var etcdErr error
  var etcdClient client.Client

  etcdClient, etcdErr = client.New(client.Config{
    Endpoints: appConfig.Etcd.Endpoints,
  })
  if etcdErr != nil {
    server.Logger.Fatal(etcdErr.Error())
  }

  etcdApi = client.NewKeysAPI(etcdClient)

  registerRoutes()

  server.Logger.Fatal(server.Start(fmt.Sprintf("%s:%d", appConfig.Server.Addr, appConfig.Server.Port)))
}

func registerRoutes() {
  server.GET("/apps", getAppsHandler)
  server.PUT("/apps", putAppsHandler)
  server.GET("/hosts", getHostsHandler)
  server.PUT("/hosts", putHostsHandler)
}
