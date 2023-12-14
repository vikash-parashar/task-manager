// // handlers/task_handler.go
package handlers

// import (
// 	"encoding/json"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/go-chi/chi"
// 	"github.com/vikash-parashar/task-remainder/controllers"
// 	"github.com/vikash-parashar/task-remainder/helpers"
// 	"github.com/vikash-parashar/task-remainder/models"
// )

// // GetTasksForUser gets all tasks for a specific user
// func GetTasksForUser(w http.ResponseWriter, r *http.Request) {
// 	userID, err := helpers.GetUserIDFromContext(r.Context())
// 	if err != nil {
// 		http.Error(w, "Invalid User ID", http.StatusInternalServerError)
// 		return
// 	}

// 	tasks, err := controllers.GetAllTasksForUser(userID)
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	helpers.RespondJSON(w, tasks)
// }

// // GetTaskForUser gets a specific task by ID for a specific user
// func GetTaskForUser(w http.ResponseWriter, r *http.Request) {
// 	userID, err := helpers.GetUserIDFromContext(r.Context())
// 	if err != nil {
// 		http.Error(w, "Invalid User ID", http.StatusInternalServerError)
// 		return
// 	}

// 	taskID, err := strconv.Atoi(chi.URLParam(r, "id"))
// 	if err != nil {
// 		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
// 		return
// 	}

// 	task, err := controllers.GetTaskByIDForUser(taskID, userID)
// 	if err != nil {
// 		http.Error(w, "Task Not Found", http.StatusNotFound)
// 		return
// 	}

// 	helpers.RespondJSON(w, task)
// }

// // CreateTaskForUser creates a new task for a specific user
// func CreateTaskForUser(w http.ResponseWriter, r *http.Request) {
// 	userID, err := helpers.GetUserIDFromContext(r.Context())
// 	if err != nil {
// 		http.Error(w, "Invalid User ID", http.StatusInternalServerError)
// 		return
// 	}

// 	var task *models.Task
// 	err = json.NewDecoder(r.Body).Decode(&task)
// 	if err != nil {
// 		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
// 		return
// 	}

// 	// Set the current time as the created_at timestamp
// 	task.CreatedAt = time.Now()
// 	task.UserID = userID

// 	err = controllers.CreateTaskForUser(task)
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// }

// // UpdateTaskForUser updates an existing task by ID for a specific user
// func UpdateTaskForUser(w http.ResponseWriter, r *http.Request) {
// 	userID, err := helpers.GetUserIDFromContext(r.Context())
// 	if err != nil {
// 		http.Error(w, "Invalid User ID", http.StatusInternalServerError)
// 		return
// 	}

// 	taskID, err := strconv.Atoi(chi.URLParam(r, "id"))
// 	if err != nil {
// 		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
// 		return
// 	}

// 	var task *models.Task
// 	err = json.NewDecoder(r.Body).Decode(&task)
// 	if err != nil {
// 		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
// 		return
// 	}

// 	err = controllers.UpdateTaskForUser(taskID, userID, task)
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// // DeleteTaskForUser deletes a task by ID for a specific user
// func DeleteTaskForUser(w http.ResponseWriter, r *http.Request) {
// 	userID, err := helpers.GetUserIDFromContext(r.Context())
// 	if err != nil {
// 		http.Error(w, "Invalid User ID", http.StatusInternalServerError)
// 		return
// 	}

// 	taskID, err := strconv.Atoi(chi.URLParam(r, "id"))
// 	if err != nil {
// 		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
// 		return
// 	}

// 	err = controllers.DeleteTaskForUser(taskID, userID)
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }
