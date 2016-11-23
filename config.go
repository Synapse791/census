package main

import (
  "github.com/Synapse791/census/utils"
  "github.com/Synapse791/census/console"
  "gopkg.in/yaml.v2"
)

const CONFIG_PATH = "/etc/census/config.yml"

type CensusConfig struct {
  Etcd  struct {
    Host  string  `yaml:"host"`
    Port  int     `yaml:"port"`
  } `yaml:"etcd"`
}

func LoadConfig() *CensusConfig {
  if ! utils.FileExists(CONFIG_PATH) {
    console.Error("Config file not found: %s", CONFIG_PATH)
  }

  rawConfig, err := utils.ReadFile(CONFIG_PATH)
  if err != nil {
    console.Error(err.Error())
  }

  appConfig := CensusConfig{}

  if err := yaml.Unmarshal(rawConfig, &appConfig); err != nil {
    console.Error(err.Error())
  }

  if appConfig.Etcd.Host == "" {
    console.Warning("etcd host is not defined in %s.             Using default: 127.0.0.1", CONFIG_PATH)
    appConfig.Etcd.Host = "127.0.0.1"
  }

  if appConfig.Etcd.Port == 0 {
    console.Warning("etcd port is not defined in %s or set to 0. Using default: 2379", CONFIG_PATH)
    appConfig.Etcd.Port = 2379
  }

  return &appConfig
}
