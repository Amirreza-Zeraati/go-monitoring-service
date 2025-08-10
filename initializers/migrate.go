package initializers

import (
	"go-monitoring-service/models"
	"log"
)

func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Monitor{})
	if err != nil {
		log.Fatal("failed to migrate:", err)
	}
	log.Println("Database migrated successfully")
}
