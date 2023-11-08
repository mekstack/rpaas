from proto.v1.project import model_pb2 as _model_pb2
from proto.v1.route import model_pb2 as _model_pb2_1
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor


class GetProjectInfoRequest(_message.Message):
    __slots__ = ["code"]
    CODE_FIELD_NUMBER: _ClassVar[int]
    code: int
    
    def __init__(self, code: _Optional[int] = ...) -> None: ...


class GetProjectInfoResponse(_message.Message):
    __slots__ = ["project"]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    project: _model_pb2.Project
    
    def __init__(self, project: _Optional[_Union[_model_pb2.Project, _Mapping]] = ...) -> None: ...


class AddProjectInfoRequest(_message.Message):
    __slots__ = ["code", "route"]
    CODE_FIELD_NUMBER: _ClassVar[int]
    ROUTE_FIELD_NUMBER: _ClassVar[int]
    code: int
    route: _model_pb2_1.Route
    
    def __init__(self, code: _Optional[int] = ...,
                 route: _Optional[_Union[_model_pb2_1.Route, _Mapping]] = ...) -> None: ...


class AddProjectInfoResponse(_message.Message):
    __slots__ = ["project"]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    project: _model_pb2.Project
    
    def __init__(self, project: _Optional[_Union[_model_pb2.Project, _Mapping]] = ...) -> None: ...
