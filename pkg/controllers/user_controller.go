package controllers

import (
	"os"
	"todo-api/pkg/models"
	"todo-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPassword)

	models.PostUser(&user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"user": fiber.Map{
			"ID":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func LoginUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	validateEmail := models.FindEmail(&user)
	if len(validateEmail) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Email not found",
		})
	}

	var PasswordSecond string
	for _, user := range validateEmail {
		PasswordSecond = user.Password
	}

	if err := bcrypt.CompareHashAndPassword([]byte(PasswordSecond), []byte(user.Password)); err != nil || user.Password == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	jwtKey := os.Getenv("SECRETKEY")
	payload := map[string]interface{}{
		"email": user.Email,
	}

	token, err := utils.GenerateToken(jwtKey, payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate access token",
		})
	}

	refreshToken, err := utils.GenerateRefreshToken(jwtKey, payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate refresh token",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Message":       "Login successfully",
		"access_token":  token,
		"refresh_token": refreshToken,
	})
}

// func RefreshToken(c *fiber.Ctx) error {

// }
