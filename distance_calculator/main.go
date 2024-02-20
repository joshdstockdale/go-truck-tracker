package main

import (
	"log"

	"github.com/joshdstockdale/go-truck-tracker/aggregator/client"
)

const (
	kafkaTopic   = "obudata"
	httpEndpoint = "http://127.0.0.1:4000"
	grpcEndpoint = ":4001"
)

func main() {
	var (
		err error
		svc CalcServicer
	)
	svc = NewCalcService()
	svc = NewLogMiddleware(svc)

	httpClient := client.NewHTTPClient(httpEndpoint)
	//grpcClient, err := client.NewGRPCClient(grpcEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, httpClient)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
