package database

import (
	"log"

	"github.com/victorbischoff/structs-api/pkg/entities"
	"github.com/victorbischoff/structs-api/pkg/utilities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func loadDsn() string {
	config, err := utilities.LoadConfigPath("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	return config.DSN
}

func InitDb() error {
	var err error
	DB, err = gorm.Open(postgres.Open(loadDsn()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&entities.User{}, &entities.ContactInfo{})
	return nil
}
