package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"robin-task/internal/model"
	"robin-task/internal/service"
	"robin-task/utils"
	"slices"
)

type TaskHandler struct {
	service          *service.TaskService
	changeLogService *service.ChangeLogService
}

func NewTaskHandler(s *service.TaskService, clSrv *service.ChangeLogService) *TaskHandler {
	return &TaskHandler{
		service:          s,
		changeLogService: clSrv,
	}
}

func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseData(false, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	task, err := h.service.GetTaskByID(utils.ParseStringToUint(c.Param("id"))) // Pass as uint
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseData(false, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var request struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseData(false, "Invalid input data", nil))
		return
	}

	//validate
	userID := utils.GetUserIDFromContext(c)
	if userID == 0 {
		c.JSON(http.StatusBadRequest, utils.ResponseData(false, "UserID are required", nil))
		return
	}

	if !slices.Contains(model.GetAllStatus(), request.Status) {
		c.JSON(http.StatusBadRequest, utils.ResponseData(false, "invalid status", nil))
		return
	}

	newTask := model.Task{
		Name:        request.Name,
		Description: request.Description,
		Status:      request.Status,
		UserID:      userID,
		Archived:    false,
	}

	err := h.service.CreateTask(&newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseData(false, err.Error(), nil))
		return
	}

	// Log the change
	go func() {
		err := h.changeLogService.CreateChangeLog(newTask.ID, &newTask, nil, userID, model.ACTION_CREATE)
		if err != nil {
			fmt.Println("create change log err:", err)
		}
	}()

	c.JSON(http.StatusOK, utils.ResponseData(true, "ok", newTask))
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {

	var request struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseData(false, "Invalid input data", nil))
		return
	}

	//find task
	task, err := h.service.GetTaskByID(utils.ParseStringToUint(c.Param("id")))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ResponseData(false, err.Error(), nil))
		return
	}

	//validate
	userID := utils.GetUserIDFromContext(c)
	if userID == 0 {
		c.JSON(http.StatusBadRequest, utils.ResponseData(false, "UserID are required", nil))
		return
	}

	if !slices.Contains(model.GetAllStatus(), request.Status) {
		c.JSON(http.StatusBadRequest, utils.ResponseData(false, "invalid status", nil))
		return
	}

	// Store the old task
	oldTask := *task

	//set new task
	task.Name = request.Name
	task.Description = request.Description
	task.Status = request.Status
	err = h.service.UpdateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseData(false, err.Error(), nil))
		return
	}

	// Log the change
	go func() {
		err := h.changeLogService.CreateChangeLog(task.ID, &oldTask, task, userID, model.ACTION_UPDATE)
		if err != nil {
			fmt.Println("create change log err:", err)
		}
	}()

	c.JSON(http.StatusOK, utils.ResponseData(true, "ok", task))
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {

	userID := utils.GetUserIDFromContext(c)
	if userID == 0 {
		c.JSON(http.StatusBadRequest, utils.ResponseData(false, "UserID are required", nil))
		return
	}

	//find task
	task, err := h.service.GetTaskByID(utils.ParseStringToUint(c.Param("id")))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ResponseData(false, err.Error(), nil))
		return
	}

	err = h.service.DeleteTask(task.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseData(false, err.Error(), nil))
		return
	}

	// Log the change
	go func() {
		err := h.changeLogService.CreateChangeLog(task.ID, task, nil, userID, model.ACTION_DELETE)
		if err != nil {
			fmt.Println("create change log err:", err)
		}
	}()

	c.JSON(http.StatusOK, utils.ResponseData(true, "ok", nil))
}

func (h *TaskHandler) ArchiveTask(c *gin.Context) {
	err := h.service.ArchiveTask(utils.ParseStringToUint(c.Param("id")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseData(false, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseData(true, "ok", nil))
}
