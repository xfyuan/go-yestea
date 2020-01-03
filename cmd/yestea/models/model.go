package models

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `gorm:"column:title" json:"title"`
	Description string `gorm:"column:description" json:"description"`
}

func (b *Todo) TableName() string {
	return "todos"
}
