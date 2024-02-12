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

		if dat.Username == "" || dat.Name == "" || dat.Email == "" || dat.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		var existingUser models.User
		database.DB.Where("username = ?", dat.Username).First(&existingUser)
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

		type UserResponse struct {
			Message  string `json:"message"`
			Username string `json:"username"`
			Name     string `json:"name"`
			Email    string `json:"email"`
		}

		return c.Status(fiber.StatusCreated).JSON(UserResponse{
			Message:  "User registered successfully",
			Username: newUser.Username,
			Name:     newUser.Name,
			Email:    newUser.Email,
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
//	@Success		200		{object}	models.User	"User authenticated"
//	@Failure		400		{object}	string		"Invalid request payload"
//	@Failure		401		{object}	string		"User not found / Password doesn't match"
//	@Router			/api/user/login [post]
func Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type body struct {
			Username string
			Password string
		}
		b := new(body)
		if err := json.Unmarshal(c.Body(), &b); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		if b.Username == "" || b.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}

		var user models.User
		database.DB.Where("username=?", b.Username).First(&user)
		if len(user.Username) == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Username not found"})
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(b.Password))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Password"})
		}

		token, err := GenerateToken(user)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token creation error"})
		}

		type UserResponse struct {
			Message  string `json:"message"`
			Token    string `json:"token"`
			Username string `json:"username"`
			Name     string `json:"name"`
			Email    string `json:"email"`
		}
		return c.Status(fiber.StatusOK).JSON(UserResponse{
			Message:  "User Authenticated",
			Token:    token,
			Username: user.Username,
			Name:     user.Name,
			Email:    user.Email,
		})
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
//	@Security		ApiKeyAuth
//	@Param			token	header		string		true	"API Key"
//
//	@Param			user	body		models.User	true	"Update Password Request"
//	@Success		200		{object}	string		"Password updated successfully"
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

		type body struct {
			Username string
			Password string
		}
		b := new(body)
		if err := json.Unmarshal(c.Body(), &b); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		if b.Username == "" || b.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}

		if username != b.Username {
			return fiber.ErrUnauthorized
		}

		var existingUser models.User
		database.DB.Where("username=?", b.Username).First(&existingUser)
		if len(existingUser.Username) == 0 {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Username doesn't exists"})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "hashing failed"})
		}

		b.Password = string(hashedPassword)
		database.DB.Model(&existingUser).Updates(b)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Password updated successfully",
		})
	}
}

// DeleteUser handles deleting user account
//
//	@Summary		Delete user account
//	@Description	Deletes the account of the authenticated user
//	@Tags			User Management
//	@Produce		json
//
//	@Security		ApiKeyAuth
//	@Param			token	header		string		true	"API Key"
//
//	@Param			user	body		models.User	true	"User deletion request"
//	@Success		200		{object}	string		"User deleted successfully"
//	@Failure		400		{object}	string		"Invalid request payload"
//	@Failure		401		{object}	string		"Unauthorized"
//	@Failure		404		{object}	string		"Username doesn't exist"
//	@Failure		500		{object}	string		"Internal Server Error"
//	@Router			/api/v2/user [delete]
func DeleteUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		username, ok := c.Locals("username").(string)
		if !ok {
			return fiber.ErrUnauthorized
		}
		type body struct {
			Username string
			Password string
		}
		b := new(body)
		if err := json.Unmarshal(c.Body(), &b); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		if b.Username == "" || b.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}

		if username != b.Username {
			return fiber.ErrUnauthorized
		}

		var existingUser models.User
		database.DB.Where("username=?", b.Username).First(&existingUser)
		if len(existingUser.Username) == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Username doesn't exists"})
		}

		err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(b.Password))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Password"})
		}
		deleteInTaskAssignment(b.Username)
		database.DB.Delete(&existingUser)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User deleted successfully",
		})
	}
}
func deleteInTaskAssignment(username string) {
	taskAssignment := new(models.TaskAssignment)
	database.DB.Where("username=?", username).Delete(&taskAssignment)
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
//	@Security		ApiKeyAuth
//	@Param			token	header	string	true	"API Key"
//
//	@Produce		json
//	@Param			user	body		models.User	true	"Refresh user token"
//	@Success		200		{object}	models.User	"Token refreshed successfully"
//	@Failure		401		{object}	string		"Unauthorized"
//	@Router			/api/v2/refreshToken [get]
func RefreshToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type body struct {
			Username string
		}
		bodyUser := new(body)
		if err := json.Unmarshal(c.Body(), &bodyUser); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		if bodyUser.Username == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}

		tokenUsername := c.Locals("username")

		if bodyUser.Username != tokenUsername {
			log.Println("username and token username mismatch.")

			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "username and token username mismatch."})
		}

		var user models.User
		database.DB.Where("username=?", bodyUser.Username).First(&user)
		if len(user.Username) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "username not found."})
		}

		token, err := GenerateToken(user)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token creation error."})
		}

		type UserResponse struct {
			Token    string `json:"token"`
			Username string `json:"username"`
			Name     string `json:"name"`
			Email    string `json:"email"`
		}

		return c.Status(fiber.StatusOK).JSON(UserResponse{
			Token:    token,
			Username: user.Username,
			Name:     user.Name,
			Email:    user.Email,
		})
	}
}

// DisplayAllUsers handles retrieving all users
//
//	@Summary		Get all users
//	@Description	Retrieve all users
//	@Tags			User Management
//	@Accept			json
//	@Produce		json
//
//	@Security		ApiKeyAuth
//	@Param			token	header		string		true	"API Key"
//
//	@Success		200		{object}	models.Task	"User retrieved successfully"
//	@Router			/api/v2/alluser [get]
func DisplayAllUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user []models.User
		database.DB.Select("username, name, email").Find(&user)
		return c.Status(fiber.StatusOK).JSON(user)
	}
}
