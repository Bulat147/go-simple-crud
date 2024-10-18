package main

import (
	"database/sql"
	"log"
	"net/http"
	"simple-crud/internal/delivery/http/handler"
	"simple-crud/internal/repository"
	"simple-crud/internal/service"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5433 user=bulat1 password=password dbname=crud-db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Open error: %s", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Ping error: %s", err)
	}

	rep := &repository.TaskRepository{DB: db}

	serv := &service.TaskService{TaskRepository: rep}

	var taskHandler = &handler.TaskHandler{TaskService: serv}

	http.HandleFunc("/task", taskHandler.GetTask)
	log.Println("Server is running on posrt :8080")
	http.ListenAndServe(":8080", nil)
}
