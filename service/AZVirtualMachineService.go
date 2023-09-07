package service

import (
	"azurecmdb/models"
	"encoding/json"
	"os/exec"
)

func GetAllVMs() (map[string][]models.AzureVirtualMachine, error) {
	subscriptions, err := GetAllResourceGroups() // Use your specific method to get subscriptions here
	if err != nil {
		return nil, err
	}

	allVMs := make(map[string][]models.AzureVirtualMachine)

	for _, subscriptionResourceGroups := range subscriptions {
		subscriptionID := subscriptionResourceGroups[0].SubscriptionID // Assuming it's the same for all resource groups

		var vms []models.AzureVirtualMachine

		for _, rg := range subscriptionResourceGroups {
			cmd := exec.Command("az", "vm", "list", "--resource-group", rg.ResourceGroupName, "--subscription", subscriptionID, "-o", "json")
			output, err := cmd.CombinedOutput()
			if err != nil {
				return nil, err
			}

			var rgVMs []models.AzureVirtualMachine
			if err := json.Unmarshal(output, &rgVMs); err != nil {
				return nil, err
			}

			vms = append(vms, rgVMs...)
		}

		allVMs[subscriptionID] = vms
	}

	return allVMs, nil
}
