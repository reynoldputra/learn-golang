package golang_website

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandle(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
	  fmt.Fprint(writer, request.Method, "\n")
		fmt.Fprint(writer, "Hello World")
	}
	server := http.Server{
	  Addr: "localhost:5003",
	  Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
	  panic(err)
	}

}
