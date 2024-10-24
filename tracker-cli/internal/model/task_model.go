package model

import (
	"gorm.io/gorm"
)

type TaskStatus string

const (
	Todo       TaskStatus = "todo"
	InProgress TaskStatus = "in-progress"
	Done       TaskStatus = "done"
)

type Task struct {
	gorm.Model
	Name        string     `gorm:"type:varchar(100);not null;unique"`
	Description string     `gorm:"type:varchar(100)"`
	Status      TaskStatus `gorm:"type:varchar(20);default:'todo'"`
}
