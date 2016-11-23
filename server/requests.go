package server

type PutHostRequest struct {
  Name      string  `json:"name"`
  IpAddress string  `json:"ip"`
}

type PutAppRequest struct {
  Name  string  `json:"name"`
  Port  string  `json:"port"`
}
