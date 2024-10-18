package service

import (
	"simple-crud/internal/domain"

	"github.com/google/uuid"
)

type TaskRepository interface {
	GetTaskById(id string) *domain.Task
}

type TaskService struct {
	TaskRepository TaskRepository
}

func (ts *TaskService) GetTaskById(id uuid.UUID) domain.Task {
	return *ts.TaskRepository.GetTaskById(id.String())
}
