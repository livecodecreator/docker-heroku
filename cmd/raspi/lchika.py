#!/usr/bin/env python3

import sys
import time
import random
import signal
import collections
import RPi.GPIO as GPIO

def timeline_handler(arg1, arg2):

    global ftime

    now = time.time()

    value = GPIO.input(PIN_INF_SENSOR)

    if value == GPIO.LOW:

        ftime = now + SEQUENCIAL_MODE_SECONDS

    if ftime < now:

        mode_random()

    else:

        mode_sequencial()

def mode_random():

    global stats
    pin = random.choice(pins)
    stat = stats.get(pin)
    stat = stati.get(stat)
    stats.update({pin: stat})
    GPIO.output(pin, stat)

def mode_sequencial():

    pinsq.rotate()
    pint = pinsq[0]

    for pin in pins:

        if pin == pint:

            GPIO.output(pin, GPIO.HIGH)

        else:

            GPIO.output(pin, GPIO.LOW)

TIMELINE_INTERVAL = 0.05

SEQUENCIAL_MODE_SECONDS = 3

PIN_INF_SENSOR = 4
PIN_LED_RED    = 17
PIN_LED_YELLOW = 22
PIN_LED_GREEN  = 24
PIN_LED_WHITE  = 25

pins = [
    PIN_LED_RED   ,
    PIN_LED_YELLOW,
    PIN_LED_GREEN ,
    PIN_LED_WHITE ,
]

pinsq = collections.deque(
    [pin for pin in pins for i in range(3)]
)

stats = {}

stati = {
    GPIO.LOW: GPIO.HIGH,
    GPIO.HIGH: GPIO.LOW,
}

ftime = 0

seqpin = 0

GPIO.setmode(GPIO.BCM)
GPIO.setup(PIN_INF_SENSOR, GPIO.IN)

for pin in pins:

    stats.update({pin: GPIO.HIGH})
    GPIO.setup(pin, GPIO.OUT, initial=GPIO.HIGH)

signal.signal(signal.SIGALRM, timeline_handler)
signal.setitimer(signal.ITIMER_REAL, TIMELINE_INTERVAL, TIMELINE_INTERVAL)

try:

    print("ctrl+c waitng...", end="", flush=True)

    while True:

        time.sleep(1)

except KeyboardInterrupt:

    print("")
    GPIO.cleanup()
    sys.exit(0)
