# @generated by generate_proto_mypy_stubs.py.  Do not edit!
import sys
from google.protobuf.descriptor import (
    Descriptor as google___protobuf___descriptor___Descriptor,
    FileDescriptor as google___protobuf___descriptor___FileDescriptor,
)

from google.protobuf.internal.containers import (
    RepeatedScalarFieldContainer as google___protobuf___internal___containers___RepeatedScalarFieldContainer,
)

from google.protobuf.message import (
    Message as google___protobuf___message___Message,
)

from typing import (
    Iterable as typing___Iterable,
    Optional as typing___Optional,
    Text as typing___Text,
)

from typing_extensions import (
    Literal as typing_extensions___Literal,
)


builtin___bool = bool
builtin___bytes = bytes
builtin___float = float
builtin___int = int


DESCRIPTOR: google___protobuf___descriptor___FileDescriptor = ...

class SummarizeDocumentRequest(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...

    @property
    def document(self) -> type___Document: ...

    def __init__(self,
        *,
        document : typing___Optional[type___Document] = None,
        ) -> None: ...
    def HasField(self, field_name: typing_extensions___Literal[u"document",b"document"]) -> builtin___bool: ...
    def ClearField(self, field_name: typing_extensions___Literal[u"document",b"document"]) -> None: ...
type___SummarizeDocumentRequest = SummarizeDocumentRequest

class SummarizeDocumentReply(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    content: builtin___bytes = ...

    def __init__(self,
        *,
        content : typing___Optional[builtin___bytes] = None,
        ) -> None: ...
    def ClearField(self, field_name: typing_extensions___Literal[u"content",b"content"]) -> None: ...
type___SummarizeDocumentReply = SummarizeDocumentReply

class Document(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    content: builtin___bytes = ...

    @property
    def source(self) -> type___DocumentSource: ...

    def __init__(self,
        *,
        content : typing___Optional[builtin___bytes] = None,
        source : typing___Optional[type___DocumentSource] = None,
        ) -> None: ...
    def HasField(self, field_name: typing_extensions___Literal[u"source",b"source"]) -> builtin___bool: ...
    def ClearField(self, field_name: typing_extensions___Literal[u"content",b"content",u"source",b"source"]) -> None: ...
type___Document = Document

class DocumentSource(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    http_uri: typing___Text = ...

    def __init__(self,
        *,
        http_uri : typing___Optional[typing___Text] = None,
        ) -> None: ...
    def ClearField(self, field_name: typing_extensions___Literal[u"http_uri",b"http_uri"]) -> None: ...
type___DocumentSource = DocumentSource

class SummarizeDocumentWorkload(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    aggregateColumns: google___protobuf___internal___containers___RepeatedScalarFieldContainer[typing___Text] = ...
    excludedColumns: google___protobuf___internal___containers___RepeatedScalarFieldContainer[typing___Text] = ...

    @property
    def request(self) -> type___SummarizeDocumentRequest: ...

    def __init__(self,
        *,
        request : typing___Optional[type___SummarizeDocumentRequest] = None,
        aggregateColumns : typing___Optional[typing___Iterable[typing___Text]] = None,
        excludedColumns : typing___Optional[typing___Iterable[typing___Text]] = None,
        ) -> None: ...
    def HasField(self, field_name: typing_extensions___Literal[u"request",b"request"]) -> builtin___bool: ...
    def ClearField(self, field_name: typing_extensions___Literal[u"aggregateColumns",b"aggregateColumns",u"excludedColumns",b"excludedColumns",u"request",b"request"]) -> None: ...
type___SummarizeDocumentWorkload = SummarizeDocumentWorkload
