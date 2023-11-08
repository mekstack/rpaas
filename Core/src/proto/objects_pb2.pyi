from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Route(_message.Message):
    __slots__ = ["target_ip", "subdomains"]
    TARGET_IP_FIELD_NUMBER: _ClassVar[int]
    SUBDOMAINS_FIELD_NUMBER: _ClassVar[int]
    target_ip: str
    subdomains: _containers.RepeatedCompositeFieldContainer[Subdomain]
    def __init__(self, target_ip: _Optional[str] = ..., subdomains: _Optional[_Iterable[_Union[Subdomain, _Mapping]]] = ...) -> None: ...

class Domain(_message.Message):
    __slots__ = ["name"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    def __init__(self, name: _Optional[str] = ...) -> None: ...

class GetDomainsPoolRequest(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class GetDomainsPoolResponse(_message.Message):
    __slots__ = ["domains"]
    DOMAINS_FIELD_NUMBER: _ClassVar[int]
    domains: _containers.RepeatedCompositeFieldContainer[Domain]
    def __init__(self, domains: _Optional[_Iterable[_Union[Domain, _Mapping]]] = ...) -> None: ...

class Subdomain(_message.Message):
    __slots__ = ["name"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    def __init__(self, name: _Optional[str] = ...) -> None: ...

class GetOccupiedSubdomainsRequest(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class GetOccupiedSubdomainsResponse(_message.Message):
    __slots__ = ["subdomains"]
    SUBDOMAINS_FIELD_NUMBER: _ClassVar[int]
    subdomains: _containers.RepeatedCompositeFieldContainer[Subdomain]
    def __init__(self, subdomains: _Optional[_Iterable[_Union[Subdomain, _Mapping]]] = ...) -> None: ...

class Project(_message.Message):
    __slots__ = ["code", "routes"]
    CODE_FIELD_NUMBER: _ClassVar[int]
    ROUTES_FIELD_NUMBER: _ClassVar[int]
    code: int
    routes: _containers.RepeatedCompositeFieldContainer[Route]
    def __init__(self, code: _Optional[int] = ..., routes: _Optional[_Iterable[_Union[Route, _Mapping]]] = ...) -> None: ...

class GetProjectInfoRequest(_message.Message):
    __slots__ = ["code"]
    CODE_FIELD_NUMBER: _ClassVar[int]
    code: int
    def __init__(self, code: _Optional[int] = ...) -> None: ...

class GetProjectInfoResponse(_message.Message):
    __slots__ = ["project"]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    project: Project
    def __init__(self, project: _Optional[_Union[Project, _Mapping]] = ...) -> None: ...

class AddProjectInfoRequest(_message.Message):
    __slots__ = ["code", "route"]
    CODE_FIELD_NUMBER: _ClassVar[int]
    ROUTE_FIELD_NUMBER: _ClassVar[int]
    code: int
    route: Route
    def __init__(self, code: _Optional[int] = ..., route: _Optional[_Union[Route, _Mapping]] = ...) -> None: ...

class AddProjectInfoResponse(_message.Message):
    __slots__ = ["project"]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    project: Project
    def __init__(self, project: _Optional[_Union[Project, _Mapping]] = ...) -> None: ...
