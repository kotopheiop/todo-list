package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	//UserId    string    `json:"user_id"`
	Name     string `json:"Name"`
	Complete bool   `json:"Complete"`
}
