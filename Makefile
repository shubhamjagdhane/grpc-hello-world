server:
	go build -o server ./server && mv server/server ./bin/server && ./bin/server

client:
	go build -o client ./client && mv client/client ./bin/client && ./bin/client

protogen: 
	protoc -Iproto --go_out=. --go_opt=module=github.com/shubhamjagdhane/grpc-hello-world --go-grpc_out=. --go-grpc_opt=module=github.com/shubhamjagdhane/grpc-hello-world proto/hello_world.proto

.PHONY: protogen server client
