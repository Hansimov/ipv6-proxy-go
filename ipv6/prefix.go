package ipv6

import (
	"fmt"
	"net"
	"strings"
)

type IPv6Info struct {
	Netint   string
	Addr     string
	Prefix   string
	MaskBits int
}

func GetIPv6Info(verbose bool) IPv6Info {
	var ipv6_infos []IPv6Info

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("× Error:", err)
		return IPv6Info{}
	}

	fmt.Println("> Get ipv6 addrs:")

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("× Error:", err)
			continue
		}
		for _, addr := range addrs {
			ip_net, ok := addr.(*net.IPNet)
			if ok && ip_net.IP.To4() == nil && ip_net.IP.To16() != nil {
				ip_addr := ip_net.IP.String()
				mask_bits, _ := ip_net.Mask.Size()
				prefix := ip_net.IP.Mask(ip_net.Mask).String()
				prefix = strings.TrimRight(prefix, ":")
				netint := iface.Name
				if strings.HasPrefix(ip_addr, "2") {
					if verbose {
						fmt.Printf("  * %s: %s\n", netint, ip_addr)
						fmt.Printf("    %s: %s [/%d]\n", netint, prefix, mask_bits)
					}
					ipv6_info := IPv6Info{
						Netint:   netint,
						Addr:     ip_addr,
						Prefix:   prefix,
						MaskBits: mask_bits,
					}
					ipv6_infos = append(ipv6_infos, ipv6_info)
				}
			}
		}
	}
	if len(ipv6_infos) > 0 {
		return ipv6_infos[0]
	}
	return IPv6Info{}
}
