package service

import (
	"azurecmdb/models"
	"encoding/json"
	"fmt"
	"os/exec"
)

func GetAllVirtualNetworks() ([]models.AzureVirtualNetwork, error) {
	// First, get all subscriptions
	subscriptions, err := GetAzureSubscriptions()
	if err != nil {
		return nil, err
	}

	var allVnets []models.AzureVirtualNetwork

	// Iterate over each subscription
	for _, subscription := range subscriptions {
		// Set active subscription
		err := exec.Command("az", "account", "set", "--subscription", subscription.SubscriptionID).Run()
		if err != nil {
			return nil, fmt.Errorf("failed to set active subscription: %s", err)
		}

		// Fetch virtual networks for the current subscription
		cmdVnets := exec.Command("az", "network", "vnet", "list", "-o", "json")
		outputVnets, err := cmdVnets.CombinedOutput()
		if err != nil {
			return nil, err
		}

		var vnets []models.AzureVirtualNetwork
		if err := json.Unmarshal(outputVnets, &vnets); err != nil {
			return nil, err
		}

		// Populate the virtual networks with their associated subnets
		if err := PopulateVirtualNetworksWithSubnets(vnets); err != nil {
			return nil, fmt.Errorf("failed to populate vnets with subnets: %s", err)
		}

		// Append the populated vnets to the main slice
		allVnets = append(allVnets, vnets...)

		// Optional: Printing for debugging
		/*fmt.Println(subscription.Name)
		for _, v := range vnets {
			for _, v := range v.Subnets {
				fmt.Println(v.Name)
			}
		}

		*/
	}

	return allVnets, nil
}

func PopulateVirtualNetworksWithSubnets(vnets []models.AzureVirtualNetwork) error {
	// Iterate over each virtual network
	for i, vnet := range vnets {

		// Fetch subnets for the virtual network
		cmdSubnets := exec.Command("az", "network", "vnet", "subnet", "list", "-g", vnet.ResourceGroupName, "--vnet-name", vnet.ResourceName, "-o", "json")
		outputSubnets, err := cmdSubnets.CombinedOutput()
		if err != nil {
			return err
		}

		var subnets []models.AzureSubnet
		if err := json.Unmarshal(outputSubnets, &subnets); err != nil {
			return err
		}
		// Use the method receiver to set subnets for the vnet
		vnet.SetSubnets(subnets)

		for _, subnet := range vnet.Subnets {
			fmt.Println(subnet.Name)
		}
		// Update the original slice (not necessary if using pointers, but added for clarity)
		vnets[i] = vnet
	}

	return nil
}
