package taskAssignment

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/saran-crayonte/task/database"
	"github.com/saran-crayonte/task/models"
)

// CreateTaskAssignment handles creating a new task assignment
//
//	@Summary		Create a new task assignment
//	@Description	Create a new task assignment with provided details
//	@Tags			Task Assignment
//	@Accept			json
//	@Produce		json
//
//	@Security		ApiKeyAuth
//	@Param			token			header		string					true	"API Key"
//
//	@Param			taskAssignment	body		models.TaskAssignment	true	"Task assignment details"
//	@Success		201				{object}	string					"Task assignment created successfully"
//	@Failure		400				{object}	string					"Invalid request payload"
//	@Failure		409				{object}	string					"Username doesn't exist / Task not found / Task is already assigned"
//	@Router			/api/v2/taskAssignment [post]
func CreateTaskAssignment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		taskAssignment := new(models.TaskAssignment)
		if err := json.Unmarshal(c.Body(), &taskAssignment); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}

		var checkAlreadyAssigned models.TaskAssignment
		database.DB.Where("task_id=?", taskAssignment.TaskID).First(&checkAlreadyAssigned)
		if checkAlreadyAssigned.ID != 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task is already assigned to somebody"})
		}

		var existingUser models.User
		database.DB.Where("username=?", taskAssignment.Username).First(&existingUser)
		if len(existingUser.Username) == 0 {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Username doesn't exists"})
		}

		var existingTask models.Task
		database.DB.Where("id=?", taskAssignment.TaskID).First(&existingTask)
		if existingTask.ID == 0 {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Task not found"})
		}

		estimatedHours := existingTask.EstimatedHours
		layout := "2006-01-02 3:04 PM"
		startDate, err := time.Parse(layout, taskAssignment.Start_Date)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid date time format"})
		}
		result := CalculateEndDate(startDate, estimatedHours)
		taskAssignment.Start_Date = startDate.Format("2006-01-02 3:04 PM")
		taskAssignment.End_Date = result.Format("2006-01-02 3:04 PM")
		database.DB.Create(taskAssignment)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Task Assignment created successfully",
		})
	}
}

func CalculateEndDate(startDate time.Time, estimatedHours int) time.Time {
	//workingHoursPerDay := 8
	endDate := startDate
	remainingHours := estimatedHours

	for remainingHours > 0 {

		if endDate.Weekday() == time.Saturday || endDate.Weekday() == time.Sunday || isHoliday(endDate) {
			endDate = endDate.AddDate(0, 0, 1)
			continue
		}

		if endDate.Hour() == 12 {
			endDate = endDate.Add(time.Hour)
			continue
		}

		if endDate.Hour() >= 18 {
			endDate = endDate.AddDate(0, 0, 1).Truncate(24 * time.Hour).Add(9 * time.Hour)
			continue
		}

		remainingHours--
		endDate = endDate.Add(time.Hour)
	}

	if endDate.Hour() > 18 {
		endDate = endDate.AddDate(0, 0, 1).Truncate(24 * time.Hour).Add(9 * time.Hour)
	}

	return endDate
}
func isHoliday(date time.Time) bool {
	holiday := new(models.Holiday)
	database.DB.Where("holiday_date = ?", date.Format("2006-01-02")).First(&holiday)
	return holiday.ID != 0
}

// GetTaskAssignment handles retrieving a task assignment by ID
//
//	@Summary		Get a task assignment by ID
//	@Description	Retrieve a task assignment by its ID
//	@Tags			Task Assignment
//	@Accept			json
//	@Produce		json
//
//	@Security		ApiKeyAuth
//	@Param			token	header		string					true	"API Key"
//
//	@Param			id		path		int						true	"Task Assignment ID"
//	@Success		200		{object}	models.TaskAssignment	"Task assignment retrieved successfully"
//	@Failure		400		{object}	string					"Invalid request payload"
//	@Failure		404		{object}	string					"Task assignment not found"
//	@Router			/api/v2/taskAssignment/{id} [get]
func GetTaskAssignment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type body struct {
			ID int
		}
		b := new(body)
		if err := json.Unmarshal(c.Body(), &b); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		var newTaskAssignment models.TaskAssignment
		database.DB.Where("id=?", b.ID).First(&newTaskAssignment)
		if newTaskAssignment.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task Assignment ID not found"})
		}
		return c.Status(fiber.StatusOK).JSON(newTaskAssignment)
	}
}

