package models

import "time"

type Post struct {
	PostID      int       `gorm:"primaryKey;autoIncrement" json:"post_id"`
	Title       string    `gorm:"column:title" json:"title"`
	Content     string    `gorm:"column:content" json:"content"`
	PublishedAt time.Time `gorm:"column:published_at" json:"published_at"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	Tags        []Tag     `gorm:"many2many:post_tags;" json:"tags,omitempty"`
}
