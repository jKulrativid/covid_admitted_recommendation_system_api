package main

import (
	databases "covid_admission_api/database"
	"covid_admission_api/entities"
	"covid_admission_api/routers"
	"log"

	"github.com/jinzhu/gorm"
)

var err error

func main() {

	databases.DB, err = gorm.Open("mysql", databases.DbURL(databases.BuildDBConfig()))
	if err != nil {
		log.Fatal(err)

	}
	defer databases.DB.Close()
	databases.DB.AutoMigrate(&entities.User{})

	databases.RedisClient = databases.NewRedisClient()

	server := routers.NewRouter()

	server.Run(":8080")

}
