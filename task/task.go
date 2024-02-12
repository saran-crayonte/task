package task

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/saran-crayonte/task/database"
	"github.com/saran-crayonte/task/models"
	"github.com/saran-crayonte/task/taskAssignment"
)

// CreateTasks handles creating a new task
//
//	@Summary		Create a new task
//	@Description	Create a new task with provided details
//	@Tags			Task Management
//	@Accept			json
//	@Produce		json
//
//	@Security		ApiKeyAuth
//	@Param			token	header		string		true	"API Key"
//
//	@Param			task	body		models.Task	true	"Task details"
//	@Success		201		{object}	string		"Task created successfully"
//	@Failure		400		{object}	string		"Invalid request payload"
//	@Router			/api/v2/task [post]
func CreateTasks() fiber.Handler {
	return func(c *fiber.Ctx) error {
		task := new(models.Task)
		if err := json.Unmarshal(c.Body(), &task); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}

		// var existingTask models.Task
		// database.DB.Where("title = ?", task.Title).First(&existingTask)
		// if existingTask.ID != 0 {
		// 	return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Task with the same title already exists"})
		// }

		database.DB.Create(&task)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Task Created successfully",
		})
	}
}

// GetTasks handles retrieving a task by ID
//
//	@Summary		Get a task by ID
//	@Description	Retrieve a task by its ID
//	@Tags			Task Management
//	@Accept			json
//	@Produce		json
//
//	@Security		ApiKeyAuth
//	@Param			token	header		string		true	"API Key"
//
//	@Param			id		path		int			true	"Task ID"
//	@Success		200		{object}	models.Task	"Task retrieved successfully"
//	@Failure		400		{object}	string		"Invalid request payload"
//	@Failure		404		{object}	string		"Task not found"
//	@Router			/api/v2/task/{id} [get]
func GetTasks() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type body struct {
			ID int
		}
		b := new(body)
		if err := json.Unmarshal(c.Body(), &b); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		var newTask models.Task
		database.DB.Where("id = ?", b.ID).First(&newTask)
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
//
//	@Security		ApiKeyAuth
//	@Param			token	header		string		true	"API Key"
//
//	@Param			task	body		models.Task	true	"Updated task details"
//	@Success		200		{object}	string		"Task updated successfully"
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
		database.DB.Where("id = ?", task.ID).First(&existingTask)
		if existingTask.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
		}

		// var existingTaskTitle models.Task
		// database.DB.Where("title = ?", task.Title).First(&existingTaskTitle)
		// if existingTask.ID != 0 {
		// 	return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Task with the same title already exists"})
		// }
		updatesInTaskAssignment(task.ID, existingTask.EstimatedHours)
		database.DB.Model(&existingTask).Updates(task)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Task updated successfully",
		})
	}
}

func updatesInTaskAssignment(id uint, est int) {
	taskAssign := new(models.TaskAssignment)
	database.DB.Where("task_id=?", id).First(&taskAssign)
	if taskAssign.ID != 0 {
		startDate, _ := time.Parse("2006-01-02 3:04 PM", taskAssign.Start_Date)
		result := taskAssignment.CalculateEndDate(startDate, est)
		newAssignment := models.TaskAssignment{
			ID:         taskAssign.ID,
			Username:   taskAssign.Username,
			TaskID:     taskAssign.TaskID,
			Start_Date: taskAssign.Start_Date,
			End_Date:   result.Format("2006-01-02 3:04 PM"),
		}
		database.DB.Model(&taskAssign).Updates(newAssignment)
	}
}

// DeleteTasks handles deleting a task by ID
//
//	@Summary		Delete a task by ID
//	@Description	Delete an existing task by its ID
//	@Tags			Task Management
//	@Accept			json
//	@Produce		json
//
//	@Security		ApiKeyAuth
//	@Param			token	header		string	true	"API Key"
//
//	@Param			id		path		int		true	"Task ID"
//	@Success		200		{object}	string	"Task deleted successfully"
//	@Failure		400		{object}	string	"Invalid request payload"
//	@Failure		404		{object}	string	"Task not found"
//	@Router			/api/v2/task/{id} [delete]
func DeleteTasks() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type body struct {
			ID int
		}
		b := new(body)
		if err := json.Unmarshal(c.Body(), &b); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		var newTask models.Task
		database.DB.Where("id = ?", b.ID).First(&newTask)
		if newTask.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
		}
		deleteInTaskAssignment(b.ID)
		database.DB.Delete(&newTask)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Task deleted successfully",
		})
	}
}

func deleteInTaskAssignment(ID int) {
	taskAssignment := new(models.TaskAssignment)
	database.DB.Where("task_id=?", ID).Delete(&taskAssignment)
}
