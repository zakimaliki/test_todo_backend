package controllers

import (
	"log"
	"todo-api/pkg/config"
	"todo-api/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	query := `INSERT INTO tasks (title, description, status, due_date) VALUES (:1, :2, :3, :4) RETURNING id INTO :5`
	_, err := config.DB.Exec(query, task.Title, task.Description, task.Status, task.DueDate, &task.ID)
	if err != nil {
		log.Println("Error inserting task:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create task"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Task created successfully", "task": task})
}

func GetAllTasks(c *fiber.Ctx) error {
	rows, err := config.DB.Query("SELECT id, title, description, status, due_date, created_at, updated_at FROM tasks")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve tasks"})
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			log.Println("Error scanning task:", err)
			continue
		}
		tasks = append(tasks, task)
	}

	return c.JSON(fiber.Map{"tasks": tasks})
}

func GetTaskByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task
	query := "SELECT id, title, description, status, due_date, created_at, updated_at FROM tasks WHERE id = :1"
	err := config.DB.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}

	return c.JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	query := `UPDATE tasks SET title = :1, description = :2, status = :3, due_date = :4 WHERE id = :5`
	_, err := config.DB.Exec(query, task.Title, task.Description, task.Status, task.DueDate, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not update task"})
	}

	return c.JSON(fiber.Map{"message": "Task updated successfully"})
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "DELETE FROM tasks WHERE id = :1"
	_, err := config.DB.Exec(query, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete task"})
	}

	return c.JSON(fiber.Map{"message": "Task deleted successfully"})
}
