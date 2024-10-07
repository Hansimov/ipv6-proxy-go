package main

import (
	"context"
	"fmt"
	"io"
	"ipv6-proxy-go/ipv6"
	"net"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Starting server on port 12333...")
	err := http.ListenAndServe(":12333", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	ipv6Info := ipv6.GetIPv6Info(false)
	randomIPv6 := ipv6.GenerateRandomIPv6(ipv6Info.Prefix, ipv6Info.MaskBits, true)

	transport := &http.Transport{
		DialContext: func(ctx context.Context, _, addr string) (net.Conn, error) {
			localAddr := &net.TCPAddr{
				IP: net.ParseIP(randomIPv6),
			}
			dialer := &net.Dialer{
				LocalAddr: localAddr,
			}
			return dialer.DialContext(ctx, "tcp", addr)
		},
	}

	client := &http.Client{Transport: transport}

	targetURL := r.URL.String()
	if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
		targetURL = "http://" + targetURL
	}

	req, err := http.NewRequest(r.Method, r.URL.String(), r.Body)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	// Copy headers from the original request
	for name, values := range r.Header {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error forwarding request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy headers from the response
	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
