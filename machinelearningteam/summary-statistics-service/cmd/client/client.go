package main

import (
	"context"
	"flag"
	pb "github.com/lofifnc/cruncher/machinelearningteam/summary-statistics-service/proto"
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/grpc"
)

const github_source_uri = "https://raw.githubusercontent.com/e-conomic/hiring-assignments/master/machinelearningteam/summary-statistics-service/test.csv"

func sendRequest(client pb.DocumentSummarizerClient, request pb.SummarizeDocumentRequest) (*pb.SummarizeDocumentReply, error) {
	ctx := context.Background()
	resp, err := client.SummarizeDocument(ctx, &request)
	if err != nil {
		log.Fatalf("could not send: %v", err)
	}
    return resp, err
}

func readAndSendRequest(client pb.DocumentSummarizerClient, file string) (*pb.SummarizeDocumentReply, error){
	document, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Couldn't read input document")
	}
	return sendRequest(client, pb.SummarizeDocumentRequest{
		Document: &pb.Document{
			Content: document,
		},
	})
}

func sendSourceRequest(client pb.DocumentSummarizerClient, source_uri string) (*pb.SummarizeDocumentReply, error){
	return sendRequest(client, pb.SummarizeDocumentRequest{
		Document: &pb.Document{
			Source: &pb.DocumentSource{
				HttpUri: source_uri,
			},
		},
	})
}

func main() {
	host := flag.String("host", "localhost:50051", "host of api")
	sourceInput := flag.Bool("source", false, "send document via uri")
	numberOfRequests := flag.Int("requests",1, "number of requests sent")
	flag.Parse()
	log.Println(*host)
	conn, err := grpc.Dial(*host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewDocumentSummarizerClient(conn)

	if *numberOfRequests > 1 {
		for i := 1; i < *numberOfRequests; i++ {
			// just fire against the client ignore errors
			_, _ = sendSourceRequest(client, github_source_uri)
		}
		os.Exit(0)
	}

	var resp *pb.SummarizeDocumentReply
	if *sourceInput {
		resp, _ = sendSourceRequest(client, github_source_uri)
	} else {
		resp, _ = readAndSendRequest(client, "test.csv")
	}
	ioutil.WriteFile("out.csv", resp.GetContent(), 0644)

}
