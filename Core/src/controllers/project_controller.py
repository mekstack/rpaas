import re
from collections import defaultdict
from typing import List, Set

from proto.objects_pb2 import Route, Subdomain
from repositories.domain_repository import DomainRepository
from repositories.project_repository import ProjectRepository
from repositories.subdomain_repository import SubdomainRepository


class ProjectController:
    def __init__(self, session):
        self.session = session
    
    @staticmethod
    def is_legit_domain(legit_domains: List[str] | Set[str], subdomain_name: str) -> bool:
        domain_layers = subdomain_name.split(".")
        return True in [
            ".".join(domain_layers[i:]) in legit_domains
            for i in range(len(domain_layers))
            if i != len(domain_layers) - 1
        ]
    
    @staticmethod
    def subdomain_checker(subdomain_name: str):
        return re.match(
            r"^(?:[a-zA-Z1-9]{1,63}\.)+(?:[a-zA-Z1-9]{1,63})+(?:\.[a-zA-Z1-9]{1,63})+$",
            subdomain_name
        )
    
    @staticmethod
    def target_ip_checker(target_ip: str):
        return re.match(
            r"(?:25[0-5]|2[0-4]\d|[0-1]?\d{1,2})(?:\.(?:25[0-5]|2[0-4]\d|[0-1]?\d{1,2})){3}$",
            target_ip
        )
    
    @staticmethod
    def parse_route(subdomains: Set[str] | List[str], target_ip: str) -> Set[str] | List[str]:
        return [
            "{}:{}".format(subdomain_name, target_ip)
            for subdomain_name in subdomains
        ]
    
    async def get_project_info(self, code: int):
        routes_from_db = await ProjectRepository(self.session).get_by_code(code)
        
        routes = defaultdict(set)
        project_routes = list()
        for route in routes_from_db:
            subdomain, target_ip = route.split(":")
            routes[target_ip].add(subdomain)
        for key, value in routes.items():
            project_routes.append(
                Route(target_ip=key, subdomains=[Subdomain(name=subdomain_name) for subdomain_name in value])
            )
        
        return project_routes
    
    async def add_routes(self, code: int, route: Route):
        subdomains, target_ip = route.subdomains, route.target_ip
        
        subdomains = [
            subdomain.name
            for subdomain in subdomains
        ]
        
        if not self.target_ip_checker(target_ip):
            raise ValueError("Target ip has incorrect format")
        
        legit_domain = await DomainRepository(self.session).get_all()
        for subdomain_name in subdomains:
            if not self.is_legit_domain(legit_domain, subdomain_name):
                raise ValueError("There is not such top-level domain")
            
            if await SubdomainRepository(self.session).is_in_list(subdomain_name):
                raise ValueError("Subdomain already exist")
            
            if not self.subdomain_checker(subdomain_name):
                raise ValueError("Subdomain has incorrect format")
        
        await (
            SubdomainRepository(self.session)
            .add_many(subdomains)
        )
        
        await (
            ProjectRepository(self.session)
            .add_routes(
                code,
                self.parse_route(
                    subdomains,
                    target_ip
                )
            ))
        
        return await self.get_project_info(code)
