package main

import (
	"log"

	"github.com/k-shimizu04/ddd_by_example/config"
	"github.com/k-shimizu04/ddd_by_example/infrastructure/database"

	"gorm.io/gorm"
)

// go run . で起動した階層をカレントとする。
const seedsPath = "./database/seeds/"

func main() {
	c, err := config.NewConfig(".env")
	if err != nil {
		log.Fatal(err)
	}

	db := database.NewDB(&c.RDB)

	defer func() {
		if db.Connection != nil {
			sqlDB, err := db.Connection.DB()
			if err != nil {
				log.Fatal(err)
			}
			if err := sqlDB.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}()

	if err := db.Connection.Transaction(func(tx *gorm.DB) error {
		s := newSeeder(tx, seedsPath)
		if err := s.Execute(); err != nil {
			return err
		}
		log.Printf("Success all seed files!!!")
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