// UpdateTaskAssignment handles updating a task assignment by ID
//
//	@Summary		Update a task assignment by ID
//	@Description	Update an existing task assignment by its ID
//	@Tags			Task Assignment
//	@Accept			json
//	@Produce		json
//
//	@Security		ApiKeyAuth
//	@Param			token			header		string					true	"API Key"
//
//	@Param			taskAssignment	body		models.TaskAssignment	true	"Updated task assignment details"
//	@Success		200				{object}	string					"Task assignment updated successfully"
//	@Failure		400				{object}	string					"Invalid request payload"
//	@Failure		404				{object}	string					"Username doesn't exist / Task not found / Task assignment not found"
//	@Router			/api/v2/taskAssignment/{id} [put]
func UpdateTaskAssignment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		taskAssignment := new(models.TaskAssignment)
		if err := json.Unmarshal(c.Body(), &taskAssignment); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		var existingTaskAssignment models.TaskAssignment
		database.DB.Where("id=?", taskAssignment.ID).First(&existingTaskAssignment)
		if existingTaskAssignment.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task Assignment ID not found"})
		}

		var existingUser models.User
		database.DB.Where("username=?", taskAssignment.Username).First(&existingUser)
		if len(existingUser.Username) == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Username doesn't exists"})
		}

		var existingTask models.Task
		database.DB.Where("id=?", taskAssignment.TaskID).First(&existingTask)
		if existingTask.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task ID not found"})
		}

		var checkAlreadyAssigned models.TaskAssignment
		database.DB.Where("task_id=?", taskAssignment.TaskID).First(&checkAlreadyAssigned)
		if checkAlreadyAssigned.ID != 0 && checkAlreadyAssigned.ID != taskAssignment.ID {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task is already assigned to somebody"})
		}

		estimatedHours := existingTask.EstimatedHours
		layout := "2006-01-02 3:04 PM"
		startDate, err := time.Parse(layout, taskAssignment.Start_Date)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid date time format"})
		}
		result := CalculateEndDate(startDate, estimatedHours)
		taskAssignment.Start_Date = startDate.Format("2006-01-02 3:04 PM")
		taskAssignment.End_Date = result.Format("2006-01-02 3:04 PM")

		database.DB.Model(&existingTaskAssignment).Updates(taskAssignment)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Task Assignment Updated successfully"})
	}
}

// DeleteTaskAssignment handles deleting a task assignment by ID
//
//	@Summary		Delete a task assignment by ID
//	@Description	Delete an existing task assignment by its ID
//	@Tags			Task Assignment
//	@Accept			json
//	@Produce		json
//
//	@Security		ApiKeyAuth
//	@Param			token	header		string	true	"API Key"
//
//	@Param			id		path		int		true	"Task Assignment ID"
//	@Success		200		{object}	string	"Task assignment deleted successfully"
//	@Failure		400		{object}	string	"Invalid request payload"
//	@Failure		404		{object}	string	"Task assignment not found"
//	@Router			/api/v2/taskAssignment/{id} [delete]
func DeleteTaskAssignment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type body struct {
			ID int
		}
		b := new(body)
		if err := json.Unmarshal(c.Body(), &b); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}

		var existingTaskAssignment models.TaskAssignment
		database.DB.Where("id=?", b.ID).First(&existingTaskAssignment)
		if existingTaskAssignment.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task Assignment id not found"})
		}

		database.DB.Delete(&existingTaskAssignment)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Task Assignment entry deleted successfully",
		})
	}
}
