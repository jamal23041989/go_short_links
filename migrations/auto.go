package main

import (
	"github.com/jamal23041989/go_short_links/internal/link"
	"github.com/jamal23041989/go_short_links/internal/stat"
	"github.com/jamal23041989/go_short_links/internal/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&link.Link{}, &user.User{}, &stat.Stat{})
}
