# go-truck-tracker
Microservices for tracking distances with GPS coordinates.

## Kafka container
```
docker compose up -d
```

## Install protobuf compiler
For linux users
```
sudo apt install -y protobuf-compiler
```

## Docs for GRPC and Protobuffer plugins for golang
1. Protobuffers
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```

2. GRPC
```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

3. Set go/bin in PATH
```
PATH="${PATH}:${HOME}/go/bin"
```

4. Install package dependencies
```
go get google.golang.org/protobuf
go get google.golang.org/grpc
```
