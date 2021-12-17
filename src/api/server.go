package main

import (
	databases "covid_admission_api/database"
	"covid_admission_api/entities"
	"covid_admission_api/routers"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var err error

func main() {

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)

	}
	// setup SQL DB
	databases.DB, err = gorm.Open("mysql", databases.DbURL(databases.BuildDBConfig()))
	if err != nil {
		log.Fatal(err)

	}
	defer databases.DB.Close()
	databases.DB.AutoMigrate(&entities.User{})

	// setup Redis Client
	databases.RedisClient, err = databases.NewRedisClient()
	if err != nil {
		log.Fatal(err)

	}
	server := routers.NewRouter()
	server.Run(":8080")

}
