package service

import (
	"azurecmdb/models"
	"encoding/json"
	"fmt"
	"os/exec"
)

func GetAllResourceGroups() (map[string][]models.AzureResourceGroup, error) {
	// First, get all subscriptions
	subscriptions, err := GetAzureSubscriptions()
	if err != nil {
		return nil, err
	}

	resourceGroupsMap := make(map[string][]models.AzureResourceGroup)

	// Iterate over each subscription
	for _, subscription := range subscriptions {
		// Set active subscription
		err := exec.Command("az", "account", "set", "--subscription", subscription.SubscriptionID).Run()
		if err != nil {
			return nil, fmt.Errorf("failed to set active subscription: %s", err)
		}

		// Fetch resource groups for the current subscription
		cmd := exec.Command("az", "group", "list", "-o", "json")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return nil, err
		}

		var resourceGroups []models.AzureResourceGroup
		if err := json.Unmarshal(output, &resourceGroups); err != nil {
			return nil, err
		}

		for i := range resourceGroups {
			resourceGroups[i].SubscriptionID = subscription.SubscriptionID
		}

		resourceGroupsMap[subscription.SubscriptionID] = resourceGroups
	}

	return resourceGroupsMap, nil
}
