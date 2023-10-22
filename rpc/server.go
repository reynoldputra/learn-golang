package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct {}

type TimeServer int64

// add a function to struct
func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
  *reply = time.Now().Unix()
  return nil
}

func main () {
  timeserver := new(TimeServer)
  rpc.Register(timeserver)
  rpc.HandleHTTP()
  
  listener, err  := net.Listen("tcp", ":1234")

  if err != nil {
    log.Fatal("Listener error: ", err)
  }

  http.Serve(listener, nil)
}
