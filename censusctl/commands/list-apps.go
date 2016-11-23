package commands

import (
  "github.com/urfave/cli"
)

func ListApps(c *cli.Context) error {
  println("this will list apps and their ports")
  return nil
}
