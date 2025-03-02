package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"robin-task/internal/model"
	"robin-task/internal/service"
	"robin-task/utils"
	"slices"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseData(false, "Invalid input", nil))
		return
	}

	//validate
	if !slices.Contains(model.GetAllUserRoles(), input.Role) {
		c.JSON(http.StatusBadRequest, utils.ResponseData(false, "invalid role", nil))
		return
	}

	token, err := h.service.Register(input.Username, input.Password, input.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseData(false, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseData(true, "ok", map[string]string{"token": token}))
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseData(false, "Invalid input", nil))
		return
	}

	token, err := h.service.Login(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.ResponseData(false, "Invalid credentials", nil))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseData(true, "ok", map[string]string{"token": token}))
}
