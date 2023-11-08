import asyncio

from configs.app_config import AppConfig
from configs.logger_config import app_logger
from grpc_server.server import GrpcServer
from proto.objects_pb2_grpc import (
    add_DomainServiceServicer_to_server,
    add_ProjectServiceServicer_to_server,
    add_SubdomainServiceServicer_to_server
)
from services.domain_service import DomainService
from services.project_service import ProjectService
from services.subdomain_service import SubdomainService

if __name__ == '__main__':
    server = GrpcServer(AppConfig.SERVER_HOST, AppConfig.SERVER_PORT)
    server.add_service(add_DomainServiceServicer_to_server, DomainService)
    server.add_service(add_SubdomainServiceServicer_to_server, SubdomainService)
    server.add_service(add_ProjectServiceServicer_to_server, ProjectService)
    
    app_logger.debug(f"Server start on {AppConfig.SERVER_HOST}:{AppConfig.SERVER_PORT}")
    asyncio.get_event_loop().run_until_complete(server.run())
