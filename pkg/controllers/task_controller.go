package controllers

import (
	"context"
	"encoding/json"
	"math"
	"strconv"
	"strings"
	"time"
	"todo-api/pkg/config"
	"todo-api/pkg/models"
	"todo-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	ctx := context.Background()
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

	// Simpan task ke database
	if err := models.PostTask(&Task); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create task",
		})
		return err
	}

	taskId := Task.ID
	// Simpan task ke Redis sebagai cache
	utils.CacheTask(ctx, taskId, &Task)

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
	pageOld := c.Query("page")
	limitOld := c.Query("limit")
	page, _ := strconv.Atoi(pageOld)
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(limitOld)
	if limit == 0 {
		limit = 5
	}
	offset := (page - 1) * limit
	sort := c.Query("sorting")
	if sort == "" {
		sort = "ASC"
	}
	sortby := c.Query("orderBy")
	if sortby == "" {
		sortby = "title"
	}
	sort = sortby + " " + strings.ToLower(sort)

	// Ambil nilai keyword dan status dari query
	keyword := c.Query("search")
	if keyword != "" {
		keyword = "%" + keyword + "%"
	} else {
		keyword = "%%"
	}

	// Set nilai default status ke "pending" jika tidak ditentukan
	status := c.Query("status")
	if status == "" {
		status = "pending"
	}

	// Ambil data tasks
	tasks, _ := models.SelectALLTasks(sort, keyword, status, limit, offset)
	totalData := models.CountData(keyword, status)
	totalPage := math.Ceil(float64(totalData) / float64(limit))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Tasks": tasks,
		"pagination": fiber.Map{
			"current_page": page,
			"total_pages":  int(totalPage),
			"total_tasks":  totalData,
		},
	})
}

func GetTaskByID(c *fiber.Ctx) error {
	ctx := context.Background()
	id := c.Params("id")
	cacheKey := "task:" + id

	// Coba ambil dari cache Redis
	cachedData, err := config.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var task *models.Task
		json.Unmarshal([]byte(cachedData), &task)
		return c.Status(fiber.StatusOK).JSON(task)
	}

	// Jika tidak ada di cache, ambil dari database
	taskId, _ := strconv.Atoi(id)
	task, _ := models.SelectTaskById(taskId)
	if task == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Task not found",
		})
	}

	// Cache data task di Redis
	utils.CacheTask(ctx, taskId, task)

	return c.Status(fiber.StatusOK).JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	ctx := context.Background()
	id, _ := strconv.Atoi(c.Params("id"))
	var Task models.Task
	c.BodyParser(&Task)
	Task.CreatedAt = time.Now()

	// Update task di database
	if err := models.UpdateTask(id, &Task); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update task",
		})
		return err
	}

	// Update cache di Redis
	utils.CacheTask(ctx, id, &Task)

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

	// Hapus task dari database
	models.DeleteTask(id)

	// Hapus task dari Redis
	utils.DeleteCacheTask(context.Background(), id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task deleted successfully",
	})
}
