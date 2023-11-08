from proto.v1.subdomain import model_pb2 as _model_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, \
    Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor


class GetOccupiedSubdomainsRequest(_message.Message):
    __slots__ = []
    
    def __init__(self) -> None: ...


class GetOccupiedSubdomainsResponse(_message.Message):
    __slots__ = ["subdomains"]
    SUBDOMAINS_FIELD_NUMBER: _ClassVar[int]
    subdomains: _containers.RepeatedCompositeFieldContainer[_model_pb2.Subdomain]
    
    def __init__(self, subdomains: _Optional[_Iterable[_Union[_model_pb2.Subdomain, _Mapping]]] = ...) -> None: ...
