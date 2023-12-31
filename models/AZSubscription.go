package models

type AzureSubscription struct {
	SubscriptionID string               `json:"id" gorm:"primaryKey;type:varchar(512);not null"`
	CloudName      string               `json:"cloudName" gorm:"type:varchar(255);not null"`
	HomeTenantID   string               `json:"homeTenantId" gorm:"type:varchar(255);index;not null"`
	Name           string               `json:"name" gorm:"type:varchar(255);not null"`
	ResourceGroups []AzureResourceGroup `gorm:"foreignKey:SubscriptionID"`
}

func (AzureSubscription) TableName() string {
	return "azure_subscriptions"
}
