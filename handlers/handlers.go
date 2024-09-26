package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var db *sql.DB

// SetDB allows passing the database connection to the handlers
func SetDB(database *sql.DB) {
	db = database
}

// CreateTask handles POST requests to /tasks
// @Summary Create a new task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param   task  body Task  true  "New Task"
// @Success 200 {object} Task
// @Router /tasks [post]
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO tasks (title, description, completed) VALUES (?, ?, ?)`
	result, err := db.Exec(query, newTask.Title, newTask.Description, newTask.Completed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	newTask.ID = int(id)

	json.NewEncoder(w).Encode(newTask)
}

// GetTasks handles GET requests to /tasks
// @Summary Get all tasks
// @Tags tasks
// @Produce  json
// @Success 200 {array} Task
// @Router /tasks [get]
func GetTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, description, completed FROM tasks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	json.NewEncoder(w).Encode(tasks)
}

// GetTask handles GET requests to /tasks/{id}
// @Summary Get a task by ID
// @Tags tasks
// @Produce  json
// @Param   id  path int  true  "Task ID"
// @Success 200 {object} Task
// @Failure 404 {string} string "Task not found"
// @Router /tasks/{id} [get]
func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	query := "SELECT id, title, description, completed FROM tasks WHERE id = ?"
	row := db.QueryRow(query, id)

	var task Task
	err = row.Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
	if err == sql.ErrNoRows {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}

// UpdateTask handles PUT requests to /tasks/{id}
// @Summary Update a task by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param   id    path int  true  "Task ID"
// @Param   task  body Task  true  "Updated Task"
// @Success 200 {object} Task
// @Failure 404 {string} string "Task not found"
// @Router /tasks/{id} [put]
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var updatedTask Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE tasks SET title = ?, description = ?, completed = ? WHERE id = ?"
	result, err := db.Exec(query, updatedTask.Title, updatedTask.Description, updatedTask.Completed, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	updatedTask.ID = id
	json.NewEncoder(w).Encode(updatedTask)
}

// DeleteTask handles DELETE requests to /tasks/{id}
// @Summary Delete a task by ID
// @Tags tasks
// @Produce  json
// @Param   id  path int  true  "Task ID"
// @Success 200 {string} string "Task deleted successfully"
// @Failure 404 {string} string "Task not found"
// @Router /tasks/{id} [delete]
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM tasks WHERE id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
}
