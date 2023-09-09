package service

import (
	"azurecmdb/models"
)

//1. Get Subscriptions

// Doesnt need to Be Concurrent
func GetAzureSubscriptions() ([]models.AzureSubscription, error) {
	return getAzureSubscriptions()
}

//Data Here is at the Subscription level
//2. Get the resourceGroups Per Subscriptsions

func GetResourceGroupsInSubscription(subscription models.AzureSubscription) ([]models.AzureResourceGroup, error) {
	return getResourceGroupsInSubscription(subscription)
}

func GetNICsInASubscription(subscription models.AzureSubscription) ([]models.AzureNIC, error) {

	return getNICsInASubscription(subscription)
}

func InsertNicIntoVM(vm []models.AzureVirtualMachine, nics []models.AzureNIC) ([]models.AzureVirtualMachine, []models.AzureNIC) {
	return getVMsWithNICs(vm, nics)
}

func GetVMInResourceGroup(rg models.AzureResourceGroup) ([]models.AzureVirtualMachine, error) {
	return getVMInResourceGroup(rg)
}
