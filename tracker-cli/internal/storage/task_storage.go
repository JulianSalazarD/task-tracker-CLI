package storage

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var once sync.Once

type TaskStorage struct {
	db *gorm.DB
}

// Instance TaskStorage
func NewTaskStorage() *TaskStorage {
	return &TaskStorage{}
}

// NewSQliteDB creates a new connection to the database
func (ts *TaskStorage) NewSQliteDB() {
	once.Do(func() {
		var err error
		ts.db, err = gorm.Open(sqlite.Open("taskmanager.db"), &gorm.Config{})
		if err != nil {
			log.Fatalf("can't open connection to database:\n %v", err)
		}
		fmt.Println("Connection to dabase has been established")
	})
}

func (ts *TaskStorage) DB() *gorm.DB {
	return ts.db
}
