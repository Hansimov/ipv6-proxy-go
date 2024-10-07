package main

import (
	"encoding/json"
	"fmt"
	"ipv6-proxy-go/ipv6"
)

func main() {
	ipv6_info := ipv6.GetIPv6Info(false)
	pretty_json, _ := json.MarshalIndent(ipv6_info, "", "  ")
	fmt.Println(string(pretty_json))
	random_ipv6 := ipv6.GenerateRandomIPv6(ipv6_info.Prefix, ipv6_info.MaskBits, false)
	fmt.Println(random_ipv6)
}

// go run main.go
