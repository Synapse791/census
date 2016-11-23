package commands

import (
  "fmt"
  "net/http"
  "time"
  "encoding/json"
  "io/ioutil"
  "github.com/urfave/cli"
)

func ListApps(c *cli.Context) error {
  netClient := &http.Client{
    Timeout: time.Second * 10,
  }

  response, resErr := netClient.Get(fmt.Sprintf("%s/apps", c.GlobalString("server-address")))
  if resErr != nil {
    return cli.NewExitError(resErr.Error(), 1)
  }
  defer response.Body.Close()

  body, readErr := ioutil.ReadAll(response.Body)
  if readErr != nil {
    return cli.NewExitError(readErr.Error(), 1)
  }

  if response.StatusCode != 200 {
    return cli.NewExitError(fmt.Sprintf("Failed to get app list: %s", string(body)), 1)
  }

  var decodedRepsonse map[string]string

  jsonErr := json.Unmarshal(body, &decodedRepsonse)
  if jsonErr != nil {
    return cli.NewExitError(jsonErr.Error(), 1)
  }

  for appName, appPort := range decodedRepsonse {
    fmt.Printf("  Name: %s, Port: %s\n", appName, appPort)
  }

  return nil
}
