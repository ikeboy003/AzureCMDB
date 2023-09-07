package models

import "strings"

type AzureResourceGroup struct {
	AzureID        string                 `json:"id" gorm:"tableName:resource_groups;primaryKey;type:varchar(512);not null"`
	Location       string                 `json:"location" gorm:"type:varchar(255);not null"`
	Name           string                 `json:"name" gorm:"type:varchar(255);not null"`
	Tags           map[string]interface{} `json:"tags" gorm:"type:jsonb"`
	Type           string                 `json:"type" gorm:"type:varchar(255);not null"`
	SubscriptionID string                 `gorm:"type:varchar(512);not null"`
	Subscription   AzureSubscription      `gorm:"references:SubscriptionID"`
}

func (AzureResourceGroup) TableName() string {
	return "azure_resource_groups"
}

func extractSubscriptionID(s string) string {
	parts := strings.Split(s, "/")
	if len(parts) >= 3 && parts[1] == "subscriptions" {
		return parts[2]
	}
	return ""
}

func (rg *AzureResourceGroup) SetSubscriptionIDFromAzureID() {
	rg.SubscriptionID = extractSubscriptionID(rg.AzureID)
}
