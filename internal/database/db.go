package database

import (
	"log"
	// "gorm.io/driver/postgres"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	var Db *gorm.DB

	// Postgres connection
	// url := "host=localhost user=postgres password=postgres dbname=student_db port=5433 sslmode=disable"

	// db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("failed to connect to database:", err)
	// }

	dsn := "root:Vedant@12345@tcp(localhost:3306)/user_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	Db = db

	log.Println("Database connected", Db)

	return Db
}
