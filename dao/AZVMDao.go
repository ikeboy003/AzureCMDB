package dao

import "azurecmdb/models"

type AZVMdao struct {
}

func (*AZVMdao) PerformTransaction(subVMMap map[string][]models.AzureVirtualMachine) error {

	tx := db.Begin()

	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, vm := range subVMMap {
		if err := tx.Create(&vm).Error; err != nil {
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
