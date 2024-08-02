package main

import (
	"log"
	"net/http"

	"github.com/mariolopezdev/go-rest-server/internal/handlers"
	"github.com/mariolopezdev/go-rest-server/internal/taskstore"
)

func main() {
	store := taskstore.New()
	taskHandler := handlers.NewTaskHandler(store)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /tasks/", taskHandler.CreateTask)
	mux.HandleFunc("GET /tasks/{id}/", taskHandler.GetTask)
	mux.HandleFunc("PUT /tasks/{id}/", taskHandler.UpdateTask)
	mux.HandleFunc("DELETE /tasks/{id}/", taskHandler.DeleteTask)
	mux.HandleFunc("GET /tasks/", taskHandler.GetAllTasks)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println(store)
	log.Println("Server listening on http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
