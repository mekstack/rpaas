from typing import List

from repositories.base_repository import BaseRepository


class ProjectRepository(BaseRepository):
    key = "project"
    
    def _get_key(self, code: int) -> str:
        return f"{self.key}:{code}"
    
    async def get_by_code(self, code: int):
        return await self._session.smembers(self._get_key(code))
    
    async def add_routes(self, code: int, routes: List[str]):
        await self._session.sadd(self._get_key(code), *routes)
        return await self.get_by_code(code)
