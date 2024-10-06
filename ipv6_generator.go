package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("List ipv6 addrs:")

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		// fmt.Println(iface)
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && ipNet.IP.To4() == nil && ipNet.IP.To16() != nil {
				ipStr := ipNet.IP.String()
				mask_ones, _ := ipNet.Mask.Size()
				prefix := ipNet.IP.Mask(ipNet.Mask).String()
				prefix = strings.TrimRight(prefix, ":")
				if strings.HasPrefix(ipStr, "2") {
					fmt.Printf("  * %s: %s\n", iface.Name, ipStr)
					fmt.Printf("    %s: %s [/%d]\n", iface.Name, prefix, mask_ones)
					break
				}
			}
		}
	}
}
