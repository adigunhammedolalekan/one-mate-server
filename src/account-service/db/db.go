package db

import (
	"github.com/jinzhu/gorm"
	"log"
)
import _ "github.com/jinzhu/gorm/dialects/postgres"

func CreateDbConnection(connectUri string) (*gorm.DB, error) {
	database, err := gorm.Open("postgres", connectUri)
	if err != nil {
		return nil, err
	}
	runMigration(database)
	return database, nil
}

func runMigration(database *gorm.DB) {
	log.Println("Running migration...")
}
