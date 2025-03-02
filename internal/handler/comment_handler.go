package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"robin-task/internal/model"
	"robin-task/internal/service"
	"robin-task/utils"
)

type CommentHandler struct {
	service *service.CommentService
}

func NewCommentHandler(s *service.CommentService) *CommentHandler {
	return &CommentHandler{service: s}
}

func (h *CommentHandler) GetCommentsByTaskID(c *gin.Context) {
	comments, err := h.service.GetCommentsByTaskID(utils.ParseStringToUint(c.Param("id")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseData(false, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseData(true, "ok", comments))
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	var request struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseData(false, "Invalid input data", nil))
		return
	}

	userID := utils.GetUserIDFromContext(c)
	taskID := utils.ParseStringToUint(c.Param("id"))

	if taskID == 0 || userID == 0 {
		c.JSON(http.StatusBadRequest, utils.ResponseData(false, "taskID and userID are required", nil))
		return
	}

	newComment := model.Comment{
		TaskID:  taskID,
		Content: request.Content,
		UserID:  userID,
	}

	err := h.service.CreateComment(&newComment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseData(false, err.Error(), nil))
		return
	}

	c.JSON(http.StatusCreated, utils.ResponseData(true, "ok", newComment))
}

func (h *CommentHandler) UpdateComment(c *gin.Context) {
	var request struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseData(false, "Invalid input data", nil))
		return
	}

	commentID := utils.ParseStringToUint(c.Param("commentID"))
	userID := utils.GetUserIDFromContext(c)

	// Fetch the comment to validate ownership
	comment, err := h.service.GetCommentByID(commentID)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ResponseData(false, err.Error(), nil))
		return
	}

	// Check if the user is the owner of the comment
	if comment.UserID != userID {
		c.JSON(http.StatusForbidden, utils.ResponseData(false, "You are not allowed to update this comment", nil))
		return
	}

	comment.Content = request.Content
	err = h.service.UpdateComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseData(false, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseData(true, "ok", comment))
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	commentID := utils.ParseStringToUint(c.Param("commentID"))
	userID := utils.GetUserIDFromContext(c)

	// Fetch the comment to validate ownership
	comment, err := h.service.GetCommentByID(commentID)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ResponseData(false, err.Error(), nil))
		return
	}

	// Check if the user is the owner of the comment
	if comment.UserID != userID {
		c.JSON(http.StatusForbidden, utils.ResponseData(false, "You are not allowed to delete this comment", nil))
		return
	}

	err = h.service.DeleteComment(commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseData(false, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseData(true, "ok", nil))
}
