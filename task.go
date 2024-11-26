package taskmanager

import "time"

type Task struct {
	Title       string `json:"title" db:"title"`
	Describtion string `json:"describtion" db:"describtion"`
}

type TaskValidated struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Describtion string    `json:"describtion" db:"describtion"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
