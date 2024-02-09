package holiday

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/saran-crayonte/task/database"
	"github.com/saran-crayonte/task/models"
)

// CreateHoliday handles creating a new holiday
//
//	@Summary		Create a new holiday
//	@Description	Create a new holiday with provided details
//	@Tags			Holiday Management
//	@Accept			json
//	@Produce		json
//
// @Security ApiKeyAuth
// @Param token header string true "API Key"
//
//	@Param			holiday	body		models.Holiday	true	"Holiday details"
//	@Success		201		{object}	models.Holiday	"Holiday created successfully"
//	@Failure		400		{object}	string			"Invalid request payload"
//	@Failure		404		{object}	string			"Holiday already defined"
//	@Router			/api/v2/holiday [post]
func CreateHoliday() fiber.Handler {
	return func(c *fiber.Ctx) error {
		holiday := new(models.Holiday)
		if err := json.Unmarshal(c.Body(), &holiday); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		var newHoliday models.Holiday
		database.DB.First(&newHoliday, holiday.HolidayDate)
		if newHoliday.ID != 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Holiday already defined"})
		}
		database.DB.Create(&holiday)
		return c.Status(fiber.StatusCreated).JSON(holiday)
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
// @Security ApiKeyAuth
// @Param token header string true "API Key"
//
//	@Param			id	path		int				true	"Holiday ID"
//	@Success		200	{object}	models.Holiday	"Holiday retrieved successfully"
//	@Failure		400	{object}	string			"Invalid request payload"
//	@Failure		404	{object}	string			"Holiday not found"
//	@Router			/api/v2/holiday/{id} [get]
func GetHoliday() fiber.Handler {
	return func(c *fiber.Ctx) error {
		holiday := new(models.Holiday)
		if err := json.Unmarshal(c.Body(), &holiday); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		var newHoliday models.Holiday
		database.DB.First(&newHoliday, holiday.ID)
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
// @Security ApiKeyAuth
// @Param token header string true "API Key"
//
//	@Param			id		path		int				true	"Holiday ID"
//	@Param			holiday	body		models.Holiday	true	"Updated holiday details"
//	@Success		200		{object}	models.Holiday	"Holiday updated successfully"
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
		database.DB.First(&newHoliday, holiday.ID)
		if newHoliday.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Holiday not found"})
		}
		database.DB.Model(&newHoliday).Updates(holiday)
		return c.Status(fiber.StatusOK).JSON(newHoliday)
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
// @Security ApiKeyAuth
// @Param token header string true "API Key"
//
//	@Param			id	path		int		true	"Holiday ID"
//	@Success		200	{object}	string	"Holiday deleted successfully"
//	@Failure		400	{object}	string	"Invalid request payload"
//	@Failure		404	{object}	string	"Holiday not found"
//	@Router			/api/v2/holiday/{id} [delete]
func DeleteHoliday() fiber.Handler {
	return func(c *fiber.Ctx) error {
		holiday := new(models.Holiday)
		if err := json.Unmarshal(c.Body(), &holiday); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		var newHoliday models.Holiday
		database.DB.First(&newHoliday, holiday.ID)
		if newHoliday.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Holiday not found"})
		}
		database.DB.Delete(&newHoliday)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Holiday deleted successfully",
		})
	}
}
