package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testpostgres/config"
	"testpostgres/models"
)

func main() {
	db, err := gorm.Open(postgres.Open(config.MASTER_DSN), &gorm.Config{})
	if err != nil {
		log.Panic("cannot init db", err)
	}
	//if res := db.Exec("drop table customers"); res.Error != nil {
	//	log.Panic("cannot drop table", res.Error)
	//}
	if err = db.AutoMigrate(models.Customer{}); err != nil {
		log.Panic("cannot auto migrate", err)
	}

	//total := 80000
	//fromID := 20000
	//for i := 0; i < total; i++ {
	//	customer := &models.Customer{
	//		Model: gorm.Model{
	//			ID: uint(fromID + i + 1),
	//		},
	//		FirstName: fmt.Sprintf("User %d", i+1),
	//	}
	//	result := db.Create(customer)
	//	if result.Error != nil {
	//		log.Panic("failed to create customer ", result.Error)
	//	}
	//	log.Println("finish create customer", customer.ID)
	//}
}
