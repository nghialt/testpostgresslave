package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testpostgres/config"
	"testpostgres/models"
)

func main() {
	db, err := gorm.Open(postgres.Open(config.SLAVE_DSN), &gorm.Config{})
	if err != nil {
		log.Panic("cannot init db", err)
	}

	limit := 100000
	offset := 0

	for {
		fmt.Println("offset ", offset)
		var customers []models.Customer
		selectRes := db.Limit(limit).Offset(offset).Find(&customers)
		if selectRes.Error != nil {
			log.Panic("failed to select ", selectRes.Error)
		}
		if len(customers) == 0 {
			log.Println("finish select all")
			break
		}
		for _, customer := range customers {
			fmt.Println(customer)
		}
		offset += limit
	}
}
