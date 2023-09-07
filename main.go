package main

import (
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

	if res1, err := service.GetNICsAcrossSubscriptions(); err != nil {
		fmt.Println(err, res1)
	}

}
