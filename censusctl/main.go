package main

import (
  "os"
  "github.com/urfave/cli"
  "github.com/Synapse791/census/censusctl/commands"
)

func main() {
  app := cli.NewApp()
  app.Name = "censusctl"
  app.Usage = "Census Control"
  app.Commands = []cli.Command{
    {
      Name: "list",
      Usage: "Lists resources of given type",
      Subcommands: []cli.Command{
        {
          Name: "apps",
          Usage: "Lists apps in the census",
          Action: commands.ListApps,
        },
        {
          Name: "hosts",
          Usage: "Lists hosts in the census",
          Action: commands.ListHosts,
        },
      },
    },
    {
      Name: "add",
      Usage: "Adds a new resource of the given type",
      Subcommands: []cli.Command{
        {
          Name: "app",
          Usage: "Adds a new app to the census",
          Action: commands.AddApp,
        },
        {
          Name: "host",
          Usage: "Adds a new host to the census",
          Action: commands.AddHost,
        },
      },
    },
  }

  app.Run(os.Args)
}
