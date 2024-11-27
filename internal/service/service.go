package service

import (
	taskmanager "task_manager"
	"task_manager/internal/repository"
)

type Task interface {
	Create(task taskmanager.Task) (taskmanager.TaskValidated, error)
}

type Service struct {
	Task
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Task: NewTaskService(repo)}
}
