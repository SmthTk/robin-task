package model

import "time"

const (
	TODO        = "todo"
	IN_PROGRESS = "in_progress"
	DONE        = "done"
)

type Task struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	UserID      uint       `json:"user_id"`
	User        User       `json:"user" gorm:"foreignKey:user_id"`
	Archived    bool       `json:"archived"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Comments    []Comment  `json:"comments" gorm:"foreignKey:task_id"`
}

func GetAllStatus() []string {
	return []string{TODO, IN_PROGRESS, DONE}
}
