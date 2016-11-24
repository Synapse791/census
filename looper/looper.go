package looper

import (
  "fmt"
  "time"
  "strings"
  "golang.org/x/net/context"
  "github.com/Synapse791/census/config"
  "github.com/Synapse791/census/console"
  "github.com/Synapse791/census/utils"
  "github.com/coreos/etcd/client"
)

var appConfig *config.CensusConfig
var etcdApi client.KeysAPI

func Run(c *config.CensusConfig) {
  appConfig = c

  etcdClient, etcdErr := client.New(client.Config{
    Endpoints: appConfig.Etcd.Endpoints,
  })
  if etcdErr != nil {
    console.ErrorAndExit(etcdErr.Error())
  }

  etcdApi = client.NewKeysAPI(etcdClient)

  console.Line("Starting worker loop")

  console.Line("Running iteration")
  if err := iteration(); err != nil {
    console.Error("Failed iteration: %s", err.Error())
  }

  for _ = range time.Tick(time.Second * 10) {
    if err := iteration(); err != nil {
      console.Error("Failed iteration: %s", err.Error())
    }
  }
}

func iteration() error {
  var etcdApps *client.Response
  var etcdHosts *client.Response
  var etcdErr error

  etcdApps, etcdErr = etcdApi.Get(context.Background(), fmt.Sprintf("%s/census/apps", appConfig.Etcd.Prefix), &client.GetOptions{Recursive: true})
  if etcdErr != nil { return etcdErr }

  var apps []LooperApp

  for _, etcdApp := range etcdApps.Node.Nodes {
    var app LooperApp
    app.Name = utils.GetLastPart(etcdApp.Key, "/")

    for _, node := range etcdApp.Nodes {
      if strings.Contains(node.Key, "port") {
        app.Port = node.Value
      } else if strings.Contains(node.Key, "hosts") && node.Nodes != nil {
        for _, etcdAppHost := range node.Nodes {
          var host LooperHost
          host.Name = utils.GetLastPart(etcdAppHost.Key, "/")
          host.IpAddress = etcdAppHost.Value
          app.Hosts = append(app.Hosts, host)
        }
      }
    }

    apps = append(apps, app)
  }

  etcdHosts, etcdErr = etcdApi.Get(context.Background(), fmt.Sprintf("%s/census/hosts", appConfig.Etcd.Prefix), &client.GetOptions{Recursive: true})
  if etcdErr != nil { return etcdErr }

  for _, app := range apps {
    port := app.Port

    for _, host := range app.Hosts {
      if ! utils.ConnectionOk(host.IpAddress, port) {
        console.Warning("Removing host %s from app %s", host.Name, app.Name)
        hostKey := fmt.Sprintf("%s/census/apps/%s/hosts/%s", appConfig.Etcd.Prefix, app.Name, host.Name)
        etcdHosts, etcdErr = etcdApi.Delete(context.Background(), hostKey, nil)
        if etcdErr != nil {
          console.Error("Failed to remove host %s from app %s", host.Name, app.Name)
          return etcdErr
        }
        console.Warning("Removed host %s from app %s successfully", host.Name, app.Name)
      }
    }

    for _, etcdHost := range etcdHosts.Node.Nodes {
      if utils.ConnectionOk(etcdHost.Value, port) {
        etcdHostName := utils.GetLastPart(etcdHost.Key, "/")

        alreadyExists := false

        for _, appHost := range app.Hosts {
          if appHost.Name == etcdHostName {
            alreadyExists = true
          }
        }
        if alreadyExists {
          continue
        }

        console.Line("Adding host %s to app %s...", etcdHostName, app.Name)
        hostKey := fmt.Sprintf("%s/census/apps/%s/hosts/%s", appConfig.Etcd.Prefix, app.Name, etcdHostName)
        etcdHosts, etcdErr = etcdApi.Set(context.Background(), hostKey, etcdHost.Value, nil)
        if etcdErr != nil {
          console.Error("Failed to add host %s to app %s", etcdHostName, app.Name)
          return etcdErr
        }

        console.Success("Added host %s to app %s successfully", etcdHostName, app.Name)
      }
    }
  }

  return nil

  // hosts, etcdErr = etcdApi.Get(context.Background(), fmt.Sprintf("%s/census/hosts"), &client.GetOptions{Recursive: true})
  // if etcdErr != nil { return etcdErr }
  //
  // for _, app := range apps.Node.Nodes {
  //
  // }
}
