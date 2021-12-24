package services

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../test.env")
	if err != nil {
		log.Fatal("test.env is not found")
	}
	code := m.Run()
	os.Exit(code)
}
