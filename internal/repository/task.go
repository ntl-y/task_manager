package repository

import (
	"fmt"
	taskmanager "task_manager"

	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task taskmanager.Task) (createdTask taskmanager.TaskValidated, err error) {
	query := fmt.Sprintf("INSERT INTO %s (TITLE, DESCRIBTION) VALUES ($1, $2) RETURNING ID, TITLE, DESCRIBTION, CREATED_AT", taskTable)
	if err = r.db.QueryRow(query, task.Title, task.Describtion).Scan(&createdTask.Id, &createdTask.Title, &createdTask.Describtion, &createdTask.CreatedAt); err != nil {
		return
	}
	return
}
