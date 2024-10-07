package main

import (
	"fmt"
	"ipv6-proxy-go/server"
	"net/http"
)

func main() {
	fmt.Println("+ Starting server on port 12333 ...")
	http.HandleFunc("/", server.ForwardRequest)
	http.ListenAndServe(":12333", nil)
}

// go run apps/server/app.go
