import grpc

from configs.redis_config import RedisConfig
from controllers.subdomain_controller import SubdomainController
from proto.v1.subdomain.model_pb2 import Subdomain
from proto.v1.subdomain.service_pb2 import GetOccupiedSubdomainsRequest, GetOccupiedSubdomainsResponse
from proto.v1.subdomain.service_pb2_grpc import SubdomainServiceServicer


class SubdomainService(SubdomainServiceServicer):
    async def GetOccupiedSubdomains(
            self,
            request: GetOccupiedSubdomainsRequest,
            context: grpc.aio.ServicerContext) -> GetOccupiedSubdomainsResponse:
        async with RedisConfig() as session:
            subdomains = [
                Subdomain(name=subdomain_name)
                for subdomain_name in await SubdomainController(session).get_all_subdomains()
            ]
        
        return GetOccupiedSubdomainsResponse(subdomains=subdomains)
