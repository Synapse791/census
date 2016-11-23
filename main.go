package main

import (
  "github.com/Synapse791/census/config"
  "github.com/Synapse791/census/server"
)

func main() {
  appConfig := config.LoadConfig()

  // read app list from etcd
  // if no apps (PREFIX/census/apps/*)
    // nothing to do
  // else
    // read hosts from etcd (PREFIX/census/hosts/*)
    // if no hosts
      // nothing to do
    // else
      // cycle through apps and check port on each host
      // if check succeeds
        // if host not in apps host list
          // add host to apps host list
        // else
          // nothing to do
      // else
        // remove host from apps host list

  server.Run(appConfig)
}
