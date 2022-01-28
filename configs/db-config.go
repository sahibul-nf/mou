package configs

import (
	"fmt"
	"log"
	"moyu/campaign"
	"moyu/transaction"
	"moyu/user"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	// dsn := "host=localhost user= password= dbname= port= sslmode=disable TimeZone=Asia/Shanghai"

	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect database %s", err.Error())
	}

	db.AutoMigrate(&campaign.Campaign{}, &user.User{}, &campaign.CampaignImage{}, &transaction.Transaction{})
	fmt.Println("Successfully to connect database")
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}

	dbSQL.Close()
}
