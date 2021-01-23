import logging
import statistics_server

if __name__ == "__main__":
    logging.basicConfig()
    statistics_server.serve()
    print("started")
