package models

type AzureVirtualMachine struct {
	AzureResource                        // Embedded AzResource
	Tags          map[string]interface{} `json:"tags" gorm:"type:jsonb"`
}

func (AzureVirtualMachine) TableName() string {
	return "azure_virtual_machine"
}
