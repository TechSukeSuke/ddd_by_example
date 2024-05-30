package database

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/k-shimizu04/ddd_by_example/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	DBMS       string
	Host       string
	Port       string
	DBName     string
	Username   string
	Password   string
	Connection *gorm.DB
}

func NewDB(c *config.RDB) *DB {
	return newDB(&DB{
		DBMS:     c.DBMS,
		Host:     c.DBHost,
		Port:     c.DBPort,
		DBName:   c.DBName,
		Username: c.DBUser,
		Password: c.DBPass,
	})
}

func newDB(d *DB) *DB {
	// connection := fmt.Sprintf("%s@tcp(%s:%s)/%s", d.Password, "127.0.0.1", "3306", d.DBName)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", d.Username, d.Password, d.Host, d.Port, d.DBName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Tokyo")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	db, err := connect(dsn, 100)
	if err != nil {
		panic(err)
	}

	d.Connection = db
	return d
}

func connect(dsn string, count int) (*gorm.DB, error) {
	for count > 1 {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			time.Sleep(time.Second * 2)
			count--
			log.Printf("failed connection. message: %s", err.Error())
			log.Printf("retry... count: %v\n", count)

			continue
		}

		return db, nil
	}

	return nil, errors.New("coudn't connecte database")
}
