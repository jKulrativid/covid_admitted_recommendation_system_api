from fastapi import FASTAPI, UploadFile
from fastapi.param_functions import File

app = FASTAPI()

@app.post('/process')
async def processHandler(file :UploadFile = File(...)):
	file.read()
	return 