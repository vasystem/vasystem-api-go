module github.com/vasystem/vasystem-api-go

go 1.24.0

require (
	buf.build/gen/go/vasystem/api/grpc/go v1.3.0-20221008143436-bf31ebf1b0ce.2
	buf.build/gen/go/vasystem/api/protocolbuffers/go v1.31.0-20221008143436-bf31ebf1b0ce.2
	golang.org/x/oauth2 v0.34.0
	google.golang.org/grpc v1.79.3
	google.golang.org/protobuf v1.36.10
)

require (
	cloud.google.com/go/compute/metadata v0.9.0 // indirect
	golang.org/x/net v0.48.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251202230838-ff82c1b0f217 // indirect
)
