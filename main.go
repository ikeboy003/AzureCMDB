package main

import (
	"azurecmdb/dao"
	"azurecmdb/service"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {

	if err := PersistResource(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Persisted Succesfully")
	}
}
func PersistFoundations() error {
	mgDAO := dao.AZManagementGroupDAO{}
	if managementGroups, err := service.GetAzureManagementGroups(); err != nil {
		return err
	} else {
		for _, mg := range managementGroups {
			fmt.Println(mg.ResourceName)
		}
		mgDAO.PerformTransaction(managementGroups)
	}

	subDAO := dao.AZSubscriptionDAO{}
	if subscriptions, err := service.GetAzureSubscriptions(); err != nil {
		return err
	} else {
		subDAO.PerformTransaction(subscriptions)
	}

	rgDAO := dao.AzResourceDAO{}
	if rg, err := service.GetAllResourceGroups(); err != nil {
		return err
	} else {
		rgDAO.PerformTransaction(rg)
	}

	return nil
}

func PersistResource() error {
	nicDAO := dao.AZNicDAo{}
	if nics, err := service.GetNICsAcrossSubscriptions(); err != nil {
		return err
	} else {
		nicDAO.PerformTransaction(nics)
	}

	return nil
}
