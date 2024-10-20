package service

import (
	"simple-crud/internal/domain"

	"github.com/google/uuid"
)

type TaskRepository interface {
	GetTaskById(id string) (*domain.Task, error)
	CreateTask(task domain.Task) error
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

func (ts *TaskService) AddTask(task domain.Task) (uuid.UUID, error) {
	taskId, err := uuid.NewRandom()
	if err != nil {
		return uuid.Nil, err
	}
	task.Id = taskId
	err = ts.TaskRepository.CreateTask(task)
	if err != nil {
		return uuid.Nil, err
	}
	return taskId, nil
}
