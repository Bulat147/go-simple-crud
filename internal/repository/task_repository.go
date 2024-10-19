package repository

import (
	"database/sql"
	"log"
	"simple-crud/internal/domain"
	ers "simple-crud/internal/errors"
)

type TaskRepository struct {
	DB *sql.DB
}

func (tr *TaskRepository) GetTaskById(id string) (*domain.Task, error) {
	task := &domain.Task{}
	query := "SELECT * FROM tasks where id = $1"

	err := tr.DB.QueryRow(query, id).Scan(&task.Id, &task.Title, &task.EndDate)
	if err != nil {
		log.Printf("parse query error: %s", err)
		return nil, ers.TaskNotFound
	}
	return task, nil
}
