# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import service.api_pb2 as api__pb2


class DocumentSummarizerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SummarizeDocument = channel.unary_unary(
            "/statistics.DocumentSummarizer/SummarizeDocument",
            request_serializer=api__pb2.SummarizeDocumentRequest.SerializeToString,
            response_deserializer=api__pb2.SummarizeDocumentReply.FromString,
        )


class DocumentSummarizerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def SummarizeDocument(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details("Method not implemented!")
        raise NotImplementedError("Method not implemented!")


def add_DocumentSummarizerServicer_to_server(servicer, server):
    rpc_method_handlers = {
        "SummarizeDocument": grpc.unary_unary_rpc_method_handler(
            servicer.SummarizeDocument,
            request_deserializer=api__pb2.SummarizeDocumentRequest.FromString,
            response_serializer=api__pb2.SummarizeDocumentReply.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        "statistics.DocumentSummarizer", rpc_method_handlers
    )
    server.add_generic_rpc_handlers((generic_handler,))


# This class is part of an EXPERIMENTAL API.
class DocumentSummarizer(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def SummarizeDocument(
        request,
        target,
        options=(),
        channel_credentials=None,
        call_credentials=None,
        insecure=False,
        compression=None,
        wait_for_ready=None,
        timeout=None,
        metadata=None,
    ):
        return grpc.experimental.unary_unary(
            request,
            target,
            "/statistics.DocumentSummarizer/SummarizeDocument",
            api__pb2.SummarizeDocumentRequest.SerializeToString,
            api__pb2.SummarizeDocumentReply.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
        )


class DocumentSummarizerBackendStub(object):
    """Froblems defining this in separate file has to life here for the moment."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SummarizeDocument = channel.unary_unary(
            "/statistics.DocumentSummarizerBackend/SummarizeDocument",
            request_serializer=api__pb2.SummarizeDocumentWorkload.SerializeToString,
            response_deserializer=api__pb2.SummarizeDocumentReply.FromString,
        )


class DocumentSummarizerBackendServicer(object):
    """Froblems defining this in separate file has to life here for the moment."""

    def SummarizeDocument(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details("Method not implemented!")
        raise NotImplementedError("Method not implemented!")


def add_DocumentSummarizerBackendServicer_to_server(servicer, server):
    rpc_method_handlers = {
        "SummarizeDocument": grpc.unary_unary_rpc_method_handler(
            servicer.SummarizeDocument,
            request_deserializer=api__pb2.SummarizeDocumentWorkload.FromString,
            response_serializer=api__pb2.SummarizeDocumentReply.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        "statistics.DocumentSummarizerBackend", rpc_method_handlers
    )
    server.add_generic_rpc_handlers((generic_handler,))


# This class is part of an EXPERIMENTAL API.
class DocumentSummarizerBackend(object):
    """Froblems defining this in separate file has to life here for the moment."""

    @staticmethod
    def SummarizeDocument(
        request,
        target,
        options=(),
        channel_credentials=None,
        call_credentials=None,
        insecure=False,
        compression=None,
        wait_for_ready=None,
        timeout=None,
        metadata=None,
    ):
        return grpc.experimental.unary_unary(
            request,
            target,
            "/statistics.DocumentSummarizerBackend/SummarizeDocument",
            api__pb2.SummarizeDocumentWorkload.SerializeToString,
            api__pb2.SummarizeDocumentReply.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
        )
