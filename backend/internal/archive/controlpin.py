import RPi.GPIO as GPIO
import time
import sys

args = sys.argv
GPIO.setuo(args[1],GPIO.out)
servo = GPIO.PWM(args[1],50)
#TODO これあってる？
if args[2] == "1":
    servo.start(27)
else if args[2] == "0":
    servo.start(13)
