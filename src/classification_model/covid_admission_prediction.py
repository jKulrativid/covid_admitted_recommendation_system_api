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

class PredictionTools:
	def __init__(self, model : torch.nn.Module, transform : transforms):
		self._Model = model
		self._Transform = transform

	def predict(self, img) -> str:
		transformed_img = self._Transform(img)[None]
		raw_prediction = self._Model(transformed_img)
		prediction = "Admit" if torch.sigmoid(raw_prediction) > .5 else "Not-Admit"
		return prediction

def get_model() -> torch.nn.Module:
	model = models.resnet101()
	model.conv1 = nn.Conv2d(1, 64, kernel_size=7, stride=2, padding=0, bias=False)
	num_ftrs = model.fc.in_features
	model.fc = nn.Linear(num_ftrs, 1)

	pth_file_path = os.path.join("model-pth-file", "covid_classification_model.pth")
	model.load_state_dict(load(pth_file_path, map_location=torch.device('cpu')))
	return model

def get_transform() -> transforms:
	img_size = (150, 150)
	normalize_mean, normalize_std = [0.485], [0.225]
	transform = transforms.Compose([
		transforms.Resize(img_size),
		transforms.ToTensor(),
		transforms.Normalize(normalize_mean, normalize_std)
	])
	return transform

def get_prediction_tools():
	return PredictionTools(get_model(), get_transform())

def get_image(img_name : str) -> torch.tensor:
	img = io.imread(img_name)
	img = Image.fromarray(img)
	return img

def get_prediction(img_name : str, prediction_tools : PredictionTools) -> str:
	img = get_image(img_name)
	prediction = prediction_tools.predict(img)
	return prediction

def save_as_json(save_path : str, results : dict):
	with open(os.path.join(save_path, "result.json"), "w") as j:
		json.dump(results, j)

def process(img_path : str, prediction_tools : PredictionTools):
	results = dict()
	for filename in os.listdir(img_path):
		# File in this directory is already preprocessed by our API. (no need to recheck)
		file_path = os.path.join(img_path, filename)
		results[filename] = get_prediction(file_path, prediction_tools)

	save_as_json(img_path, results)


def get_image_path():
	err_message = "File Path Error : No File Path Given"
	if len(sys.argv) < 2:
		sys.exit(err_message)
	
	return sys.argv[1]


def __main():
	#img_path = get_image_path()
	img_path = os.path.join(".", "test_img")
	prediction_tools = get_prediction_tools()
	process(img_path, prediction_tools)


if __name__ == '__main__':
	__main()
