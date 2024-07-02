#!/usr/bin/env python3

import importlib.metadata
import os
import sys
import argparse

from rsavesync.lib import log_factory
from rsavesync.lib.constants import PROGRAM_NAME


def main():
    if len(sys.argv) == 1:
        sys.argv.append('--help')

    metadata = importlib.metadata.metadata(PROGRAM_NAME)
    description = metadata.get('Summary', 'No description available')

    parser = argparse.ArgumentParser(description=description)
    parser.add_argument('--version', action='store_true', help='Print the version of rsavesync')
    parser.add_argument(
        '--alias',
        dest='alias',
        metavar='<alias>',
        required=False,
        help='(Optional) Use this flag when running a non-steam game, not strictly required for steam games')

    parser.add_argument(
        '--command',
        dest='command',
        metavar='<command>',
        required=True,
        help='(Required) Pass the arguments to run the game'
    )

    args = parser.parse_args()

    if args.version:
        print_version_and_exit()

    env = os.environ.copy()

    logger = log_factory.create()

    logger.info(f'executing: {args.command}')
    logger.info('--------------------------')
    logger.info(f'environment variables:')
    logger.info(env)

    exit(0)


def print_version_and_exit():
    version = importlib.metadata.version(PROGRAM_NAME)
    print(f'conductor-cli {version}')
    sys.exit(0)


if __name__ == '__main__':
    main()
