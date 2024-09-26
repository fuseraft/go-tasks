package routes

import (
	"log"

	"go-tasks/config"
	"go-tasks/handlers"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// InitializeRoutes sets up the routes from config and initializes Swagger
func InitializeRoutes(r *mux.Router, cfg config.Config) {
	// Iterate through the routes in the config and register them
	for _, route := range cfg.Routes {
		switch route.Handler {
		case "createTask":
			r.HandleFunc(route.Path, handlers.CreateTask).Methods(route.Method)
		case "getTasks":
			r.HandleFunc(route.Path, handlers.GetTasks).Methods(route.Method)
		case "getTask":
			r.HandleFunc(route.Path, handlers.GetTask).Methods(route.Method)
		case "updateTask":
			r.HandleFunc(route.Path, handlers.UpdateTask).Methods(route.Method)
		case "deleteTask":
			r.HandleFunc(route.Path, handlers.DeleteTask).Methods(route.Method)
		default:
			log.Fatalf("Handler %s not defined", route.Handler)
		}
	}

	// Serve Swagger documentation
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
