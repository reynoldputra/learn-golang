
package golang_website

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServerMux(t *testing.T) {
  mux := http.NewServeMux()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello world")
  })

  mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hi")
  })

	server := http.Server{
	  Addr: "localhost:5003",
	  Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
	  panic(err)
	}

}
