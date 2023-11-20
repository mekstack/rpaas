from repositories.subdomain_repository import SubdomainRepository


class SubdomainController:
    def __init__(self, session):
        self.session = session
    
    async def get_all_subdomains(self):
        return await SubdomainRepository(self.session).get_all()
