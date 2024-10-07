# ipv6-proxy-go
A service in go to proxy network requests with ipv6 addrs

## Init

```sh
go mod init ipv6-proxy-go
```

## Test ipv6

```sh
go run tests/test_ipv6/test.go
```

## Run server

```sh
go run apps/server/app.go
```

which would output:

```sh
+ Starting server on port 12333...
```

and use curl:

```sh
curl --proxy http://127.0.0.1:12333 http://test.ipw.cn
```

which would output a random ipv6 addr.