from redis.asyncio import Redis

from configs.app_config import AppConfig


class RedisConfig(Redis):
    def __init__(self):
        super().__init__(
            host=AppConfig.REDIS_HOST,
            port=AppConfig.REDIS_PORT,
            decode_responses=True
        )
