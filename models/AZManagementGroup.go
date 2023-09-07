package models

type AzureManagementGroup struct {
	TenantID    string `json:"tenantId" gorm:"primaryKey;type:varchar(255);not null"`
	DisplayName string `json:"displayName" gorm:"column:display_name;type:varchar(255);not null"`
	GroupID     string `json:"id" gorm:"column:id;type:varchar(512);uniqueIndex;not null"`
	Name        string `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Type        string `json:"type" gorm:"column:type;type:varchar(255);not null"`
}

func (AzureManagementGroup) TableName() string {
	return "azure_management_groups"
}
