package ipv6

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
)

func CheckIPAddress(ipv6_addr string, verbose bool) {
	fmt.Println("> Check IP addr:")

	url := "https://test.ipw.cn"
	if verbose {
		fmt.Println("  * url   :", url)
		fmt.Println("  * ipv6  :", ipv6_addr)
	}

	transport := &http.Transport{
		DialContext: func(ctx context.Context, _, addr string) (net.Conn, error) {
			localAddr := &net.TCPAddr{
				IP: net.ParseIP(ipv6_addr),
			}
			dialer := &net.Dialer{
				LocalAddr: localAddr,
			}
			return dialer.DialContext(ctx, "tcp", addr)
		},
	}

	client := &http.Client{Transport: transport}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("  × Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("  × Error:", err)
		return
	}
	if verbose {
		fmt.Println("  * Status:", resp.Status)
		fmt.Println("  * Result:", string(body))
	}
}
