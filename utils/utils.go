package utils

import (
  "os"
  "io/ioutil"
)

func FileExists(path string) bool {
  if _, err := os.Stat(path); os.IsNotExist(err) {
    return false
  }
  return true
}

func ReadFile(path string) ([]byte, error) {
  return ioutil.ReadFile(path)
}
