package service

import (
	"azurecmdb/models"
	"encoding/json"
	"os/exec"
)

func GetAllVMsinSubscription() ([]models.AzureVirtualMachine, error) {
	cmd := exec.Command("az", "vm", "list", "-o", "json")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	var vms []models.AzureVirtualMachine
	if err := json.Unmarshal(output, &vms); err != nil {
		return nil, err
	}

	return vms, nil
}

func getVMInResourceGroup(rg models.AzureResourceGroup) ([]models.AzureVirtualMachine, error) {

	cmd := exec.Command("az", "vm", "list", "--resource-group", rg.ResourceGroupName, "--subscription", rg.SubscriptionID, "-o", "json")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	var vms []models.AzureVirtualMachine
	if err := json.Unmarshal(output, &vms); err != nil {
		return nil, err
	}

	return vms, nil
}

func getVMsWithNICs(nonNicVm []models.AzureVirtualMachine, allNics []models.AzureNIC) ([]models.AzureVirtualMachine, []models.AzureNIC) {
	// Map to index NICs by their ID for efficient lookup
	nicIDIndex := make(map[string]models.AzureNIC)
	for _, nic := range allNics {
		nicIDIndex[nic.ResourceID] = nic
	}

	// Iterate through VMs to associate NICs with their respective VMs
	for idx := range nonNicVm {
		for _, nicInt := range nonNicVm[idx].NetworkProfile.NetworkInterfaces {
			if nic, exists := nicIDIndex[nicInt.ID]; exists {
				nonNicVm[idx].Nics = append(nonNicVm[idx].Nics, nic)
				// Remove the NIC from the index after appending to prevent duplicates
				delete(nicIDIndex, nicInt.ID)
			}
		}
	}

	// Convert remaining NICs in nicIDIndex back to a slice
	allNicsWithoutVM := make([]models.AzureNIC, 0, len(nicIDIndex))
	for _, nic := range nicIDIndex {
		allNicsWithoutVM = append(allNicsWithoutVM, nic)
	}

	return nonNicVm, allNicsWithoutVM
}
