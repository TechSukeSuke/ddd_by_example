package main

import (
	"log"

	"github.com/k-shimizu04/ddd_by_example/config"
	"github.com/k-shimizu04/ddd_by_example/infrastructure"
	"github.com/k-shimizu04/ddd_by_example/infrastructure/database"
)

func main() {
	c, err := config.NewConfig(".env")
	if err != nil {
		log.Fatal(err)
	}

	db := database.NewDB(&c.RDB)
	// for example of other infrastructures
	// redis := redis.NewRedis(&c.Redis)
	// mail := mail.NewMail(&c.Mail)
	// aws := aws.NewAWS(&c.AWS)

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

	r := infrastructure.InitRouting(&c.APIServer, db.Connection)
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
