package models

import "time"

type User struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string `gorm:"column:username;unique" json:"username"`
	Email     string `gorm:"column:email;unique" json:"email"`
	Role      string `gorm:"role;" json:"role"`
	Password  string `gorm:"column:password" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}