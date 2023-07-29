package initializers

import (
	"fmt"
	"labpro/single-service/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(isAutoMigrate bool) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", Cfg.DB_HOST, Cfg.DB_USERNAME, Cfg.DB_PASSWORD, Cfg.DB_DATABASE, Cfg.DB_PORT)
	
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // logger.Silent, logger.Info
	})
	if err != nil {
		log.Fatal("Failed to connect to the database.")
	}
	fmt.Println("Successfully connected to the database.")

	if isAutoMigrate {
		err = DB.AutoMigrate(&models.Barang{}, &models.Perusahaan{}, &models.User{})
		if err != nil {
			log.Fatal("Failed to auto-migrate.")
		}
	}
}