package model

import (
	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	Title   string `gorm:"index;not null"`
	Content string `gorm:"types:longtext"`
}

type Post struct {
	gorm.Model
	TopicId uint   `gorm:"not null"`
	Content string `gorm:"types:longtext"`
}
