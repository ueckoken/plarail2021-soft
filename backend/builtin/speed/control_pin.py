import RPi.GPIO as GPIO
import http.server
import http
import urllib.parse


class ControlPin(http.server.BaseHTTPRequestHandler):
    def do_GET(self):
        query: dict = urllib.parse.parse_qs(
            urllib.parse.urlparse(self.path).query)
        if "speed" not in query:
            self.send_response(http.HTTPStatus.BAD_REQUEST)
            return

        if not 0 <= float(query['speed'][0]) <= 100:
            change_speed(float(query["speed"][0]))
        self.send_response(http.HTTPStatus.OK)


def setup_gpio():
    pin_channel = 10
    frequency_hz = 50
    
    GPIO.setmode(GPIO.BOARD)
    GPIO.setup(pin_channel, GPIO.out)
    return GPIO.PWM(pin_channel, frequency_hz)


def change_speed(speed: float):
    global pwm
    pwm.changeDutyCycle(speed)


# start HTTP server
def start_server():
    server_addr = ("127.0.0.1", 8081)
    
    with http.server.HTTPServer(server_addr, ControlPin) as httpd:
       httpd.serve_forever()


def main():
    global pwm
    pwm = setup_gpio()
    start_server()

main()
