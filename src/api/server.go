package main

import (
	"covid_admission_api/database"
	"covid_admission_api/routers"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var err error

func main() {

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)

	}
	// setup SQL DB
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)

	}
	defer db.Close()
	db.AutoMigrate()

	// setup Redis Client
	redisClient, err := database.NewRedisClient()
	if err != nil {
		log.Fatal(err)

	}
	server := routers.NewRouter(db, redisClient)
	server.Start(":8080")

}
