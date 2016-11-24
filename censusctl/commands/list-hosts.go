package commands

import (
  "fmt"
  "encoding/json"
  "io/ioutil"
  "github.com/urfave/cli"
)

func ListHosts(c *cli.Context) error {
  netClient := newHttpClient()

  response, resErr := netClient.Get(fmt.Sprintf("%s/hosts", c.GlobalString("server-address")))
  if resErr != nil {
    return cli.NewExitError(resErr.Error(), 1)
  }
  defer response.Body.Close()

  body, readErr := ioutil.ReadAll(response.Body)
  if readErr != nil {
    return cli.NewExitError(readErr.Error(), 1)
  }

  if response.StatusCode != 200 {
    return cli.NewExitError(fmt.Sprintf("Failed to get host list: %s", string(body)), 1)
  }

  var decodedRepsonse map[string]string

  jsonErr := json.Unmarshal(body, &decodedRepsonse)
  if jsonErr != nil {
    return cli.NewExitError(jsonErr.Error(), 1)
  }

  table := newTableWriter()
  table.SetHeader([]string{"Name", "IP Address"})

  for hostName, hostIp := range decodedRepsonse {
    table.Append([]string{hostName, hostIp})
  }

  table.Render()

  return nil
}
