package repository

import (
	taskmanager "task_manager"

	"github.com/jmoiron/sqlx"
)

type Task interface {
	Create(task taskmanager.Task) (taskmanager.TaskValidated, error)
}

type Repository struct {
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Task: NewTaskRepository(db)}
}
