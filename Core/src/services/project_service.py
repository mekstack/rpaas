from grpc import aio, StatusCode

from configs.logger_config import app_logger
from configs.redis_config import RedisConfig
from controllers.project_controller import ProjectController
from proto.objects_pb2 import (
    Project,
    GetProjectInfoRequest,
    GetProjectInfoResponse,
    AddProjectInfoRequest,
    AddProjectInfoResponse,
    Route
)
from proto.objects_pb2_grpc import ProjectServiceServicer


class ProjectService(ProjectServiceServicer):
    async def GetProjectInfo(
            self,
            request: GetProjectInfoRequest,
            context: aio.ServicerContext) -> GetProjectInfoResponse:
        project_code = request.code
        
        async with RedisConfig() as session:
            project_routes = await ProjectController(session).get_project_info(project_code)
        
        return GetProjectInfoResponse(
            project=Project(code=project_code, routes=[
                Route(target_ip=route.target_ip, port=route.port, subdomains=route.subdomains)
                for route in project_routes
            ])
        )
    
    async def AddProjectInfo(
            self,
            request: AddProjectInfoRequest,
            context: aio.ServicerContext) -> AddProjectInfoResponse:
        project_code = request.code
        
        async with RedisConfig() as session:
            try:
                project_routes = await ProjectController(session).add_routes(project_code, request.route)
            except ValueError as error:
                app_logger.error(error)
                if isinstance(error, ValueError):
                    context.set_code(StatusCode.INVALID_ARGUMENT)
                    context.set_details(str(error))
                    return AddProjectInfoResponse()
        
        return AddProjectInfoResponse(
            project=Project(code=project_code, routes=[
                Route(target_ip=route.target_ip, port=route.port, subdomains=route.subdomains)
                for route in project_routes
            ])
        )
