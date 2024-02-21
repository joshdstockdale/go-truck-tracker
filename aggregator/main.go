package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/joshdstockdale/go-truck-tracker/types"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	var (
		store    = makeStore()
		svc      = NewInvoiceAggregator(store)
		httpAddr = os.Getenv("AGG_HTTP_ENDPOINT")
		grpcAddr = os.Getenv("AGG_GRPC_ENDPOINT")
	)
	svc = NewMetricsMiddleware(svc)
	svc = NewLogMiddleware(svc)
	go func() {
		log.Fatal(makeGRPCTransport(grpcAddr, svc))
	}()
	log.Fatal(makeHttpTransport(httpAddr, svc))

}

func makeHttpTransport(listenAddr string, svc Aggregator) error {
	var (
		aggMetrichandler = newHTTPMetricsHandler("aggregate")
		invMetrichandler = newHTTPMetricsHandler("invoice")
	)
	http.HandleFunc("/aggregate", aggMetrichandler.instrument(handleAggregate(svc)))
	http.HandleFunc("/invoice", invMetrichandler.instrument(handleGetInvoice(svc)))
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("HTTP transport running on ", listenAddr)
	return http.ListenAndServe(listenAddr, nil)
}

func makeGRPCTransport(listenAddr string, svc Aggregator) error {
	fmt.Println("GRPC transport running on ", listenAddr)
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	server := grpc.NewServer([]grpc.ServerOption{}...)
	types.RegisterAggregatorServer(server, NewGRPCAggregatorServer(svc))
	return server.Serve(ln)
}

func makeStore() Storer {
	storeType := os.Getenv("AGG_STORE_TYPE")
	switch storeType {
	case "memory":
		return NewMemoryStore()
	default:
		log.Fatalf("Invalid store type")
	}
	return nil
}

func writeJSON(rw http.ResponseWriter, status int, v any) error {
	rw.WriteHeader(status)
	rw.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(rw).Encode(v)
}
