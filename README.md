# go-boilerplate
go-boilerplate is not a framework.

go-boilerplate integrates go-cloud, wire, gin, gRPC, grpc-gateway to proivde a baseline for new projects. The code can be copied and modified to suit requirement of a particular project.

## How to run the example
There are 3 examples in the codebase:

* HTTP Service, implemented with gin, provides HTTP endpoint
* gRPC Service, implemented with gRPC
* gRPC HTTP Gateway, a gateway that translate existing gRPC services to HTTP endpoints
### HTTP Service
Run the following command to start the HTTP service.
```
$ go run cmd/httpsrv/main.go 
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /v1/ping                  --> github.com/57blocks/go-boilerplate/internal/handler/ping.(*Resource).Ping-fm (1 handlers)
2021/03/24 09:20:36 starting server listening at :8080 ...
```
Test with cURL
```
$ curl http://localhost:8080/v1/ping
{"value":"pong"}
```
### gRPC Service
Run the following command to start the gRPC service.
```
$ go run cmd/grpcsrv/main.go 
2021/03/24 09:24:48 starting server listening at :50051 ...
```
Run the following command to test the service.
```
$ go run cmd/grpcclnt/main.go 
2021/03/24 09:25:40 Echo: hello grpc, from: default echo service
```
### gRPC with HTTP gateway
With the gRPC service started up, run the following command to start the gateway.
```
$ go run cmd/gateway/main.go 
2021/03/24 10:48:16 starting gateway listening at :8081 ...
```
And test with cURL
```
$ curl -X POST http://localhost:8081/v1/echo
{"value":", from: default echo service"}                        
```