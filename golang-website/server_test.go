package golang_website

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
  server := http.Server{
    Addr: "localhost:5003",
  }

  err := server.ListenAndServe()

  if err != nil {
    panic(err)
  }

}
