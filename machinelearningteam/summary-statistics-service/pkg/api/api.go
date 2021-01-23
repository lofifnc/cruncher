package api

import (
	"context"
	"fmt"
	api "github.com/lofifnc/cruncher/machinelearningteam/summary-statistics-service/proto"
	"google.golang.org/grpc"
	"log"
	"os"
	"strings"
	"time"
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

type serverContext struct {
	// client to GRPC service
	backendClient api.DocumentSummarizerBackendClient

	// default timeout
	timeout time.Duration
}

// constructor for server context
func newServerContext(endpoint string) (*serverContext, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &serverContext{
		backendClient: api.NewDocumentSummarizerBackendClient(conn),
		timeout: time.Second,
	}
	return ctx, nil
}

// Server is a server implementing the proto API
type Server struct {
	api.UnimplementedDocumentSummarizerServer
	context *serverContext
}

func NewServer() *Server {
	serverCtx, err := newServerContext(serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	return &Server{context: serverCtx}
}

func (s *Server) SummarizeDocument(
	ctx context.Context,
	req *api.SummarizeDocumentRequest,
) (*api.SummarizeDocumentReply, error) {
	fmt.Println("Received document...")
	clientCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	reply, err := s.context.backendClient.SummarizeDocument(clientCtx, &api.SummarizeDocumentWorkload{
		Request: req,
		ExcludedColumns: excludedColumns,
		AggregateColumns: aggregateColumns,
	})
	if err != nil {
		log.Printf(".GetFeatures(_) = _, %v: ", err)
		return nil, err
	}
	return reply, err
}
