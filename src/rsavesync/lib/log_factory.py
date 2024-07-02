import logging
import os.path

from rsavesync.lib.constants import PROGRAM_NAME


def create() -> logging.Logger:
    log_dir = os.path.expanduser(f'~/.config/{PROGRAM_NAME}/logs')
    os.makedirs(log_dir, exist_ok=True)

    logger = logging.getLogger(PROGRAM_NAME)
    logger.setLevel(logging.DEBUG)

    file_handler = logging.FileHandler(os.path.join(log_dir, f'{PROGRAM_NAME}.log'))
    console_handler = logging.StreamHandler()

    file_handler.setLevel(logging.DEBUG)
    console_handler.setLevel(logging.DEBUG)

    formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
    file_handler.setFormatter(formatter)
    console_handler.setFormatter(formatter)

    logger.addHandler(file_handler)
    logger.addHandler(console_handler)

    return logger
