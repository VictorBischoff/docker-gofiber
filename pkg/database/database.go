package database

import (
	"log"

	"github.com/victorbischoff/structs-api/pkg/apistructs"
	"github.com/victorbischoff/structs-api/pkg/utilities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

//const dsn string = "host=100.84.8.61 user=victor password=victor dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Copenhagen"

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
	DB.AutoMigrate(&apistructs.User{}, &apistructs.ContactInfo{})
	return nil
}
