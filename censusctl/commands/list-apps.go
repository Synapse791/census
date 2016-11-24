package commands

import (
  "fmt"
  "encoding/json"
  "io/ioutil"
  "github.com/urfave/cli"
)

func ListApps(c *cli.Context) error {
  netClient := newHttpClient()

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

  table := newTableWriter()
  table.SetHeader([]string{"Name", "Port"})

  for appName, appPort := range decodedRepsonse {
    table.Append([]string{appName, appPort})
  }

  table.Render()

  return nil
}
