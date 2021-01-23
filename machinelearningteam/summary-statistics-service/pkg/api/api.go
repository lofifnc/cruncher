package api

import (
	"context"
	"fmt"
	api "github.com/lofifnc/cruncher/machinelearningteam/summary-statistics-service/proto"
	"google.golang.org/grpc"
	"log"
	"os"
	"strings"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// get configuration from environment
var serverAddr = getEnv("SERVER_ADDR", "localhost:10001")
var aggregateColumns = strings.Split(getEnv("AGGREGATE_COLUMNS", "AccountNumber"),",")
var excludedColumns = strings.Split(getEnv("EXCLUDED_COLUMNS", "AccountTypeName"), ",")

// Server is a server implementing the proto API
type Server struct {
	api.UnimplementedDocumentSummarizerServer
}

func (s *Server) SummarizeDocument(
	ctx context.Context,
	req *api.SummarizeDocumentRequest,
) (*api.SummarizeDocumentReply, error) {

	fmt.Println("Received document...")
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Printf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := api.NewDocumentSummarizerBackendClient(conn)
	reply, err := client.SummarizeDocument(ctx, &api.SummarizeDocumentWorkload{
		Request: req,
		ExcludedColumns: excludedColumns,
		AggregateColumns: aggregateColumns,
	})
	if err != nil {
		log.Printf("%v.GetFeatures(_) = _, %v: ", client, err)
	}
	return reply, err
}
