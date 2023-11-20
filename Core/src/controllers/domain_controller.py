from repositories.domain_repository import DomainRepository


class DomainController:
    def __init__(self, session):
        self.session = session
    
    async def get_all_domains(self):
        return await DomainRepository(self.session).get_all()
