package main

import (
	"encoding/json"
	"fmt"
	"ipv6-proxy-go/ipv6"
)

func main() {
	ipv6_info := ipv6.GetIPv6Info(ipv6.Args{Verbose: false})
	pretty_json, _ := json.MarshalIndent(ipv6_info, "", "  ")
	fmt.Println(string(pretty_json))
}

// go run main.go
