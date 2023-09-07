package dao

import "azurecmdb/models"

type AZSubscriptionDAO struct {
}

func (*AZSubscriptionDAO) PerformTransaction(subscriptions []models.AzureSubscription) error {

	tx := db.Begin()

	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, sub := range subscriptions {
		if err := tx.Create(&sub).Error; err != nil {
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
