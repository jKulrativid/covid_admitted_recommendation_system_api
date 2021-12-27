package database

import (
	"covid_admission_api/entities"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database interface {
	AutoMigrate() error
	GetConnection() *gorm.DB
	Close()
}

type database struct {
	db *gorm.DB
}

func NewDatabase() (Database, error) {
	fmt.Println("Initializing database...")
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3306"
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "root"
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return nil, fmt.Errorf("DB_PASSWORD is not set")
	}
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		return nil, fmt.Errorf("DB_NAME is not set")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}))
	if err != nil {
		return nil, err
	}
	return &database{db}, nil
}

func (d *database) GetConnection() *gorm.DB {
	return d.db
}

func (d *database) AutoMigrate() error {
	return d.db.AutoMigrate(&entities.User{})
}

func (d *database) Close() {
	db, _ := d.db.DB()
	db.Close()
}
