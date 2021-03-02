package main

import (
	"fmt"
	"log"
	"testpostgres/config"
	"testpostgres/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		step := 0
		for step < 5 {
			step = time.Now().Nanosecond() / 1000 % 1000
			fmt.Println("step ", step)
			time.Sleep(124 * time.Millisecond)
		}
		for i := 0; i < len(customers); i += step {
			customer := customers[i]
			fmt.Println("customer ", customer.ID, " count ", customer.Count)
		}
		//for _, customer := range customers {
		//	fmt.Println("customer ", customer.ID, " count ", customer.Count)
		//}
		//offset += limit
		time.Sleep(500 * time.Millisecond)
	}
}
