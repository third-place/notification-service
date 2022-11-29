package main

import (
	"github.com/joho/godotenv"
	"github.com/third-place/notification-service/internal/db"
	"github.com/third-place/notification-service/internal/entity"
)

func main() {
	_ = godotenv.Load()
	conn := db.CreateDefaultConnection()
	conn.AutoMigrate(
		&entity.Post{},
	)
}
