package utils

import (
	"azurecmdb/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetGormConnection() (*gorm.DB, error) {
	connStr := "user=postgres password=postgres1 dbname=azurecmdb sslmode=verify-full"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&models.AzureVirtualMachine{})
	if err != nil {
		fmt.Println(err)
	}
	err = db.AutoMigrate(
		&models.AzureResourceGroup{})
	if err != nil {
		fmt.Println(err)
	}
	err = db.AutoMigrate(&models.AzureSubscription{})

	if err != nil {
		fmt.Println(err)
	}

	err = db.AutoMigrate(&models.AzureManagementGroup{})
	if err != nil {
		fmt.Println(err)
	}
	return db, nil
}
