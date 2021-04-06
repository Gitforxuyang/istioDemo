
proto:
	protoc --proto_path=proto --go_out=plugins=grpc:./proto proto/hello.proto

v1:
	GOOS=linux GOARCH=amd64 go build -o 'bin/server' -mod=vendor grpc-v1/main.go
	docker build -t grpc:v1 .

v2:
	GOOS=linux GOARCH=amd64 go build -o 'bin/server' -mod=vendor grpc-v2/main.go
	docker build -t grpc:v2 .

web:
	GOOS=linux GOARCH=amd64 go build -o 'bin/server' -mod=vendor web/main.go
	docker build -t web:v1 .

.PHONY: proto
.PHONY: web