package service

import (
	"azurecmdb/models"
	"encoding/json"
	"os/exec"
)

func GetAzureManagementGroups() ([]models.AzureManagementGroup, error) {
	cmd := exec.Command("az", "account", "management-group", "list", "-o", "json")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	var groups []models.AzureManagementGroup
	if err := json.Unmarshal(output, &groups); err != nil {
		return nil, err
	}

	return groups, nil
}
