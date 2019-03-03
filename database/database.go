package database

import (
	"Smidgen/config"
	"Smidgen/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type(
	Pastes struct {
		Id string `gorm:"type:varchar(8);unique_index;primary_key"`
		Content string `gorm:"type:text"`
		CreationTime time.Time
	}
)

var(
	database gorm.DB
)

func Connect() {
	log.Println("Connecting to database")

	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.Database.Username,
		config.Conf.Database.Password,
		config.Conf.Database.Host,
		config.Conf.Database.Port,
		config.Conf.Database.Database)); if err != nil {
		panic(err)
	}

	database = *db

	log.Println("Connected to database")
}

func CreateTables() {
	log.Println("Creating tables")
	database.Exec("CREATE TABLE IF NOT EXISTS pastes(id VARCHAR(8) UNIQUE, content TEXT, creation_time TIMESTAMP);")
}

func PasteExists(id string) bool {
	var count int
	err := database.Where("id = ?", id).Find(Pastes{}).Count(&count); if err != nil {
		count = 0
	}

	return count > 0
}

func GenerateId() string {
	id := utils.GenerateId(8)

	for PasteExists(id) {
		id = utils.GenerateId(8)
	}

	return id
}

func CreatePaste(content string) string {
	id := GenerateId()

	paste := Pastes{
		Id: id,
		Content: content,
		CreationTime: time.Now(),
	}

	database.Create(&paste)

	return id
}

func GetContent(id string)  *string {
	var paste Pastes
	if database.Where("id = ?", id).First(&paste).RecordNotFound() {
		return nil
	}

	return &paste.Content
}
