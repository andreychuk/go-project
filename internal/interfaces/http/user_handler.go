package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"project/internal/domain/models"
	"project/internal/usecase"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(r *gin.Engine, uUsecase usecase.UserUsecase) {
	handler := &UserHandler{
		userUsecase: uUsecase,
	}

	r.POST("/users", handler.CreateUser)
	r.GET("/users/:uid", handler.GetUserByID)
	r.GET("/users", handler.GetAllUsers)
	r.PUT("/users/:uid", handler.UpdateUser)
	r.DELETE("/users/:uid", handler.DeleteUser)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Status   int    `json:"status"`
		RoleID   uint   `json:"role_id"`
		GroupIDs []uint `json:"group_ids"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user := models.User{
		Name:   input.Name,
		Email:  input.Email,
		Status: input.Status,
		RoleID: input.RoleID,
	}

	if err := h.userUsecase.CreateUser(&user, input.GroupIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userUsecase.GetUserByID(uint(uid))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userUsecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Status   int    `json:"status"`
		RoleID   uint   `json:"role_id"`
		GroupIDs []uint `json:"group_ids"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user := models.User{
		UID:    uint(uid),
		Name:   input.Name,
		Email:  input.Email,
		Status: input.Status,
		RoleID: input.RoleID,
	}

	if err := h.userUsecase.UpdateUser(&user, input.GroupIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.userUsecase.DeleteUser(uint(uid)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
