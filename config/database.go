package config

import (
	"fmt"
	"log"

	// _ "github.com/lib/pq"
	"github.com/niteshKrr/gin-framework/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect_DB() {
	connStr := "postgresql://my_db_owner:npg_eEKQ8cMoW7sU@ep-damp-bush-a11f0s53-pooler.ap-southeast-1.aws.neon.tech/my_db?sslmode=require"
	var err error
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Connected to database successfully! 🎉")
}

func MigrateDB() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
	fmt.Println("✅ Database migrated successfully!")
}
