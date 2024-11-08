package models

import (
	"time"
	"todo-api/pkg/config"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func PostUser(item *User) error {
	result := config.DB.Create(&item)
	return result.Error
}

func FindEmail(input *User) []User {
	items := []User{}
	config.DB.Raw("SELECT * FROM users WHERE email = ?", input.Email).Scan(&items)
	return items
}

func FindID(id uint) User {
	items := User{}
	config.DB.Raw("SELECT * FROM users WHERE id = ?", id).Scan(&items)
	return items
}
