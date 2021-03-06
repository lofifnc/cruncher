from http.client import HTTPResponse

from service import api_pb2
from service import api_pb2_grpc
from statistics import calc_statistics
from urllib import request
from prometheus_client import Summary

from concurrent import futures
import grpc
import logging
import os

PORT = os.getenv("PORT", "10001")

REQUEST_TIME = Summary("request_processing_seconds", "Time spent processing request")


def download_csv(csv_uri: str) -> bytes:
    response: HTTPResponse = request.urlopen(csv_uri)
    return response.read()


@REQUEST_TIME.time()
def summarize_document(request):
    document = request.request.document
    csv_bytes = (
        document.content if document.content else download_csv(document.source.http_uri)
    )
    statistics = calc_statistics(
        csv_bytes.decode("utf-8"), request.aggregateColumns, request.excludedColumns
    )
    reply = api_pb2.SummarizeDocumentReply()
    reply.content = statistics.encode("utf-8")
    return reply


class DocumentSummarizerBackendServicer(api_pb2_grpc.DocumentSummarizerBackendServicer):
    def SummarizeDocument(self, request: api_pb2.SummarizeDocumentWorkload, context):
        return summarize_document(request)


def serve():
    logging.info(f"Starting server on port: {PORT}")
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    api_pb2_grpc.add_DocumentSummarizerBackendServicer_to_server(
        DocumentSummarizerBackendServicer(), server
    )
    server.add_insecure_port(f"[::]:{PORT}")
    server.start()
    server.wait_for_termination()
