#!/usr/bin/env python3

import time
import json
import logging
import websocket

def connect():

    try:

        url = "wss://livecodecreator.herokuapp.com/public/raspi/message"

        header = [
            "TokenCode: %s" % "XXXXXXXX"
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
    if data.get("message") == "keepalive": return
    logging.info("websocket message: {}".format(data.get("message")))

def main():

    logging.basicConfig(level=logging.DEBUG, format="%(levelname)s: %(message)s")

    while True:

        connect()
        time.sleep(5)

if __name__ == "__main__":

    main()
