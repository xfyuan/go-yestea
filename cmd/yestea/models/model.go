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
	Name string `gorm:"column:name" json:"name"`}
