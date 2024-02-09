package task

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/saran-crayonte/task/database"
	"github.com/saran-crayonte/task/models"
)

// CreateTasks handles creating a new task
//
//	@Summary		Create a new task
//	@Description	Create a new task with provided details
//	@Tags			Task Management
//	@Accept			json
//	@Produce		json
//	@Param			task	body		models.Task	true	"Task details"
//	@Success		201		{object}	models.Task	"Task created successfully"
//	@Failure		400		{object}	string		"Invalid request payload"
//	@Failure		409		{object}	string		"Task with the same title already exists"
//	@Router			/api/v2/task [post]
func CreateTasks() fiber.Handler {
	return func(c *fiber.Ctx) error {
		task := new(models.Task)
		if err := json.Unmarshal(c.Body(), &task); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}

		var existingTask models.Task
		database.DB.Where("title = ?", task.Title).First(&existingTask)
		if existingTask.ID != 0 {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Task with the same title already exists"})
		}

		database.DB.Create(&task)
		return c.Status(fiber.StatusCreated).JSON(task)
	}
}

// GetTasks handles retrieving a task by ID
//
//	@Summary		Get a task by ID
//	@Description	Retrieve a task by its ID
//	@Tags			Task Management
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int			true	"Task ID"
//	@Success		200	{object}	models.Task	"Task retrieved successfully"
//	@Failure		400	{object}	string		"Invalid request payload"
//	@Failure		404	{object}	string		"Task not found"
//	@Router			/api/v2/task/{id} [get]
func GetTasks() fiber.Handler {
	return func(c *fiber.Ctx) error {
		task := new(models.Task)
		if err := json.Unmarshal(c.Body(), &task); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		var newTask models.Task
		database.DB.First(&newTask, task.ID)
		if newTask.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
		}
		return c.Status(fiber.StatusOK).JSON(newTask)
	}
}

// UpdateTasks handles updating a task by ID
//
//	@Summary		Update a task by ID
//	@Description	Update an existing task by its ID
//	@Tags			Task Management
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"Task ID"
//	@Param			task	body		models.Task	true	"Updated task details"
//	@Success		200		{object}	models.Task	"Task updated successfully"
//	@Failure		400		{object}	string		"Invalid request payload"
//	@Failure		404		{object}	string		"Task not found"
//	@Router			/api/v2/task/{id} [put]
func UpdateTasks() fiber.Handler {
	return func(c *fiber.Ctx) error {
		task := new(models.Task)
		if err := json.Unmarshal(c.Body(), &task); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})

		}

		var existingTask models.Task
		database.DB.First(&existingTask, task.ID)
		if existingTask.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
		}
		database.DB.Model(&existingTask).Updates(task)
		return c.Status(fiber.StatusOK).JSON(existingTask)
	}
}

// DeleteTasks handles deleting a task by ID
//
//	@Summary		Delete a task by ID
//	@Description	Delete an existing task by its ID
//	@Tags			Task Management
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int		true	"Task ID"
//	@Success		200	{object}	string	"Task deleted successfully"
//	@Failure		400	{object}	string	"Invalid request payload"
//	@Failure		404	{object}	string	"Task not found"
//	@Router			/api/v2/task/{id} [delete]
func DeleteTasks() fiber.Handler {
	return func(c *fiber.Ctx) error {
		task := new(models.Task)
		if err := json.Unmarshal(c.Body(), &task); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		var newTask models.Task
		database.DB.First(&newTask, task.ID)
		if newTask.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
		}

		database.DB.Delete(&newTask)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Task deleted successfully",
		})
	}
}
