package models

import (
	"time"
)

type Model struct {
	ID uint `gorm:"primary_key;column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

type Todo struct {
	Model
	Title string `gorm:"column:title" json:"title"`
	Description string `gorm:"column:description" json:"description"`
}

func (b *Todo) TableName() string {
	return "todos"
}