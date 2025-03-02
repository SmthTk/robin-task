package service

import (
	"robin-task/internal/model"
	"robin-task/internal/repository"
)

type CommentService struct {
	commentRepo repository.CommentRepository
}

func NewCommentService(commentRepo repository.CommentRepository) *CommentService {
	return &CommentService{commentRepo: commentRepo}
}

func (s *CommentService) GetCommentsByTaskID(taskID uint) ([]model.Comment, error) {
	return s.commentRepo.GetCommentsByTaskID(taskID)
}

func (s *CommentService) CreateComment(comment *model.Comment) error {
	return s.commentRepo.CreateComment(comment)
}

func (s *CommentService) GetCommentByID(commentID uint) (*model.Comment, error) {
	return s.commentRepo.GetCommentByID(commentID)
}

func (s *CommentService) UpdateComment(comment *model.Comment) error {
	return s.commentRepo.UpdateComment(comment)
}

func (s *CommentService) DeleteComment(commentID uint) error {
	return s.commentRepo.DeleteComment(commentID)
}
