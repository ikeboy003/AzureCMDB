package main

import (
	"azurecmdb/dao"
	"azurecmdb/service"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	vmDao := dao.AZVMdao{}
	GetData()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter VM name (or press 'q' to quit):")
		scanner.Scan() // Get user input from console
		vmName := scanner.Text()

		if vmName == "q" {
			break
		}

		if res, err := vmDao.IsThisNameTaken(vmName); err == nil && res {
			fmt.Println("The name is taken!")
		} else if err != nil {
			fmt.Println("error")
		} else {
			fmt.Println("This name is not taken")
		}
	}
}

func GetData() {

	mgDAO, vmDao := dao.AZManagementGroupDAO{}, dao.AZVMdao{}
	managementGroups, err := service.GetAzureManagementGroups()

	if err != nil {
		fmt.Println(err)
	} else {
		managementGroups[0].Subscriptions, err = service.GetAzureSubscriptions()

		if err != nil {
			fmt.Println("Couldnt Get Subscription")
		} else {

			for idx, subscription := range managementGroups[0].Subscriptions {

				err := exec.Command("az", "account", "set", "--subscription", subscription.SubscriptionID).Run()
				if err != nil {
					fmt.Println(err)
				}
				nics, err := service.GetNICsInASubscription(subscription)
				if err != nil {
					fmt.Println("Couldnt Get Nic in Sub: ", subscription.Name, err)
				}
				managementGroups[0].Subscriptions[idx].ResourceGroups, err = service.GetResourceGroupsInSubscription(subscription)
				if err != nil {
					fmt.Println("Couldnt Get Resource Groups in Sub: ", subscription.Name, err)
				}

				if err != nil {
					fmt.Println("Couldnt Get RG in Sub: ", subscription.Name, err)
				}
				vms, err := service.GetAllVMsinSubscription()
				if err != nil {
					fmt.Println("Couldnt Get VMs in Sub: ", subscription.Name, err)
				}

				vms, nics = service.InsertNicIntoVM(vms, nics)

				vmDao.PerformSliceTransaction(vms)

			}
		}
	}

	if err = mgDAO.PerformTransaction(managementGroups); err != nil {
		fmt.Println("couldnt Persist")
	}

}
