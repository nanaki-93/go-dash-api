package config

import (
	"fmt"
	"go-dash-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var (
	host     = Config("DB_HOST")
	user     = Config("DB_USER")
	password = Config("DB_PASSWORD")
	db       = Config("DB_NAME")
	port     = Config("DB_PORT")
)

func NewDatabaseConnection() *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, db, port)
	database, e := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if e != nil {
		panic(e)
	}
	log.Println("Connected")
	database.Logger = logger.Default.LogMode(logger.Silent)
	log.Println("running migrations")
	err := database.AutoMigrate(&model.Entity{}, &model.Schema{})
	if err != nil {
		log.Fatal("Error running migrations", err)
	}

	doGorm, err := database.DB()
	if err != nil {
		panic(err)
	}

	gormErr := doGorm.Ping()
	if gormErr != nil {
		panic(gormErr)
	}

	return database
}
