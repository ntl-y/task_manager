package service

import (
	taskmanager "task_manager"
	"task_manager/internal/repository"
)

type TaskService struct {
	repo *repository.Repository
}

func (s *TaskService) Create(task taskmanager.Task) (created taskmanager.TaskValidated, err error) {
	return s.repo.Task.Create(task)
}

func NewTaskService(repo *repository.Repository) *TaskService {
	return &TaskService{repo: repo}
}
