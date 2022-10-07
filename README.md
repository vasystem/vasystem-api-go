# vasystem-api-go

This is a Go client library for the [VASystem API](https://docs.vasystem.dev).

You will need to have a client ID and secret, please see the [VASystem API documentation](https://docs.vasystem.dev)
for more information.

For usage information, please see the example at [`cmd/vasystem-api-example-go/main.go`](cmd/vasystem-api-example-go/main.go).

## Installation

```bash
go get github.com/vasystem/vasystem-api-go
```

## Running the example

```shell
go build ./cmd/vaystem-api-example-go
./vasystem-api-example-go -domain "staralliancevirtual.org" -client-id "client-id" -client-secret "client-secret"
```

## Regenerating protos

Install [buf](https://buf.build) and run:

```shell
buf generate
```
