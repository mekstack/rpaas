import os
from enum import Enum


class EnvEnum(Enum):
    DEV = "dev"
    PROD = "prod"


class AppConfig:
    # App
    APP_NAME = "Core"
    ENV = EnvEnum.DEV
    
    # Server
    SERVER_HOST = os.getenv("SERVER_HOST")
    SERVER_PORT = int(os.getenv("SERVER_PORT"))
    
    # Database
    REDIS_HOST = os.getenv("REDIS_HOST")
    REDIS_PORT = int(os.getenv("REDIS_PORT"))
