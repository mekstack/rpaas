from redis.asyncio import Redis


class BaseRepository:
    def __init__(self, session: Redis):
        self._session: Redis = session
