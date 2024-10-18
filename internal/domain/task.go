package domain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	EndDate time.Time `json:"endDate"`
}
