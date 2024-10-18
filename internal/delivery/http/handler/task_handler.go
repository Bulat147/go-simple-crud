package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-crud/internal/domain"

	"github.com/google/uuid"
)

type TaskService interface {
	GetTaskById(id uuid.UUID) domain.Task
}

type TaskHandler struct {
	TaskService TaskService
}

func (th *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var taskId, err = uuid.Parse(r.URL.Query().Get("taskId"))
		if err != nil {
			http.Error(w, "Failed to get taskId", http.StatusInternalServerError)
			log.Printf("Parse uuid error: %s", err)
			return
		}
		taskRs, err := json.Marshal(th.TaskService.GetTaskById(taskId))
		if err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
			log.Printf("Marshal Error: %s", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(taskRs)
	}
}
