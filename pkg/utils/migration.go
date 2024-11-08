package utils

import (
	"todo-api/pkg/config"
	"todo-api/pkg/models"
)

func Migration() {
	config.DB.AutoMigrate(&models.Task{}, &models.User{})
}
