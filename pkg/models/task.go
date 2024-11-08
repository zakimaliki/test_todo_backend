package models

import (
	"fmt"
	"time"
	"todo-api/pkg/config"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// func SelectALLTasks() *gorm.DB {
// 	// items := []Article{}
// 	// config.DB.Raw("SELECT * FROM articles").Scan(&items)
// 	// return items

// 	items := []Task{}
// 	return config.DB.Find(&items)

// }

//	func SelectALLTasks() ([]Task, error) {
//		items := []Task{}
//		result := config.DB.Find(&items)
//		return items, result.Error
//	}
func SelectALLTasks(sort, keyword, status string, limit, offset int) ([]Task, error) {
	var tasks []Task
	query := config.DB

	// Kondisi untuk title atau description
	if keyword != "%%" {
		query = query.Where("title LIKE ? OR description LIKE ?", keyword, keyword)
	}

	// Kondisi untuk status, default ke "pending" jika tidak diatur
	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Order(sort).
		Limit(limit).
		Offset(offset).
		Find(&tasks).Error
	return tasks, err
}

func SelectTaskById(id int) (*Task, error) {
	var item Task
	result := config.DB.First(&item, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

func PostTask(item *Task) error {
	fmt.Printf("Task to insert: %+v\n", item)
	result := config.DB.Create(&item)
	if result.Error != nil {
		fmt.Println("Error inserting task:", result.Error)
	}
	return result.Error
}

func UpdateTask(id int, newTask *Task) error {
	var item Task
	result := config.DB.Model(&item).Where("id = ?", id).Updates(newTask)
	return result.Error
}

func DeleteTask(id int) error {
	var item Task
	result := config.DB.Delete(&item, "id = ?", id)
	return result.Error
}

func CountData(keyword, status string) int64 {
	var count int64
	query := config.DB.Model(&Task{})

	// Kondisi untuk title atau description
	if keyword != "%%" {
		query = query.Where("title LIKE ? OR description LIKE ?", keyword, keyword)
	}

	// Kondisi untuk status, default ke "pending" jika tidak diatur
	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&count)
	return count
}
