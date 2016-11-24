package server

import (
  "fmt"
  "net/http"
  "strings"
  "golang.org/x/net/context"
  "github.com/labstack/echo"
  "github.com/coreos/etcd/client"
)

func getAppsHandler(c echo.Context) error {
  apps, err := etcdApi.Get(context.Background(), fmt.Sprintf("%s/census/apps", appConfig.Etcd.Prefix), &client.GetOptions{Recursive: true})
  if err != nil {
    server.Logger.Error(err.Error())
    return c.String(http.StatusInternalServerError, "Internal server error")
  }

  responseList := make(map[string]string)

  for _, app := range apps.Node.Nodes {
    keyParts := strings.Split(app.Key, "/")
    appName := keyParts[len(keyParts) - 1]
    portNode, portErr := etcdApi.Get(context.Background(), app.Key + "/port", &client.GetOptions{Recursive: true})
    if portErr != nil {
      server.Logger.Error(portErr.Error())
      return c.String(http.StatusInternalServerError, "Internal server error")
    }

    responseList[appName] = portNode.Node.Value
  }

  return c.JSON(http.StatusOK, responseList)
}

func putAppsHandler(c echo.Context) error {
  req := new(PutAppRequest)

  if err := c.Bind(req); err != nil {
    return err
  }

  if req.Name == "" {
    return c.String(http.StatusBadRequest, "The name field is required")
  }

  if req.Port == "" {
    return c.String(http.StatusBadRequest, "The port field is required")
  }

  if _, err := etcdApi.Set(context.Background(),  fmt.Sprintf("%s/census/apps/%s/port", appConfig.Etcd.Prefix, req.Name), req.Port, &client.SetOptions{}); err != nil {
    server.Logger.Error(err.Error())
    return c.String(http.StatusInternalServerError, "Internal server error")
  }

  return c.JSON(http.StatusCreated, req)
}

func getHostsHandler(c echo.Context) error {
  hosts, err := etcdApi.Get(context.Background(), fmt.Sprintf("%s/census/hosts", appConfig.Etcd.Prefix), &client.GetOptions{Recursive: true})
  if err != nil {
    server.Logger.Error(err.Error())
    return c.String(http.StatusInternalServerError, "Internal server error")
  }

  responseList := make(map[string]string)

  for _, host := range hosts.Node.Nodes {
    keyParts := strings.Split(host.Key, "/")
    hostName := keyParts[len(keyParts) - 1]

    responseList[hostName] = host.Value
  }

  return c.JSON(http.StatusOK, responseList)
}

func putHostsHandler(c echo.Context) error {
  req := new(PutHostRequest)

  if err := c.Bind(req); err != nil {
    return err
  }

  if req.Name == "" {
    return c.String(http.StatusBadRequest, "The name field is required")
  }

  if req.IpAddress == "" {
    return c.String(http.StatusBadRequest, "The ip field is required")
  }

  if _, err := etcdApi.Set(context.Background(),  fmt.Sprintf("%s/census/hosts/%s", appConfig.Etcd.Prefix, req.Name), req.IpAddress, &client.SetOptions{}); err != nil {
    server.Logger.Error(err.Error())
    return c.String(http.StatusInternalServerError, "Internal server error")
  }

  return c.JSON(http.StatusCreated, req)
}
