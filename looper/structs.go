package looper

type LooperApp struct {
  Name  string
  Port  string
  Hosts []LooperHost
}

type LooperHost struct {
  Name      string
  IpAddress string
}
