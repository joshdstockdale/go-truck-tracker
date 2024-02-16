package main

import (
	"context"
	"log"
	"time"

	"github.com/joshdstockdale/go-truck-tracker/aggregator/client"
	"github.com/joshdstockdale/go-truck-tracker/types"
)

func main() {

	c, err := client.NewGRPCClient(":3001")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := c.Aggregate(context.Background(), &types.AggregateRequest{
		ObuID: 1,
		Value: 22.22,
		Unix:  time.Now().UnixNano(),
	}); err != nil {
		log.Fatal(err)
	}
}
