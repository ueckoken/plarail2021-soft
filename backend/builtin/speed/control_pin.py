import RPi.GPIO as GPIO
import http.server
import http
import urllib.parse
import time


class ControlPin(http.server.BaseHTTPRequestHandler):
    def do_GET(self):
        query: dict = urllib.parse.parse_qs(
            urllib.parse.urlparse(self.path).query)
        if "speed" not in query:
            self.send_response(http.HTTPStatus.BAD_REQUEST)
            return

        user_speed = float(query['speed'][0])
        if user_speed == 0.0:
            change_speed(0)
        elif 0 < user_speed <= 100:
            change_speed(60)
            time.sleep(0.1)
            change_speed(user_speed / 100 * 17.5 + 17.5)
        self.create_msg()

    def create_msg(self):
        # content_len  = int(self.headers.get("content-length"))
        # req_body = self.rfile.read(content_len).decode("utf-8")
        # body = "body: " + req_body + "\n"
        self.send_response(200)
        # self.send_header('Content-type', 'text/html; charset=utf-8')
        # self.send_header('Content-length', str(body.encode()))
        # self.end_headers()
        # self.wfile.write(body.encode())


def setup_gpio():
    pin_channel = 10
    frequency_hz = 50
    GPIO.setmode(GPIO.BOARD)
    GPIO.setup(pin_channel, GPIO.OUT)
    return GPIO.PWM(pin_channel, frequency_hz)


def change_speed(speed: float):
    global pwm
    pwm.ChangeDutyCycle(speed)


# start HTTP server
def start_server():
    server_addr = ("0.0.0.0", 8081)
    with http.server.HTTPServer(server_addr, ControlPin) as httpd:
        httpd.serve_forever()


def main():
    global pwm
    pwm = setup_gpio()
    pwm.start(0)
    start_server()


main()
