import logging

from configs.app_config import AppConfig, EnvEnum

log_formatter = logging.Formatter("%(asctime)s [%(threadName)-s] [%(levelname)-5.5s]  %(message)s")
app_logger = logging.getLogger()
app_logger.setLevel(
    level=logging.DEBUG if AppConfig.ENV == EnvEnum.DEV else logging.INFO,
)
console_handler = logging.StreamHandler()
console_handler.setFormatter(log_formatter)
app_logger.addHandler(console_handler)
