package commands

import (
  "fmt"
  "net/http"
  "bytes"
  "encoding/json"
  "io/ioutil"
  "github.com/urfave/cli"
  "github.com/Synapse791/census/server"
)

func AddApp(c *cli.Context) error {
  if c.String("name") == "" {
    return cli.NewExitError("The name flag is required", 1)
  }

  if c.String("port") == "" {
    return cli.NewExitError("The port flag is required", 1)
  }

  netClient := newHttpClient()

  payload := server.PutAppRequest{}
  payload.Name = c.String("name")
  payload.Port = c.String("port")

  jsonPayload, jsonErr := json.Marshal(payload)
  if jsonErr != nil {
    return cli.NewExitError(jsonErr.Error(), 1)
  }

  req, reqErr := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/apps", c.GlobalString("server-address")), bytes.NewBuffer(jsonPayload))
  if reqErr != nil {
    return cli.NewExitError(reqErr.Error(), 1)
  }

  req.Header.Set("Content-Type", "application/json")

  response, resErr := netClient.Do(req)
  if resErr != nil {
    return cli.NewExitError(resErr.Error(), 1)
  }
  defer response.Body.Close()

  body, readErr := ioutil.ReadAll(response.Body)
  if readErr != nil {
    return cli.NewExitError(readErr.Error(), 1)
  }

  if response.StatusCode != 201 {
    return cli.NewExitError(fmt.Sprintf("Failed to add app: %s", string(body)), 1)
  }

  fmt.Println("App added successfully")

  return nil
}
