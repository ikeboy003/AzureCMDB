package models

type AzureVirtualMachine struct {
	ResourceID        string                 `json:"id" gorm:"primaryKey;type:varchar(512);not null"`
	Name              string                 `json:"name" gorm:"type:varchar(255);not null"`
	Location          string                 `json:"location" gorm:"type:varchar(255);not null"`
	Tags              map[string]interface{} `json:"tags" gorm:"type:jsonb"`
	ResourceType      string                 `json:"type" gorm:"type:varchar(255);not null"`
	ResourceGroupName string                 `json:"resourceGroup" gorm:"type:varchar(255);not null"`

	// Other fields specific to VMs
}
