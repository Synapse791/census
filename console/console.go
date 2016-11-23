package console

import (
  "fmt"
  "time"
  "os"
)

func Line(message string, args... interface{}) {
  formattedMessage := fmt.Sprintf(message, args...)
  fmt.Printf("[%s]  %s\n", getTimestamp(), formattedMessage)
}

func Read(message string, args... interface{}) string {
  var data string
  formattedMessage := fmt.Sprintf(message, args...)
  fmt.Printf("[\033[36m%s\033[00m]  \033[36m%s: \033[00m", getTimestamp(), formattedMessage)
  fmt.Scan(&data)
  return data
}

func Success(message string, args... interface{}) {
  formattedMessage := fmt.Sprintf(message, args...)
  fmt.Printf("[\033[32m%s\033[00m]  [\033[32mSUCC\033[00m]  \033[32m%s\033[00m\n", getTimestamp(), formattedMessage)
}

func Warning(message string, args... interface{}) {
  formattedMessage := fmt.Sprintf(message, args...)
  fmt.Printf("[\033[33m%s\033[00m]  [\033[33mWARN\033[00m]  \033[33m%s\033[00m\n", getTimestamp(), formattedMessage)
}

func Error(message string, args... interface{}) {
  formattedMessage := fmt.Sprintf(message, args...)
  fmt.Printf("[\033[31m%s\033[00m]  [\033[31mERRO\033[00m]  \033[31m%s\033[00m\n", getTimestamp(), formattedMessage)
  os.Exit(1)
}

func getTimestamp() string {
  currentTime := time.Now()
  return fmt.Sprintf("%02d:%02d:%02d", currentTime.Hour(), currentTime.Minute(), currentTime.Second())
}
