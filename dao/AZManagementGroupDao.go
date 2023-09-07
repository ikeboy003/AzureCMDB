package dao

import "azurecmdb/models"

type AZManagementGroupDAO struct {
}

func (*AZManagementGroupDAO) PerformTransaction(mgGroups []models.AzureManagementGroup) error {

	tx := db.Begin()

	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, mg := range mgGroups {
		if err := tx.Create(&mg).Error; err != nil {
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
