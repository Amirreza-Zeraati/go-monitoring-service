package initializers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	// dsn := "host=localhost user=postgres password=postgres dbname=gorm port=5432"
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB, err = gorm.Open(sqlite.Open("C:/Users/arisu/Desktop/Go/go-monitoring-service/gonitor.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
}
