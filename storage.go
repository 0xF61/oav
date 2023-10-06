package main

import (
	"log"

	"github.com/AlienVault-OTX/OTX-Go-SDK/src/otxapi"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Store *PostgresStore

type PostgresStore struct {
	db *gorm.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	dsn := "host=192.168.254.2 user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	db.AutoMigrate(&otxapi.PulseDetail{})

	return &PostgresStore{
		db: db,
	}, nil

}
