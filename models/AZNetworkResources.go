package models

type AzureVirtualNetwork struct {
	ResourceID        string        `json:"id" gorm:"primaryKey;type:varchar(512);not null"` // Primary key column
	Name              string        `json:"name" gorm:"type:varchar(255);uniqueIndex"`       // Unique index on the Name
	Location          string        `json:"location" gorm:"type:varchar(255)"`
	SubscriptionID    string        `gorm:"type:varchar(255)"`
	ResourceGroupName string        `json:"resourceGroup" gorm:"type:varchar(255);not null"`
	Subnets           []AzureSubnet `json:"subnets" gorm:"foreignKey:VNetID"` // Setting up one-to-many relationship
}

func (vnet *AzureVirtualNetwork) SetSubnets(subnets []AzureSubnet) {
	vnet.Subnets = subnets
}

type AzureSubnet struct {
	ResourceID    string `json:"id" gorm:"primaryKey;type:varchar(512);not null"`
	Name          string `json:"name" gorm:"type:varchar(255);uniqueIndex"` // Unique index on the Name
	AddressPrefix string `json:"addressPrefix" gorm:"type:varchar(255)"`
	VNetID        uint   `gorm:"index"` // Index on VNetID for faster lookups
}

type AzureNIC struct {
	ResourceID    string            `json:"resourceId" gorm:"primaryKey;type:varchar(512);not null"`
	Name          string            `json:"name" gorm:"type:varchar(255);not null"`
	Location      string            `json:"location" gorm:"type:varchar(255);not null"`
	ResourceGroup string            `json:"resourceGroup" gorm:"type:varchar(255);not null"`
	IPConfigs     []IPConfiguration `json:"ipConfigurations"`
}

type IPConfiguration struct {
	Name                      string `json:"name"`
	Primary                   bool   `json:"primary"`
	PrivateIPAddress          string `json:"privateIPAddress"`
	PrivateIPAddressVersion   string `json:"privateIPAddressVersion"`
	PrivateIPAllocationMethod string `json:"privateIPAllocationMethod"`
	ResourceGroup             string `json:"resourceGroup"`
	Type                      string `json:"type"`
}
