package main

import (
	"github.com/third-place/notification-service/internal/db"
	"github.com/third-place/notification-service/internal/entity"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	conn := db.CreateDefaultConnection()
	conn.AutoMigrate(
		&entity.Post{},
	)
}
