package service

import (
	"robin-task/internal/model"
	"robin-task/internal/repository"
)

type TaskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) GetAllTasks() ([]model.Task, error) {
	return s.taskRepo.GetAllTasks()
}

func (s *TaskService) GetAllTaskByRole(userRole string) ([]model.Task, error) {
	if userRole == model.ROLE_ADMIN {
		return s.taskRepo.GetAllTasks()
	}

	return s.taskRepo.GetTasksWithConditions(
		map[string]interface{}{
			"archived": false,
		},
	)
}

func (s *TaskService) GetTaskByID(id uint) (*model.Task, error) {
	return s.taskRepo.GetTaskByID(id)
}

func (s *TaskService) GetTaskByIDAndRole(id uint, userRole string) (*model.Task, error) {
	if userRole == model.ROLE_ADMIN {
		return s.taskRepo.GetTaskByID(id)
	}

	return s.taskRepo.GetTaskByIDAndCondition(
		id,
		map[string]interface{}{
			"archived": false,
		},
	)
}

func (s *TaskService) CreateTask(task *model.Task) error {
	return s.taskRepo.CreateTask(task)
}

func (s *TaskService) UpdateTask(task *model.Task) error {
	return s.taskRepo.UpdateTask(task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.taskRepo.DeleteTask(id)
}

func (s *TaskService) ArchiveTask(id uint) error {
	task, err := s.taskRepo.GetTaskByID(id)
	if err != nil {
		return err
	}

	// Archive the task
	task.Archived = true
	return s.taskRepo.UpdateTask(task)
}
