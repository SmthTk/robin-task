package model

import "time"

const (
	ACTION_CREATE = "create"
	ACTION_UPDATE = "update"
	ACTION_DELETE = "delete"
)

type Changelog struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	TaskID    uint       `json:"task_id"`
	OldValue  string     `json:"old_value"`
	NewValue  string     `json:"new_value"`
	UserID    uint       `json:"user_id"`
	User      User       `json:"user" gorm:"foreignKey:user_id"`
	ChangedAt *time.Time `json:"changed_at"`
	Action    string     `json:"action"`
}

func GetAllChangelogActions() []string {
	return []string{ACTION_CREATE, ACTION_UPDATE, ACTION_DELETE}
}

func (Changelog) TableName() string {
	return "changelogs"
}
