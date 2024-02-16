package main

import (
	"log"

	"github.com/joshdstockdale/go-truck-tracker/aggregator/client"
)

const (
	kafkaTopic  = "obudata"
	aggEndpoint = "http://127.0.0.1:3000/aggregate"
)

func main() {
	var (
		err error
		svc CalcServicer
	)
	svc = NewCalcService()
	svc = NewLogMiddleware(svc)
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, client.NewClient(aggEndpoint))
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
