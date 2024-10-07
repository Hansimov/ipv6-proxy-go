package server

import (
	"context"
	"fmt"
	"io"
	"ipv6-proxy-go/ipv6"
	"net"
	"net/http"
)

func createHTTPClient(random_ipv6 string) *http.Client {
	transport := &http.Transport{
		DialContext: func(ctx context.Context, _, addr string) (net.Conn, error) {
			local_iddr := &net.TCPAddr{
				IP: net.ParseIP(random_ipv6),
			}
			dialer := &net.Dialer{
				LocalAddr: local_iddr,
			}
			return dialer.DialContext(ctx, "tcp", addr)
		},
	}
	return &http.Client{Transport: transport}
}

func copyRequestHeaders(r *http.Request, req *http.Request) {
	for name, values := range r.Header {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}
}

func createRequest(r *http.Request, target_url string) *http.Request {
	req, err := http.NewRequest(r.Method, target_url, r.Body)
	if err != nil {
		fmt.Println("× Error creating request:", err)
		return nil
	}
	copyRequestHeaders(r, req)
	return req
}

func sendRequest(req *http.Request, client *http.Client) *http.Response {
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("× Error sending request:", err)
		return nil
	}
	return resp
}

func copyResponseHeaders(writer http.ResponseWriter, resp *http.Response) {
	for name, values := range resp.Header {
		for _, value := range values {
			writer.Header().Add(name, value)
		}
	}
}

func copyResponse(writer http.ResponseWriter, resp *http.Response) {
	copyResponseHeaders(writer, resp)
	writer.WriteHeader(resp.StatusCode)
	io.Copy(writer, resp.Body)
}

func ForwardRequest(writer http.ResponseWriter, r *http.Request) {
	ipv6_info := ipv6.GetIPv6Info(false)
	random_ipv6 := ipv6.GenerateRandomIPv6(ipv6_info.Prefix, ipv6_info.MaskBits, true)
	client := createHTTPClient(random_ipv6)
	req := createRequest(r, r.URL.String())
	resp := sendRequest(req, client)
	defer resp.Body.Close()
	copyResponse(writer, resp)
}
