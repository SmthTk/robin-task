package model

import "time"

type Comment struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	TaskID    uint       `json:"task_id"`
	Content   string     `json:"content"`
	UserID    uint       `json:"user_id"`
	User      User       `json:"user" gorm:"foreignKey:user_id"`
	CreatedAt *time.Time `json:"created_at"`
}
