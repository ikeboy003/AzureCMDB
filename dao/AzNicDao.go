package dao

import "azurecmdb/models"

type AZNicDAo struct {
}

func (AZNicDAo) PerformUpdateTransaction(nics []models.AzureNIC) error {

	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, nic := range nics {
		if err := tx.Create(&nic).Error; err != nil {
			tx.Rollback()
			return err

		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (AZNicDAo) PerformCreateTransaction(nics []models.AzureNIC) error {

	tx := db.Begin()

	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, nic := range nics {
		if err := tx.Create(&nic).Error; err != nil {
			tx.Rollback()
			return err

		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
