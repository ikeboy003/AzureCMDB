package main

import (
	"azurecmdb/dao"
	service "azurecmdb/service"
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

	if res, err := service.GetAllVMs(); err != nil {
		fmt.Println(err)
	} else {
		vmDao := dao.AZVMdao{}
		vmDao.PerformTransaction(res)

	}

}
