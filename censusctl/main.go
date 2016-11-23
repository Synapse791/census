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
  app.Version = "1.0"
  app.Flags = []cli.Flag{
    cli.StringFlag{
      Name: "server-address, s",
      Usage: "Address of the census server",
      Value: "http://localhost:3100",
    },
  }
  app.Commands = []cli.Command{
    cli.Command{
      Name: "list",
      Usage: "Lists resources of given type",
      Subcommands: []cli.Command{
        cli.Command{
          Name: "apps",
          Usage: "Lists apps in the census",
          Action: commands.ListApps,
        },
        cli.Command{
          Name: "hosts",
          Usage: "Lists hosts in the census",
          Action: commands.ListHosts,
        },
      },
    },
    cli.Command{
      Name: "add",
      Usage: "Adds a new resource of the given type",
      Subcommands: []cli.Command{
        cli.Command{
          Name: "app",
          Usage: "Adds a new app to the census",
          Flags: []cli.Flag{
            cli.StringFlag{
              Name: "name, n",
              Usage: "The name for the new app",
            },
            cli.StringFlag{
              Name: "port, p",
              Usage: "The port for the new app",
            },
          },
          Action: commands.AddApp,
        },
        cli.Command{
          Name: "host",
          Usage: "Adds a new host to the census",
          Flags: []cli.Flag{
            cli.StringFlag{
              Name: "name, n",
              Usage: "The name for the new host",
            },
            cli.StringFlag{
              Name: "ip-address, i",
              Usage: "The ip address for the new host",
            },
          },
          Action: commands.AddHost,
        },
      },
    },
  }

  app.Run(os.Args)
}
