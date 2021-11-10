from torchvision import models, transforms
from torch import load, device
import torch
import torch.nn as nn

from PIL import Image, ImageOps
from skimage import io

import sys
import os
import json

global model, transform

def setup_model():
	model = models.resnet101()
	model.conv1 = nn.Conv2d(1, 64, kernel_size=7, stride=2, padding=0, bias=False)
	num_ftrs = model.fc.in_features
    model.fc = nn.Linear(num_ftrs, 1)

	pth_file_path = os.path.join("model-pth-file", "covid_classification_model.pth")
	model.load_state_dict(load(pth_file_path, map_location=torch.device('cpu')))


def get_preprocessed_image(img_name : str):
	img = io.imread(img_name)
	img = Image.fromarray(img)
	img = transforms.__call__(img)
	return img


def get_result(img_name : str) -> str:
	img = get_preprocessed_image(img_name)
	raw_prediction = model(img)
	prediction = 'Admitted' if torch.sigmoid(raw_prediction) > 0.5 else 'NotAdmitted'
	return prediction


def save_as_json(results : dict):
	with open("result.json", "w") as j:
		json.dumps(results, j)


def evaluate(img_path : str):
	results = dict()
	for filename in os.listdir(img_path):
		# File in this directory is already preprocessed by our API.
		result[filename] = get_result(filename)

	save_as_json(results)


def get_image_path():
	err_message = "File Path Error : No File Path Given"
	if len(sys.argv) < 2:
		sys.exit(err_message)


def main():
	setup_model()
	img_path = get_image_path()
	evaluate(img_path)

if __name__ == '__main__':
	main()
