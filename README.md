# ipv6-proxy-go
A service in go to proxy network requests with ipv6 addrs

## Setup

Init:

```sh
go mod init ipv6-proxy-go
```

Setup:
 
```sh
go mod tidy
```

## Install dependencies

```sh
go mod download
```

## Test ipv6

```sh
go run tests/test_ipv6/test.go
```

## Run server

```sh
go run apps/server/app.go
# go run apps/server/app.go -p 12333
```

This would output:

```sh
+ Starting ForwardRequest server on port: [12333]
```

and use curl:

```sh
curl --proxy http://127.0.0.1:12333 http://test.ipw.cn
```

This should output the generated random ipv6 addr.