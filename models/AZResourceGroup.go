package models

import "strings"

type AzureResourceGroup struct {
	AzureID        string                 `gorm:"tableName:resource_groups;primaryKey;type:varchar(512);not null" json:"id"`
	Location       string                 `gorm:"type:varchar(255);not null" json:"location"`
	Name           string                 `gorm:"type:varchar(255);not null" json:"name"`
	Tags           map[string]interface{} `gorm:"type:jsonb" json:"tags"`
	Type           string                 `gorm:"type:varchar(255);not null" json:"type"`
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
