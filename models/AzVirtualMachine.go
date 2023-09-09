package models

type AzureVirtualMachine struct {
	AzureResource                         // Embedded AzResource
	Tags           map[string]interface{} `json:"tags" gorm:"type:jsonb"`
	NetworkProfile NetworkProfile         `json:"networkProfile" gorm:"-"`
	Nics           []AzureNIC             `gorm:"foreignKey:VMName;references:ResourceName"`
}

func (AzureVirtualMachine) TableName() string {
	return "azure_virtual_machine"
}

type NetworkProfile struct {
	NetworkInterfaces []NetworkInterface `json:"networkInterfaces"`
}
type NetworkInterface struct {
	ID string `json:"id"`
}

type Resource interface {
	AzureNIC | AzureVirtualMachine | AzureVirtualNetwork
}
