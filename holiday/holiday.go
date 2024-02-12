package holiday

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/saran-crayonte/task/database"
	"github.com/saran-crayonte/task/models"
	"github.com/saran-crayonte/task/task"
)

// CreateHoliday handles creating a new holiday
//
//	@Summary		Create a new holiday
//	@Description	Create a new holiday with provided details
//	@Tags			Holiday Management
//	@Accept			json
//	@Produce		json
//
//	@Security		ApiKeyAuth
//	@Param			token	header		string			true	"API Key"
//
//	@Param			holiday	body		models.Holiday	true	"Holiday details"
//	@Success		201		{object}	string			"Holiday created successfully"
//	@Failure		400		{object}	string			"Invalid request payload"
//	@Failure		404		{object}	string			"Holiday already defined"
//	@Router			/api/v2/holiday [post]
func CreateHoliday() fiber.Handler {
	return func(c *fiber.Ctx) error {
		holiday := new(models.Holiday)
		if err := json.Unmarshal(c.Body(), &holiday); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		layout := "2006-01-02"
		_, err := time.Parse(layout, holiday.HolidayDate)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid date time format"})
		}
		var newHoliday models.Holiday
		database.DB.Where("holiday_date=?", holiday.HolidayDate).First(&newHoliday)
		if newHoliday.ID != 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Holiday already defined"})
		}
		database.DB.Create(&holiday)
		UpdateHolidayInAssignment()
		type UserResponse struct {
			Message     string `json:"message"`
			HolidayID   string `json:"holidayID"`
			HolidayName string `json:"holidayName"`
			HolidayDate string `json:"holidayDate"`
		}
		return c.Status(fiber.StatusCreated).JSON(UserResponse{
			Message:     "Holiday created successfully",
			HolidayID:   string(rune(holiday.ID)),
			HolidayName: holiday.HolidayName,
			HolidayDate: holiday.HolidayDate,
		})
	}
}

// GetHoliday handles retrieving a holiday by ID
//
//	@Summary		Get a holiday by ID
//	@Description	Retrieve a holiday by its ID
//	@Tags			Holiday Management
//	@Accept			json
//	@Produce		json
//
//	@Security		ApiKeyAuth
//	@Param			token	header		string			true	"API Key"
//
//	@Param			id		path		int				true	"Holiday ID"
//	@Success		200		{object}	models.Holiday	"Holiday retrieved successfully"
//	@Failure		400		{object}	string			"Invalid request payload"
//	@Failure		404		{object}	string			"Holiday not found"
//	@Router			/api/v2/holiday/{id} [get]
func GetHoliday() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type body struct {
			ID int
		}
		b := new(body)
		if err := json.Unmarshal(c.Body(), &b); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		var newHoliday models.Holiday
		database.DB.Where("id=?", b.ID).First(&newHoliday)
		if newHoliday.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Holiday not found"})
		}
		return c.Status(fiber.StatusOK).JSON(newHoliday)
	}
}

// UpdateHoliday handles updating a holiday by ID
//
//	@Summary		Update a holiday by ID
//	@Description	Update an existing holiday by its ID
//	@Tags			Holiday Management
//	@Accept			json
//	@Produce		json
//
//	@Security		ApiKeyAuth
//	@Param			token	header		string			true	"API Key"
//
//	@Param			holiday	body		models.Holiday	true	"Updated holiday details"
//	@Success		200		{object}	string			"Holiday updated successfully"
//	@Failure		400		{object}	string			"Invalid request payload"
//	@Failure		404		{object}	string			"Holiday not found"
//	@Router			/api/v2/holiday/{id} [put]
func UpdateHoliday() fiber.Handler {
	return func(c *fiber.Ctx) error {
		holiday := new(models.Holiday)
		if err := json.Unmarshal(c.Body(), &holiday); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		var newHoliday models.Holiday
		database.DB.Where("id=?", holiday.ID).First(&newHoliday)
		if newHoliday.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Holiday not found"})
		}
		layout := "2006-01-02"
		_, err := time.Parse(layout, holiday.HolidayDate)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid date time format"})
		}
		var existingHoliday models.Holiday
		database.DB.Where("holiday_date=?", holiday.HolidayDate).First(&existingHoliday)
		if existingHoliday.ID != 0 && existingHoliday.ID != holiday.ID {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Holiday already defined"})
		}
		database.DB.Model(&newHoliday).Updates(holiday)
		UpdateHolidayInAssignment()
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Holiday Updated Successfully"})
	}
}

// DeleteHoliday handles deleting a holiday by ID
//
//	@Summary		Delete a holiday by ID
//	@Description	Delete an existing holiday by its ID
//	@Tags			Holiday Management
//	@Accept			json
//	@Produce		json
//
//	@Security		ApiKeyAuth
//	@Param			token	header		string	true	"API Key"
//
//	@Param			id		path		int		true	"Holiday ID"
//	@Success		200		{object}	string	"Holiday deleted successfully"
//	@Failure		400		{object}	string	"Invalid request payload"
//	@Failure		404		{object}	string	"Holiday not found"
//	@Router			/api/v2/holiday/{id} [delete]
func DeleteHoliday() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type body struct {
			ID int
		}
		b := new(body)
		if err := json.Unmarshal(c.Body(), &b); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		var newHoliday models.Holiday
		database.DB.Where("id=?", b.ID).First(&newHoliday)
		if newHoliday.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Holiday not found"})
		}
		database.DB.Delete(&newHoliday)
		UpdateHolidayInAssignment()
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Holiday deleted successfully",
		})
	}
}

func UpdateHolidayInAssignment() {
	var taskAssignments []models.TaskAssignment
	database.DB.Find(&taskAssignments)
	for _, taskAssignment := range taskAssignments {
		var findTask models.Task
		database.DB.Where("id=?", taskAssignment.TaskID).First(&findTask)
		if findTask.ID != 0 {
			task.UpdatesInTaskAssignment(findTask.ID, findTask.EstimatedHours)
		}
	}
}

// DisplayAllHolidays handles retrieving all holidays
//
//	@Summary		Get all holidays
//	@Description	Retrieve all holidays
//	@Tags			Holiday Management
//	@Accept			json
//	@Produce		json
//
//	@Security		ApiKeyAuth
//	@Param			token	header		string			true	"API Key"
//
//	@Success		200		{object}	models.Holiday	"Holiday retrieved successfully"
//	@Router			/api/v2/task [get]
func DisplayAllHolidays() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var holiday []models.Holiday
		database.DB.Find(&holiday)
		return c.Status(fiber.StatusOK).JSON(holiday)
	}
}
