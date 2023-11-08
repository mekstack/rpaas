from proto.v1.route import model_pb2 as _model_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, \
    Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor


class Project(_message.Message):
    __slots__ = ["code", "routes"]
    CODE_FIELD_NUMBER: _ClassVar[int]
    ROUTES_FIELD_NUMBER: _ClassVar[int]
    code: int
    routes: _containers.RepeatedCompositeFieldContainer[_model_pb2.Route]
    
    def __init__(self, code: _Optional[int] = ...,
                 routes: _Optional[_Iterable[_Union[_model_pb2.Route, _Mapping]]] = ...) -> None: ...
