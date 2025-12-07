package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDb() *gorm.DB {
	var Db *gorm.DB

	url := "host=localhost user=postgres password=postgres dbname=student_db port=5433 sslmode=disable"
	
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	Db = db

	log.Println("Database connected", Db)

	return Db
}
