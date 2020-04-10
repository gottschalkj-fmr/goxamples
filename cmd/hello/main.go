package main

import (
	"fmt"
	"net/http"

	"github.com/krautbax/goxamples/pkg/hello"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, hello.Greeting())
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":9090", nil)
}
