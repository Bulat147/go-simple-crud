package repository

import (
	"database/sql"
	"simple-crud/internal/domain"
)

type TaskRepository struct {
	DB *sql.DB
}

func (tr *TaskRepository) GetTaskById(id string) *domain.Task {
	task := &domain.Task{}
	query := "SELECT * FROM tasks where id = $1"

	err := tr.DB.QueryRow(query, id).Scan(&task.Id, &task.Title, &task.EndDate)
	if err != nil {
		return nil
	}
	return task
}
