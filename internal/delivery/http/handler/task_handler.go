package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"simple-crud/internal/domain"
	ers "simple-crud/internal/errors"

	"github.com/google/uuid"
)

type TaskService interface {
	GetTaskById(id uuid.UUID) (domain.Task, error)
}

type TaskHandler struct {
	TaskService TaskService
}

func (th *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		taskId, err := parseUUIDFromRequestParam("taskId", r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		task, err := th.TaskService.GetTaskById(taskId)
		if err == ers.TaskNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		taskRs, err := json.Marshal(task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Marshal Error: %s", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(taskRs)
	}
}

func parseUUIDFromRequestParam(param string, r *http.Request) (uuid.UUID, error) {
	idParam := r.URL.Query().Get(param)
	if idParam == "" {
		return uuid.Nil, fmt.Errorf("param %s is empty", param)
	}
	var id, err = uuid.Parse(idParam)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid %s param", param)
	}
	return id, nil
}
