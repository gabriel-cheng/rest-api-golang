package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `json:"id" gorm:"primaryKey`
	name        string         `json:"name"`
	description string         `json:"description"`
	mediumPrice string         `json:"mediumPrice"`
	author      string         `json:"author"`
	imageURL    string         `json:"imgUrl"`
	CreatedAt   time.Time      `json:"created`
	UpdatedAt   time.Time      `json:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"`
}
