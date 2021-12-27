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

def load_model() -> models:
	model = models.resnet101