package models

import (
	"gorm.io/gorm"
	"www.github.com/NirajSalunke/sql-maker/config"
)

type QueryRequest struct {
	NaturalText string `json:"naturalText" gorm:"not null"`
	DatabaseID  uint   `json:"databaseId"`
	SqlEngine   string `json:"sqlEngine" default:"trino"`
}

type QueryResponse struct {
	Data []map[string]interface{} `json:"data"`
}

type Conversation struct {
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserInput  string `json:"userInput"`
	AiOutput   string `json:"aiOutput"`
	DatabaseID uint   `json:"databaseId"`
}

type Database struct {
	gorm.Model
	UserID        uint           `json:"userId"`
	DSN           string         `json:"dsn"`
	Conversations []Conversation `json:"conversations" gorm:"foreignKey:DatabaseID"`
}

type User struct {
	gorm.Model
	Name      string     `json:"name"`
	Email     string     `json:"email" gorm:"unique;not null"`
	Password  string     `json:"password"`
	Databases []Database `json:"databases" gorm:"foreignKey:UserID"`
}

func MigrateModels() {
	config.DB.AutoMigrate(&User{}, &Database{}, &Conversation{})
}
