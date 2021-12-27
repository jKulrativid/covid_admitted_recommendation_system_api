from torchvision import models, transforms
from torch import load
import torch
import torch.nn as nn

from PIL import Image, ImageOps
from skimage import io

import numpy as np

import sys
import os
import json

from http.server import BaseHTTPRequestHandler, HTTPServer
import time

def load_model() -> models:
	model = models.resnet101

class ProcessHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header("Content-type", "text/html")
        self.end_headers()
        self.wfile.write(bytes("<html><head><title>https://pythonbasics.org</title></head>", "utf-8"))
        self.wfile.write(bytes("<p>Request: %s</p>" % self.path, "utf-8"))
        self.wfile.write(bytes("<body>", "utf-8"))
        self.wfile.write(bytes("<p>This is an example web server.</p>", "utf-8"))
        self.wfile.write(bytes("</body></html>", "utf-8"))


if __name__ == "__main__":    

	hostName = "localhost"
	serverPort = 2711

    webServer = HTTPServer((hostName, serverPort), ProcessHandler)
    print("Server started http://%s:%s" % (hostName, serverPort))

    try:
        webServer.serve_forever()
    except KeyboardInterrupt:
        pass

    webServer.server_close()
    print("Server stopped.")