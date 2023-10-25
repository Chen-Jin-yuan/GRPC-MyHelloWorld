module myHelloWorld

go 1.21.3

require (
	google.golang.org/grpc v1.59.0
	google.golang.org/grpc/examples/helloworld/helloworld v0.0.0-20231019174947-e88e8498c6df
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.16.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace google.golang.org/grpc/examples/helloworld/helloworld => ./helloworld