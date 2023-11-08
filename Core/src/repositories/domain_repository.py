from typing import Set

from repositories.base_repository import BaseRepository


class DomainRepository(BaseRepository):
    key = "domains"
    
    async def get_all(self) -> Set[str]:
        return await self._session.smembers(self.key)
