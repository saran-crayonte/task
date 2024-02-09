package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saran-crayonte/task/holiday"
	"github.com/saran-crayonte/task/task"
	"github.com/saran-crayonte/task/taskAssignment"
	"github.com/saran-crayonte/task/user"
)

// SetupRoutes configures the API routes for the application
func SetupRoutes(app *fiber.App) {
	// Define base API group
	ap := app.Group("/api")

	// User routes
	ap.Post("/user", user.Register())
	ap.Post("/user/login", user.Login())

	// Authenticated API routes
	api := ap.Group("/v2", user.Authenticate())
	api.Get("/refreshToken", user.RefreshToken())
	api.Put("/user", user.UpdatePassword())
	api.Delete("/user", user.DeleteUser())

	// Task routes
	api.Post("/task", task.CreateTasks())
	api.Get("/task/:id", task.GetTasks())
	api.Put("/task/:id", task.UpdateTasks())
	api.Delete("/task/:id", task.DeleteTasks())

	// Task assignment routes
	api.Post("/taskAssignment", taskAssignment.CreateTaskAssignment())
	api.Get("/taskAssignment/:id", taskAssignment.GetTaskAssignment())
	api.Put("/taskAssignment/:id", taskAssignment.UpdateTaskAssignment())
	api.Delete("/taskAssignment/:id", taskAssignment.DeleteTaskAssignment())

	// Holiday routes
	api.Post("/holiday", holiday.CreateHoliday())
	api.Get("/holiday/:id", holiday.GetHoliday())
	api.Put("/holiday/:id", holiday.UpdateHoliday())
	api.Delete("/holiday/:id", holiday.DeleteHoliday())
}
