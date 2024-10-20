package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-crud/internal/domain"
	ers "simple-crud/internal/errors"

	"simple-crud/internal/utils"

	"github.com/google/uuid"
)

type TaskService interface {
	GetTaskById(id uuid.UUID) (domain.Task, error)
	AddTask(task domain.Task) (uuid.UUID, error)
}

type TaskHandler struct {
	TaskService TaskService
}

func (th *TaskHandler) InitRoutes() {
	http.HandleFunc("/tasks", th.handleTasks)
}

func (th *TaskHandler) handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		th.handleGetTaskById(w, r)
	case "POST":
		th.handlePostTask(w, r)
	}

}

func (th *TaskHandler) handleGetTaskById(w http.ResponseWriter, r *http.Request) {
	taskId, err := utils.ParseUUIDFromRequestParam("taskId", r)
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

func (th *TaskHandler) handlePostTask(w http.ResponseWriter, r *http.Request) {
	task := &domain.Task{}
	err := json.NewDecoder(r.Body).Decode(task)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Unmarshal error: %s", err)
		return
	}

	id, err := th.TaskService.AddTask(*task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	idResponse, err := json.Marshal(domain.IdResponse{Id: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Marshal Error: %s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(idResponse)
}
