from grpc import aio


class GrpcServer:
    def __init__(self, host: str, port: int):
        self._server = aio.server()
        self._server.add_insecure_port(f'{host}:{port}')
    
    async def run(self):
        await self._server.start()
        await self._server.wait_for_termination()
    
    def add_service(self, adder_function, service_call):
        adder_function(service_call(), self._server)
