# go-truck-tracker
Microservices for tracking distances with GPS coordinates.

![alt text](https://github.com/joshdstockdale/go-truck-tracker/blob/main/trackerWhite.png?raw=true)

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

## Prometheus
```
docker run --name prometheus -d -v ./.config/prometheus.yml:/etc/prometheus/prometheus.yml -p 9090:9090 --add-host=host.docker.internal:host-gateway prom/prometheus
```

### Golang client
```
go get github.com/prometheus/client_golang/prometheus
```

## Grafana
```
docker run -d -p 3000:3000 --add-host=host.docker.internal:host-gateway  --name=grafana grafana/grafana-enterprise
```


# Running
Once all the docker containers above are running: 
```
make receiver
make obu
make dcalc
make aggregator
make gateway
```

reciever -> localhost:30000/ws

aggregator -> localhost:4000 (http), localhost:4001 (grpc)

gateway -> localhost:6000
