package main

import (
	"fmt"
	"log"
	"net/http"
)

func main () {
  /// Handle without mux
  // var handler http.HandlerFunc = func(writer http.ResponseWriter, req *http.Request){
  //   fmt.Fprint(writer, "Hello world")
  // }

  mux := http.NewServeMux()

  mux.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
    fmt.Fprint(writer, "endpoint /")
  })
  
  mux.HandleFunc("/home-page", func(writer http.ResponseWriter, req *http.Request) {
    log.Print(req)
    writer.WriteHeader(200)
    fmt.Fprint(writer, "endpoint /home-page")
  })
  
  server := http.Server{
    Addr : "localhost:5050",
    Handler: mux,
  }

  err := server.ListenAndServe()

  if err != nil {
    panic(err)
  }

}
