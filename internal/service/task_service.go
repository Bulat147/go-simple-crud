package service

import (
	"simple-crud/internal/domain"

	"github.com/google/uuid"
)

type TaskRepository interface {
	GetTaskById(id string) (*domain.Task, error)
}

type TaskService struct {
	TaskRepository TaskRepository
}

func (ts *TaskService) GetTaskById(id uuid.UUID) (domain.Task, error) {
	task, err := ts.TaskRepository.GetTaskById(id.String())
	if err != nil {
		return domain.Task{}, err
	}
	return *task, nil
}
