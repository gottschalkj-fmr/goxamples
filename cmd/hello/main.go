/*
 * Copyright (c) 2020 gottschalkj-fmr.
 * Licensed under the Apache License, Version 2.0
 * http://www.apache.org/licenses/LICENSE-2.0
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gottschalkjfmr/goxamples/pkg/hello"
)

var (
	quits   chan bool      = make(chan bool)
	errors  chan error     = make(chan error)
	signals chan os.Signal = make(chan os.Signal)
)

func init() {
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.LUTC)
}

func health(w http.ResponseWriter, r *http.Request) {
	log.Println("Health checked")
	fmt.Fprintln(w, "OK")
}

func greet(w http.ResponseWriter, r *http.Request) {
	log.Println("Greeting served")
	fmt.Fprintln(w, hello.Greeting())
	quits <- true
}

func run() {
	log.Println("Server starting")
	errors <- http.ListenAndServe(":9090", nil)
}

func wait() {
	select {
	case <-quits:
		log.Println("Server exiting")
		time.Sleep(2 * time.Second)
	case err := <-errors:
		log.Println(err)
	case sig := <-signals:
		log.Printf("OS shutdown signal %+v\n", sig)
	}
}

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/greet", greet)
	go run()
	wait()
}
