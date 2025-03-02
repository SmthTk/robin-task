package repository

import (
	"gorm.io/gorm"
	"robin-task/internal/model"
)

type CommentRepository interface {
	GetCommentsByTaskID(taskID uint) ([]model.Comment, error)
	CreateComment(comment *model.Comment) error
	GetCommentByID(commentID uint) (*model.Comment, error)
	UpdateComment(comment *model.Comment) error
	DeleteComment(commentID uint) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) GetCommentsByTaskID(taskID uint) ([]model.Comment, error) {
	var comments []model.Comment
	err := r.db.Where("task_id = ?", taskID).Find(&comments).Error
	return comments, err
}

func (r *commentRepository) CreateComment(comment *model.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) GetCommentByID(commentID uint) (*model.Comment, error) {
	var comment model.Comment
	err := r.db.First(&comment, commentID).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepository) UpdateComment(comment *model.Comment) error {
	return r.db.Save(comment).Error
}

func (r *commentRepository) DeleteComment(commentID uint) error {
	return r.db.Delete(&model.Comment{}, commentID).Error
}
