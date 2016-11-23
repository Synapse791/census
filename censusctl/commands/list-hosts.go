package commands

import (
  "github.com/urfave/cli"
)

func ListHosts(c *cli.Context) error {
  println("this will list hosts")
  return nil
}
