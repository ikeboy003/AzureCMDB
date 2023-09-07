package models

type AzureVirtualMachine struct {
	ResourceID        string                 `gorm:"primaryKey;type:varchar(512);not null" json:"id"`
	Name              string                 `gorm:"type:varchar(255);not null" json:"name"`
	Location          string                 `gorm:"type:varchar(255);not null" json:"location"`
	Tags              map[string]interface{} `gorm:"type:jsonb" json:"tags"`
	ResourceType      string                 `gorm:"type:varchar(255);not null" json:"type"`
	ResourceGroupName string                 `gorm:"type:varchar(255);not null" json:"resourceGroup"`

	// Other fields specific to VMs
}
