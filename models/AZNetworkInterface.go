package models

type AzureNIC struct {
	AzureResource                   // Embedded AzResource
	ETag          string            `json:"etag" gorm:"type:varchar(255)"` // Ensure it's unique
	IPConfigs     []IPConfiguration `json:"ipConfigurations" gorm:"foreignKey:AllocatedNicName"`
	VMName        *string           `json:"-"`
}

type IPConfiguration struct {
	AllocatedNicName          string `gorm:"type:varchar(255);not null"`
	ResourceID                string `json:"id" gorm:"primaryKey;type:varchar(512);not null;uniqueindex"`
	Name                      string `json:"name"`
	Primary                   bool   `json:"primary"`
	PrivateIPAddress          string `json:"privateIPAddress"`
	PrivateIPAddressVersion   string `json:"privateIPAddressVersion"`
	PrivateIPAllocationMethod string `json:"privateIPAllocationMethod"`
	ResourceGroup             string `json:"resourceGroup"`
	Type                      string `json:"type"`
	Etag                      string `json:"etag"`
}
