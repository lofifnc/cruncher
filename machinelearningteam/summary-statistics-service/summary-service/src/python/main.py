import logging
import statistics_server
import sys

if __name__ == "__main__":
    logging.basicConfig(stream=sys.stdout, level=logging.DEBUG)
    statistics_server.serve()
