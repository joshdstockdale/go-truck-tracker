package main

import (
	"log"
)

var kafkaTopic = "obudata"

func main() {
	var (
		err error
		svc CalcServicer
	)
	svc = NewCalcService()
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
