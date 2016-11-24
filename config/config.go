package config

import (
  "strings"
  "github.com/Synapse791/census/utils"
  "github.com/Synapse791/census/console"
  "gopkg.in/yaml.v2"
)

const CONFIG_PATH = "/etc/census/config.yml"

type CensusConfig struct {
  Server struct {
    Addr  string  `yaml:"address"`
    Port  int     `yaml:"port"`
  } `yaml:"server"`
  Etcd  struct {
    Endpoints []string  `yaml:"endpoints"`
    Prefix    string    `yaml:"prefix"`
  } `yaml:"etcd"`
}

func LoadConfig() *CensusConfig {
  if ! utils.FileExists(CONFIG_PATH) {
    console.ErrorAndExit("Config file not found: %s", CONFIG_PATH)
  }

  rawConfig, err := utils.ReadFile(CONFIG_PATH)
  if err != nil {
    console.ErrorAndExit(err.Error())
  }

  appConfig := CensusConfig{}

  if err := yaml.Unmarshal(rawConfig, &appConfig); err != nil {
    console.ErrorAndExit(err.Error())
  }

  if appConfig.Server.Addr == "" {
    console.Warning("server host is not defined in %s.  Using default: \"127.0.0.1\"", CONFIG_PATH)
    appConfig.Server.Addr = "127.0.0.1"
  }

  if appConfig.Server.Port == 0 {
    console.Warning("server port is not defined in %s.  Using default: 3100", CONFIG_PATH)
    appConfig.Server.Port = 3100
  }

  if len(appConfig.Etcd.Endpoints) < 1 {
    console.Warning("no etcd endpoints defined in %s.   Using default: \"http://localhost:2379\"", CONFIG_PATH)
    appConfig.Etcd.Endpoints = []string{"http://localhost:2379"}
  }

  if appConfig.Etcd.Prefix == "" {
    console.Warning("etcd prefix is not defined in %s.  Using default: \"\"", CONFIG_PATH)
  }

  if ! strings.HasPrefix(appConfig.Etcd.Prefix, "/") {
    appConfig.Etcd.Prefix = "/" + appConfig.Etcd.Prefix
  }

  return &appConfig
}
