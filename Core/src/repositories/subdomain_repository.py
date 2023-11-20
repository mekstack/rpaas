from typing import List, Set

from repositories.base_repository import BaseRepository


class SubdomainRepository(BaseRepository):
    key = "subdomains"
    
    async def get_all(self) -> Set[str]:
        return await self._session.smembers(self.key)
    
    async def add_many(self, routes: List[str]):
        await self._session.sadd(self.key, *routes)
    
    async def is_in_list(self, subdomain_name: str) -> bool:
        return bool(await self._session.sismember(self.key, subdomain_name))
