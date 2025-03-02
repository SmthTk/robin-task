package repository

import (
	"encoding/json"
	"gorm.io/gorm"
	"robin-task/internal/model"
	"time"
)

type ChangeLogRepository interface {
	GetChangeLogsByTaskID(taskID uint) ([]model.Changelog, error)
	CreateChangeLog(taskID uint, oldTask, newTask *model.Task, userID uint, action string) error
}

type changelogRepository struct {
	db *gorm.DB
}

func NewChangeLogRepository(db *gorm.DB) ChangeLogRepository {
	return &changelogRepository{db: db}
}

func (r *changelogRepository) GetChangeLogsByTaskID(taskID uint) ([]model.Changelog, error) {
	var logs []model.Changelog
	err := r.db.Where("task_id = ?", taskID).Find(&logs).Error
	return logs, err
}

func (r *changelogRepository) CreateChangeLog(taskID uint, oldTask, newTask *model.Task, userID uint, action string) error {

	var oldTaskJSON, newTaskJSON []byte
	var err error

	if oldTask != nil {
		oldTaskJSON, err = json.Marshal(oldTask)
		if err != nil {
			return err
		}
	} else {
		oldTaskJSON = []byte("")
	}

	if newTask != nil {
		newTaskJSON, err = json.Marshal(newTask)
		if err != nil {
			return err
		}
	} else {
		newTaskJSON = []byte("")
	}

	now := time.Now()
	logEntry := model.Changelog{
		TaskID:    taskID,
		OldValue:  string(oldTaskJSON),
		NewValue:  string(newTaskJSON),
		UserID:    userID,
		Action:    action,
		ChangedAt: &now,
	}

	return r.db.Create(&logEntry).Error
}
