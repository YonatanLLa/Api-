package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yonatanlla/crud-go/db"
	"github.com/yonatanlla/crud-go/models"
	"github.com/yonatanlla/crud-go/routes"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandle)

	// ROUTES USER

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHanlder).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	// ROUTES TASK

	r.HandleFunc("/task", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/task/{id}", routes.GetByIdTaskHandler).Methods("GET")
	r.HandleFunc("/task", routes.PostTaskHandler).Methods("POST")
	r.HandleFunc("/task/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
