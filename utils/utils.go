package utils

import (
  "os"
  "net"
  "fmt"
  "strings"
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

func GetLastPart(s string, delimeter string) string {
  parts := strings.Split(s, delimeter)
  return parts[len(parts) - 1]
}

func ConnectionOk(host string, port string) bool {
  conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
  if err != nil {
    return false
  }
  conn.Close()

  return true
}
