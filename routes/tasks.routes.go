package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yonatanlla/crud-go/db"
	"github.com/yonatanlla/crud-go/models"
)

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	tasksAll := db.DB.Find(&tasks)

	err := tasksAll.Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error geting users"))
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&tasks)

}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task models.Task

	json.NewDecoder(r.Body).Decode(&task)
	createTask := db.DB.Create(&task)
	err := createTask.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(&task)
}

func GetByIdTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task

	db.DB.First(&task, params)
	// err := getById.Error

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&task)

}
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task

	result := db.DB.Unscoped().Delete(&task, params)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))

}
