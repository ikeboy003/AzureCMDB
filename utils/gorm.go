package utils

import (
	"azurecmdb/models"
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

	if err = resetFoundationalResouces(db); err != nil {
		return nil, err
	}
	if err = resetResources(db); err != nil {
		return nil, err
	}

	return db, nil
}

func resetFoundationalResouces(db *gorm.DB) error {
	migrator := db.Migrator()
	if err := migrator.DropTable(&models.AzureManagementGroup{}, &models.AzureSubscription{}, &models.AzureResourceGroup{}); err != nil {
		return err
	}
	if err := migrator.CreateTable(&models.AzureManagementGroup{}, &models.AzureSubscription{}, &models.AzureResourceGroup{}); err != nil {
		return err
	}
	return nil
}
func resetResources(db *gorm.DB) error {
	migrator := db.Migrator()
	if err := migrator.DropTable(&models.AzureVirtualMachine{}, &models.AzureNIC{}, &models.AzureVirtualNetwork{}, &models.IPConfiguration{}); err != nil {
		return err
	}

	if err := migrator.CreateTable(&models.AzureVirtualMachine{}, &models.AzureNIC{}, &models.AzureVirtualNetwork{}, &models.IPConfiguration{}); err != nil {
		return err
	}

	return nil
}
