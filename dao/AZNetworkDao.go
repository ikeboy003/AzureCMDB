package dao

import "azurecmdb/models"

type AzNetworkDAO struct {
}

func (AzNetworkDAO) PerformTransaction(subRgMap map[string][]models.AzureResourceGroup) error {

	tx := db.Begin()

	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, rg := range subRgMap {
		if err := tx.Create(&rg).Error; err != nil {
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
