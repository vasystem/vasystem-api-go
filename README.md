# vasystem-api-example-go

This is an example of how to use the [VASystem API](https://docs.vasystem.dev) in Go.

You will need to have a client ID and secret, please see the [VASystem API documentation](https://docs.vasystem.dev)
for more information.

## Running the example

```shell
go build -o vasystem-api-example-go
./vasystem-api-example-go -domain "staralliancevirtual.org" -client-id "client-id" -client-secret "client-secret"
```

## Regenerating protos

Install [buf](https://buf.build) and run:

```shell
buf generate
```
