package service

import (
	"robin-task/internal/model"
	"robin-task/internal/repository"
)

type ChangeLogService struct {
	changelogRepo repository.ChangeLogRepository
}

func NewChangeLogService(changelogRepo repository.ChangeLogRepository) *ChangeLogService {
	return &ChangeLogService{changelogRepo: changelogRepo}
}

func (s *ChangeLogService) GetChangeLogsByTaskID(taskID uint) ([]model.Changelog, error) {
	return s.changelogRepo.GetChangeLogsByTaskID(taskID)
}

func (s *ChangeLogService) CreateChangeLog(taskID uint, oldTask, newTask *model.Task, userID uint, action string) error {
	return s.changelogRepo.CreateChangeLog(taskID, oldTask, newTask, userID, action)
}
