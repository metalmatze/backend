package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is the database used by the backend
var DB *gorm.DB

// ConnectDatabase connects to the database DB
func ConnectDatabase() error {
	var err error
	var db *gorm.DB

	// Check with database driver to use. If DB_HOST is set, assume postgresql
	_, ok := os.LookupEnv("DB_HOST")
	if ok {
		log.Println("DB_HOST is set, using postgresql")
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		log.Println("DB_HOST is not set, using sqlite database")
		db, err = gorm.Open(sqlite.Open("data/gorm.db"), &gorm.Config{})
	}

	if err != nil {
		log.Println(err.Error())
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(Budget{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(Account{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(Category{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(Envelope{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(Transaction{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(Allocation{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}
