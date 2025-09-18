package models

import "time"

type Comment struct {
	CommentID  int       `json:"comment_id"`
	PostID     int       `json:"post_id"`
	AuthorName string    `json:"author_name"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}
