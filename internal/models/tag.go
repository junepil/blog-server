package models

type Tag struct {
	TagID int    `gorm:"primaryKey;autoIncrement" json:"tag_id"`
	Name  string `gorm:"column:name;unique" json:"name"`
	Posts []Post `gorm:"many2many:post_tags;" json:"posts,omitempty"`
}
