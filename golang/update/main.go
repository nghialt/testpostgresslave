package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testpostgres/config"
	"testpostgres/models"
	"time"
)

func main() {
	db, err := gorm.Open(postgres.Open(config.MASTER_DSN), &gorm.Config{})
	if err != nil {
		log.Panic("cannot init db", err)
	}

	count := 0
	for {
		start := time.Time{}
		//result := db.Model(&models.Customer{}).Where("id > 30000").Update("count", count)
		result := db.Model(&models.Customer{}).Where("id > 0").Update("count", count)
		if result.Error != nil {
			log.Panic("failed to update ", db.Error)
		}
		log.Println("finish update ", time.Time{}.Sub(start).Milliseconds())
		count = (count + 1 ) % 10000
		time.Sleep(500 * time.Millisecond)
	}
}
