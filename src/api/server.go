package main

import (
	databases "covid_admission_api/database"
	"covid_admission_api/entities"
	"covid_admission_api/routers"
	"log"

	_ "github.com/go-sql-driver/mysql"
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

	databases.RedisClient, err = databases.NewRedisClient()
	if err != nil {
		log.Fatal(err)

	}

	server := routers.NewRouter()

	server.Run(":8080")

}
