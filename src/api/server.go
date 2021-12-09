package main

import (
	"covid_admission_api/routers"
)

func main() {

	server := routers.NewRouter()

	server.Run(":8080")

}
