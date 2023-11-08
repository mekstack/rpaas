from proto.v1.domain import model_pb2 as _model_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, \
    Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor


class GetDomainsPoolRequest(_message.Message):
    __slots__ = []
    
    def __init__(self) -> None: ...


class GetDomainsPoolResponse(_message.Message):
    __slots__ = ["domains"]
    DOMAINS_FIELD_NUMBER: _ClassVar[int]
    domains: _containers.RepeatedCompositeFieldContainer[_model_pb2.Domain]
    
    def __init__(self, domains: _Optional[_Iterable[_Union[_model_pb2.Domain, _Mapping]]] = ...) -> None: ...
