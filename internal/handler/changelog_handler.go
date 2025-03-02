package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"robin-task/internal/service"
	"robin-task/utils"
)

type ChangeLogHandler struct {
	service *service.ChangeLogService
}

func NewChangeLogHandler(s *service.ChangeLogService) *ChangeLogHandler {
	return &ChangeLogHandler{service: s}
}

func (h *ChangeLogHandler) GetChangeLogsByTaskID(c *gin.Context) {
	changeLogs, err := h.service.GetChangeLogsByTaskID(utils.ParseStringToUint(c.Param("id")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseData(false, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseData(true, "ok", changeLogs))
}
