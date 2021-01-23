import logging
import statistics_server
from prometheus_client import start_http_server
import sys

PROMETHEUS_PORT = 5001


if __name__ == "__main__":
    logging.basicConfig(stream=sys.stdout, level=logging.DEBUG)
    start_http_server(PROMETHEUS_PORT)
    statistics_server.serve()
