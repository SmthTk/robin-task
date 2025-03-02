package repository

import (
	"gorm.io/gorm"
	"robin-task/internal/model"
)

type TaskRepository interface {
	GetAllTasks() ([]model.Task, error)
	GetTaskByID(id uint) (*model.Task, error)
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task) error
	DeleteTask(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

// PreloadUser is a helper function to preload the User fields (excluding Password)
func PreloadUser(db *gorm.DB) *gorm.DB {
	return db.Select([]string{"id", "username", "role"}) // Only select id, username, and role for the User
}

func PreloadComment(db *gorm.DB) *gorm.DB {
	return db.Select([]string{"id", "content", "created_at", "user_id", "task_id"}) // Select relevant fields for Comment
}

func (r *taskRepository) GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Preload("User", PreloadUser).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTaskByID(id uint) (*model.Task, error) {
	var task model.Task

	if err := r.db.
		Preload("User", PreloadUser).
		Preload("Comments", PreloadComment).
		Preload("Comments.User", PreloadUser).
		First(&task, id).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *taskRepository) CreateTask(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) UpdateTask(task *model.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) DeleteTask(id uint) error {
	return r.db.Delete(&model.Task{}, id).Error
}
