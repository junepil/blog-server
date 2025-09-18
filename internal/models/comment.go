package models

import "time"

type Comment struct {
	CommentID  int       `gorm:"primaryKey;autoIncrement" json:"comment_id"`
	PostID     int       `gorm:"column:post_id" json:"post_id"`
	AuthorName string    `gorm:"column:author_name" json:"author_name"`
	Content    string    `gorm:"column:content" json:"content"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
}
