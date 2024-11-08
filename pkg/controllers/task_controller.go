package controllers

import (
	"strconv"
	"time"
	"todo-api/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	var Task models.Task
	c.BodyParser(&Task)

	dueDate, err := time.Parse("2006-01-02 15:04:05", Task.DueDate.Format("2006-01-02 15:04:05"))
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid due date format",
		})
		return err
	}
	Task.DueDate = dueDate

	Task.CreatedAt = time.Now()
	if err := models.PostTask(&Task); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create task",
		})
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Task created successfully",
		"task": fiber.Map{
			"title":       Task.Title,
			"description": Task.Description,
			"status":      Task.Status,
			"due_date":    Task.DueDate.Format("2006-01-02"),
		},
	})
}

func GetAllTasks(c *fiber.Ctx) error {
	tasks, err := models.SelectALLTasks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch tasks"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully fetched tasks",
		"tasks":   tasks,
	})
}

func GetTaskByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	task, _ := models.SelectTaskById(id)
	if task == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Task not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var Task models.Task
	c.BodyParser(&Task)
	Task.CreatedAt = time.Now()
	if err := models.UpdateTask(id, &Task); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update task",
		})
		return err
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task updated successfully",
		"task": fiber.Map{
			"title":       Task.Title,
			"description": Task.Description,
			"status":      Task.Status,
			"due_date":    Task.DueDate.Format("2006-01-02"),
		},
	})
}

func DeleteTask(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	models.DeleteTask(id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task deleted successfully",
	})
}
