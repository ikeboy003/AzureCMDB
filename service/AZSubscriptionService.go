package service

import (
	"azurecmdb/models"
	"encoding/json"
	"os/exec"
)

func GetAzureSubscriptions() ([]models.AzureSubscription, error) {
	cmd := exec.Command("az", "account", "list", "-o", "json")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	var subscriptions []models.AzureSubscription
	if err := json.Unmarshal(output, &subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}
