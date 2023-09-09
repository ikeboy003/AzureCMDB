package service

import (
	"azurecmdb/models"
	"encoding/json"
	"fmt"
	"os/exec"
)

func getNICsInASubscription(subscription models.AzureSubscription) ([]models.AzureNIC, error) {
	// Fetch NICs for the given subscription
	cmdNICs := exec.Command("az", "network", "nic", "list", "-o", "json")
	outputNICs, err := cmdNICs.CombinedOutput()
	if err != nil {
		return nil, err
	}

	var nics []models.AzureNIC
	if err := json.Unmarshal(outputNICs, &nics); err != nil {
		return nil, err
	}

	return nics, nil
}

func GetNICsAcrossSubscriptions() ([]models.AzureNIC, error) {

	subscriptions, err := GetAzureSubscriptions()
	if err != nil {
		return nil, err
	}

	var allNICs []models.AzureNIC

	// Iterate over each subscription
	for _, subscription := range subscriptions {

		// Set active subscription
		err := exec.Command("az", "account", "set", "--subscription", subscription.Name).Run()

		if err != nil {
			return nil, fmt.Errorf("failed to set active subscription: %s", err)
		}

		// Fetch NICs for the current subscription
		cmdNICs := exec.Command("az", "network", "nic", "list", "-o", "json")
		outputNICs, err := cmdNICs.CombinedOutput()
		if err != nil {
			return nil, err
		}

		var nics []models.AzureNIC
		if err := json.Unmarshal(outputNICs, &nics); err != nil {
			return nil, err
		}

		// Append to the allNICs slice
		allNICs = append(allNICs, nics...)
	}

	return allNICs, nil
}
