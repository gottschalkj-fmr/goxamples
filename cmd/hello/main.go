/*
 * Copyright (c) 2020 krautbax.
 * Licensed under the Apache License, Version 2.0
 * http://www.apache.org/licenses/LICENSE-2.0
 */

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
