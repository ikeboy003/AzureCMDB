package models

type AzureVirtualNetwork struct {
	AzureResource                // Embedded AzResource
	SubscriptionID string        `gorm:"type:varchar(255)"`
	Subnets        []AzureSubnet `json:"subnets" gorm:"foreignKey:VNetID"` // Setting up one-to-many relationship
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
	AzureResource                   // Embedded AzResource
	ETag          string            `json:"etag" gorm:"type:varchar(255)"` // Ensure it's unique
	IPConfigs     []IPConfiguration `json:"ipConfigurations" gorm:"foreignKey:AllocatedNicName"`
}

type IPConfiguration struct {
	AllocatedNicName          *string `gorm:"type:varchar(255);not null"`
	ResourceID                string  `json:"id" gorm:"primaryKey;type:varchar(512);not null;uniqueindex"`
	Name                      string  `json:"name"`
	Primary                   bool    `json:"primary"`
	PrivateIPAddress          string  `json:"privateIPAddress"`
	PrivateIPAddressVersion   string  `json:"privateIPAddressVersion"`
	PrivateIPAllocationMethod string  `json:"privateIPAllocationMethod"`
	ResourceGroup             string  `json:"resourceGroup"`
	Type                      string  `json:"type"`
	Etag                      string  `json:"etag"`
}
