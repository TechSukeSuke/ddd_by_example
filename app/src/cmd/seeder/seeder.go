package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/gocarina/gocsv"
	"gorm.io/gorm"
)

type (
	Seeder struct {
		tx        *gorm.DB
		seedsPath string
	}
)

func newSeeder(tx *gorm.DB, seedsPath string) *Seeder {
	return &Seeder{
		tx:        tx,
		seedsPath: seedsPath,
	}

}

func (s *Seeder) Execute() error {
	files, err := os.ReadDir(s.seedsPath)
	if err != nil {
		return err
	}

	for _, f := range files {
		file, err := os.Open(fmt.Sprintf("%s/%s", s.seedsPath, f.Name()))
		if err != nil {
			return err
		}
		defer file.Close()

		csv, err := gocsv.CSVToMaps(file)
		if err != nil {
			return err
		}

		if len(csv) <= 0 {
			continue
		}

		name, _, _ := strings.Cut(f.Name(), ".csv")

		if err := s.tx.Table(name).Create(cast(csv)).Error; err != nil {
			log.Printf("During execution Create() using the %v, error occurred.", name)
			return err
		}

		log.Printf("Success seed: %v", file.Name())
	}

	return nil
}

func cast(csv []map[string]string) []map[string]interface{} {
	response := []map[string]interface{}{}

	for _, line := range csv {
		columns := make(map[string]interface{})

		iter := reflect.ValueOf(line).MapRange()
		for iter.Next() {
			columnName := iter.Key().String()
			columns[columnName] = line[columnName]
		}
		response = append(response, columns)
	}

	return response
}
