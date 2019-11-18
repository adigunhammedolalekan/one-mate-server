package main

import (
	"github.com/adigunhammedolalekan/services/src/account-service/db"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env file. Quiting...")
	}
	database, err := db.CreateDbConnection(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	_ = database
}

func startRpcService(port string) error {

}