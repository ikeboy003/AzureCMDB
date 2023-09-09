package models

type AzureResource struct {
	ResourceID        string `json:"id" gorm:"primaryKey;type:varchar(512);not null;uniqueIndex"`
	ResourceName      string `json:"name" gorm:"type:varchar(255);not null;unique"`
	Location          string `json:"location" gorm:"type:varchar(255);not null"`
	ResourceGroupName string `json:"resourceGroup" gorm:"type:varchar(255);not null"`
	ResourceType      string `json:"type,omitempty" gorm:"type:varchar(255)"`
}
