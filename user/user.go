package user

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/saran-crayonte/task/database"
	"github.com/saran-crayonte/task/models"
	"golang.org/x/crypto/bcrypt"
)

// Register handles user registration
//
//	@Summary		Register a new user
//	@Description	Register a new user with username, name, email, and password
//	@Tags			User Management
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.User	true	"User registration details"
//	@Success		201		{object}	string		"User registered successfully"
//	@Failure		400		{object}	string		"Invalid request payload"
//	@Failure		409		{object}	string		"Username already exists"
//	@Router			/api/user [post]
func Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		dat := new(models.User)
		if err := json.Unmarshal(c.Body(), &dat); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}

		var existingUser models.User
		database.DB.First(&existingUser, "username = ?", dat.Username)
		if len(existingUser.Username) != 0 {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "this username already exists"})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dat.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "hashing failed"})
		}

		newUser := models.User{
			Username: dat.Username,
			Name:     dat.Name,
			Email:    dat.Email,
			Password: string(hashedPassword),
		}
		database.DB.Create(&newUser)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "User registered successfully",
		})
	}
}

// Login handles user login
//
//	@Summary		Login user
//	@Description	Logs in a user with username and password
//	@Tags			User Management
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.User	true	"User login credentials"
//	@Success		202		{object}	models.User	"User authenticated"
//	@Failure		400		{object}	string		"Invalid request payload"
//	@Failure		401		{object}	string		"User not found / Password doesn't match"
//	@Router			/api/user/login [post]
func Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		returnObject := fiber.Map{
			"status": "",
			"msg":    "Something went wrong.",
		}

		var formData models.User

		if err := json.Unmarshal(c.Body(), &formData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}

		var user models.User

		database.DB.First(&user, "username = ?", formData.Username)
		if len(user.Username) == 0 {
			returnObject["msg"] = "User not found."

			return c.Status(fiber.StatusNotFound).JSON(returnObject)
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formData.Password))

		if err != nil {
			returnObject["msg"] = "Password doesnt match"
			return c.Status(fiber.StatusUnauthorized).JSON(returnObject)
		}

		token, err := GenerateToken(user)
		if err != nil {
			returnObject["msg"] = "Token creation error."
			return c.Status(fiber.StatusUnauthorized).JSON(returnObject)
		}

		returnObject["token"] = token
		returnObject["user"] = user
		returnObject["status"] = "OK"
		returnObject["msg"] = "User authenticated"
		return c.Status(fiber.StatusAccepted).JSON(returnObject)
	}
}

// UpdatePassword handles updating user password
//
//	@Summary		Update user password
//	@Description	Updates the password of the authenticated user
//	@Tags			User Management
//	@Accept			json
//	@Produce		json
//
// @Security ApiKeyAuth
// @Param token header string true "API Key"
//
//	@Param			user	body		models.User	true	"Update Password Request"
//	@Success		202		{object}	string		"Password updated successfully"
//	@Failure		400		{object}	string		"Invalid request payload"
//	@Failure		401		{object}	string		"Unauthorized"
//	@Failure		404		{object}	string		"Username doesn't exist"
//	@Router			/api/v2/user [put]
func UpdatePassword() fiber.Handler {
	return func(c *fiber.Ctx) error {

		username, ok := c.Locals("username").(string)
		if !ok {
			return fiber.ErrUnauthorized
		}

		var user models.User
		if err := json.Unmarshal(c.Body(), &user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		if username != user.Username {
			return fiber.ErrUnauthorized
		}

		var existingUser models.User
		database.DB.First(&existingUser, "username = ?", user.Username)
		if len(existingUser.Username) == 0 {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Username doesn't exists"})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "hashing failed"})
		}

		user.Password = string(hashedPassword)
		database.DB.Model(&existingUser).Updates(user)

		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message": "Password updated successfully",
		})
	}
}

// DeleteUser handles deleting user account
//
//	@Summary		Delete user account
//	@Description	Deletes the account of the authenticated user
//	@Tags			User Management
//	@Produce	json
//
// @Security ApiKeyAuth
// @Param token header string true "API Key"
//
//	@Param		user	body		models.User	true	"User deletion request"
//	@Success	200		{object}	string		"User deleted successfully"
//	@Failure	400		{object}	string		"Invalid request payload"
//	@Failure	401		{object}	string		"Unauthorized"
//	@Failure	404		{object}	string		"Username doesn't exist"
//	@Failure	500		{object}	string		"Internal Server Error"
//	@Router		/api/v2/user [delete]
func DeleteUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		username, ok := c.Locals("username").(string)
		if !ok {
			return fiber.ErrUnauthorized
		}

		var user models.User
		if err := json.Unmarshal(c.Body(), &user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		if username != user.Username {
			return fiber.ErrUnauthorized
		}
		var existingUser models.User
		database.DB.First(&existingUser, "username = ?", user.Username)
		if len(existingUser.Username) == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Username doesn't exists"})
		}
		err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Passward"})
		}
		database.DB.Delete(&existingUser)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User deleted successfully",
		})
	}
}

type CustomClaims struct {
	Email    string
	Username string

	jwt.RegisteredClaims
}

var secret string = "secret"

// GenerateToken generates JWT token for user authentication
func GenerateToken(user models.User) (string, error) {

	claims := CustomClaims{
		user.Email,
		user.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Println("Error in token signing.", err)
		return "", err
	}

	return t, nil

}

// ValidateToken validates JWT token
func ValidateToken(clientToken string) (*CustomClaims, string) {
	token, err := jwt.ParseWithClaims(clientToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err.Error()
	}

	claims, ok := token.Claims.(*CustomClaims)

	if !ok {
		return nil, "Invalid token claims"
	}

	return claims, ""
}

// Authenticate middleware authenticates user based on JWT token
func Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("token")

		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token not present."})
		}

		claims, msg := ValidateToken(token)

		log.Println(claims)

		if msg != "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": msg})
		}

		c.Locals("username", claims.Username)

		return c.Next()
	}
}

// RefreshToken handles refreshing authentication token
//
//	@Summary		Refresh authentication token
//	@Description	Refreshes the authentication token
//	@Tags			User Management
//
// @Security ApiKeyAuth
// @Param token header string true "API Key"
// @Param username header string true "Enter username"
//
//	@Produce		json
//	@Success		200	{object}	models.User	"Token refreshed successfully"
//	@Failure		401	{object}	string		"Unauthorized"
//	@Router			/api/v2/refreshToken [get]
func RefreshToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		returnObject := fiber.Map{
			"status": "OK",
			"msg":    "Refresh Token route",
		}

		username := c.Get("username")
		if username == "" {
			log.Println("Username key not found.")

			returnObject["msg"] = "username not found."
			return c.Status(fiber.StatusUnauthorized).JSON(returnObject)
		}

		tokenUsername := c.Locals("username")
		if username != tokenUsername {
			log.Println("username and token username mismatch.")
			returnObject["msg"] = "Invalid username."
			return c.Status(fiber.StatusUnauthorized).JSON(returnObject)
		}

		var user models.User
		database.DB.First(&user, "username = ?", username)

		if len(user.Username) == 0 {
			returnObject["msg"] = "Username not found."

			return c.Status(fiber.StatusBadRequest).JSON(returnObject)
		}

		token, err := GenerateToken(user)

		if err != nil {
			returnObject["msg"] = "Token creation error."
			return c.Status(fiber.StatusUnauthorized).JSON(returnObject)
		}

		returnObject["token"] = token
		returnObject["user"] = user

		return c.Status(fiber.StatusOK).JSON(returnObject)
	}
}
