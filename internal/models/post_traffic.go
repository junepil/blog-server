package models

import "time"

// PostTraffic represents a record of a view on a post.
type PostTraffic struct {
	ViewID    int       `json:"view_id"`
	PostID    int       `json:"post_id"`
	ViewedAt  time.Time `json:"viewed_at"`
	IPAddress string    `json:"ip_address"`
}
