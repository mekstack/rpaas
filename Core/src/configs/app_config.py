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
    SERVER_HOST = os.environ["SERVER_HOST"]
    SERVER_PORT = int(os.environ["SERVER_PORT"])
    
    # Database
    REDIS_HOST = os.environ["REDIS_HOST"]
    REDIS_PORT = int(os.environ["REDIS_PORT"])
