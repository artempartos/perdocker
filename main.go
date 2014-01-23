package main

import (
	"./perd"
	"flag"
)

var httpPort string

func main() {
	var rubyWorkers, nodejsWorkers int
	var timeout int64

	flag.StringVar(&httpPort, "port", "8080", "HTTP server port.")
	flag.IntVar(&rubyWorkers, "ruby-workers", 1, "Count of ruby workers.")
	flag.IntVar(&nodejsWorkers, "nodejs-workers", 1, "Count of nodejs workers.")
	flag.Int64Var(&timeout, "timeout", 30, "Max execution time.")
	flag.Parse()

	workers := map[string]int{"ruby": rubyWorkers, "nodejs": nodejsWorkers}
	server := perd.NewServer(httpPort, workers, timeout)

	server.Run()
}
