#!/usr/bin/env python3

import os
import time
import signal
import psutil
import requests
import threading

from http import HTTPStatus

def timeline_handler(a, b):

    try:

        post_status()

    except Exception:

        pass

def post_status():

    endpoint = "https://livecodecreator.herokuapp.com/raspi/status"
    payload = {
        "token": os.getenv("RASPI_TOKEN"),
        "cpu": float(psutil.cpu_percent()),
        "disk": float(psutil.disk_usage(path='/').percent),
        "memory": float(psutil.virtual_memory().percent),
        "bootTime": int(psutil.boot_time()),
    }
    requests.post(endpoint, json=payload)

def main():

    INTERVAL_SECONDS = 60
    timeline_handler(None, None)
    signal.signal(signal.SIGALRM, timeline_handler)
    signal.setitimer(signal.ITIMER_REAL, INTERVAL_SECONDS, INTERVAL_SECONDS)

    try:

        threading.Event().wait()

    except KeyboardInterrupt:

        pass

if __name__ == "__main__":

    exit(main())
