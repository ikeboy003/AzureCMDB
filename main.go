package main

import (
	"azurecmdb/dao"
	"azurecmdb/service"
	"fmt"
	"log"
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

	mgDAO, vmDao, nicDao := dao.AZManagementGroupDAO{}, dao.AZVMdao{}, dao.AZNicDAo{}
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

				nicDao.PerformCreateTransaction(nics)
				vmDao.PerformSliceTransaction(vms)

			}
		}
	}

	if err = mgDAO.PerformTransaction(managementGroups); err != nil {
		fmt.Println("couldnt Persist")
	}

}
