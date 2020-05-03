/*
 * Copyright (c) 2020 krautbax.
 * Licensed under the Apache License, Version 2.0
 * http://www.apache.org/licenses/LICENSE-2.0
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/krautbax/goxamples/pkg/hello"
)

var (
	quits   chan bool      = make(chan bool)
	errors  chan error     = make(chan error)
	signals chan os.Signal = make(chan os.Signal)
)

var (
	writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	logger *log.Logger   = log.New(writer, "", log.LstdFlags|log.Lmicroseconds|log.LUTC)
)

func init() {
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, hello.Greeting())
	quits <- true
}

func run() {
	logger.Println("Server starting")
	writer.Flush()
	errors <- http.ListenAndServe(":9090", nil)
}

func wait() {
	defer time.Sleep(2 * time.Second)
	defer writer.Flush()

	select {
	case <-quits:
		logger.Println("Server exiting")
	case err := <-errors:
		logger.Println(err)
	case sig := <-signals:
		logger.Printf("OS shutdown signal %+v\n", sig)
	}
}

func main() {
	http.HandleFunc("/", greet)
	go run()
	wait()
}
