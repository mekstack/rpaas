import grpc

from configs.redis_config import RedisConfig
from controllers.domain_controller import DomainController
from proto.v1.domain.model_pb2 import Domain
from proto.v1.domain.service_pb2 import GetDomainsPoolRequest, GetDomainsPoolResponse
from proto.v1.domain.service_pb2_grpc import DomainServiceServicer


class DomainService(DomainServiceServicer):
    async def GetDomainsPool(
            self,
            request: GetDomainsPoolRequest,
            context: grpc.ServicerContext) -> GetDomainsPoolResponse:
        async with RedisConfig() as session:
            domains = [
                Domain(name=domain_name)
                for domain_name in await DomainController(session).get_all_domains()
            ]
        return GetDomainsPoolResponse(domains=domains)
