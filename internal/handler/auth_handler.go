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
		Avatar   string `json:"avatar"`
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

	token, err := h.service.Register(input.Username, input.Password, input.Role, input.Avatar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseData(false, err.Error(), nil))
		return
	}

	user, err := h.service.GetUserByUsername(input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseData(false, err.Error(), nil))
	}

	var userRes *model.UserResponse
	err = utils.CastType(user, &userRes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseData(false, err.Error(), nil))
	}

	c.JSON(http.StatusOK, utils.ResponseData(true, "ok", map[string]interface{}{"token": token, "user": userRes}))
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

	user, token, err := h.service.Login(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.ResponseData(false, "Invalid credentials", nil))
		return
	}

	var userRes *model.UserResponse
	err = utils.CastType(user, &userRes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseData(false, err.Error(), nil))
	}

	c.JSON(http.StatusOK, utils.ResponseData(true, "ok", map[string]interface{}{"token": token, "user": userRes}))
}
