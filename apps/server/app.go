package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"ipv6-proxy-go/server"
	"net/http"
	"strconv"
)

func parseArgs() *int {
	port := pflag.IntP("port", "p", 12333, "Port to run the server on")
	pflag.Parse()
	return port
}

func main() {
	port := parseArgs()
	fmt.Printf("+ Starting ForwardRequest server on port: [%d]\n", *port)
	http.HandleFunc("/", server.ForwardRequest)
	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}

// go run apps/server/app.go
