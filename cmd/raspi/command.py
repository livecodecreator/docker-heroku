#!/usr/bin/env python3

import os
import time
import json
import logging
import requests
import websocket

def connect():

    try:

        url = "wss://livecodecreator.herokuapp.com/raspi/command"
        # url = "ws://127.0.0.1:8080/raspi/command"

        raspi_token = os.getenv("RASPI_TOKEN")

        header = [
            "Authorization: %s" % raspi_token
        ]

        # websocket.enableTrace(True)
        ws = websocket.WebSocketApp(url,
            header = header,
            on_open = on_open,
            on_close = on_close,
            on_error = on_error,
            on_message = on_message,
        )

        ws.run_forever()

    except Exception as e:

        logging.error(e)

def on_open(ws):

    pass

def on_close(ws):

    logging.info("websocket close")

def on_error(ws, error):

    logging.error("websocket error: {}".format(error))

def on_message(ws, message):

    data = json.loads(message)
    command = data.get("command")
    function_name = "on_command_{}".format(command)

    if callable(globals().get(function_name)):
        globals().get(function_name)()

def on_command_keepalive():

    pass

def on_command_takeshot():

    logging.info("websocket command: takeshot")
    raspi_slack = os.getenv("RASPI_SLACK")
    payload = {
        "text": "HELLO",
    }
    requests.post(raspi_slack, json=payload)

def main():

    logging.basicConfig(level=logging.DEBUG, format="%(levelname)s: %(message)s")

    while True:

        connect()
        time.sleep(5)

if __name__ == "__main__":

    main()
