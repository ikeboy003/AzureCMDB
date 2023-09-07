package models

type AzureManagementGroup struct {
	TenantID      string              `json:"tenantId" gorm:"primaryKey;type:varchar(255);not null"`
	DisplayName   string              `json:"displayName" gorm:"type:varchar(255);not null"`
	ResourceID    string              `json:"id" gorm:"type:varchar(512);uniqueIndex;not null"`
	ResourceName  string              `json:"name" gorm:"type:varchar(255);not null"`
	Type          string              `json:"type" gorm:"type:varchar(255);not null"`
	Subscriptions []AzureSubscription `gorm:"foreignKey:HomeTenantID"`
}

func (AzureManagementGroup) TableName() string {
	return "azure_management_groups"
}
