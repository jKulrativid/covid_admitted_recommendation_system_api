package main

import (
	"covid_admission_api/routers"
)

func main() {

	server := routers.NewImageRouter()

	server.Run(":8080")

}
