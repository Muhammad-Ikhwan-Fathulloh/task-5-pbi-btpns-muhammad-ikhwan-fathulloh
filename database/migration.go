package database

import (
	"log"
	"task-5-pbi-btpns-muhammad-ikhwan-fathulloh/models"
)

func Migrate() {
	Instance.AutoMigrate(&models.User{}, &models.Photo{})
	log.Println("Database Migration Completed!")
}
