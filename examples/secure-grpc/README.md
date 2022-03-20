# secure-grpc

This example shows how to connect to an Access Node using the secure gRPC endpoint.

Note: This example uses the `common.FlowClient` library to create the `flow-go-sdk/client`. This 
implementation adds a custom TLS verification module that verifies the server's certificate has
a special signed extension that allows you to verify it was created by the specific node you are
connecting to. (See the [libp2p spec](https://github.com/libp2p/specs/blob/master/tls/tls.md#libp2p-public-key-extension) for more details)

## Usage:
```
go run main.go
```