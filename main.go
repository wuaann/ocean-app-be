package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {

	dsn := os.Getenv("POSTGRES_CONN_STRING")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	print(err, db)
}
