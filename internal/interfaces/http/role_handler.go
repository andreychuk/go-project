package http

import (
	"net/http"
	"strconv"

	"project/internal/domain/models"
	"project/internal/usecase"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	roleUsecase usecase.RoleUsecase
}

func NewRoleHandler(r *gin.Engine, rUsecase usecase.RoleUsecase) {
	handler := &RoleHandler{
		roleUsecase: rUsecase,
	}

	r.POST("/roles", handler.CreateRole)
	r.GET("/roles/:uid", handler.GetRoleByID)
	r.GET("/roles", handler.GetAllRoles)
	r.PUT("/roles/:uid", handler.UpdateRole)
	r.DELETE("/roles/:uid", handler.DeleteRole)
}

func (h *RoleHandler) CreateRole(c *gin.Context) {
	var role models.Role

	if err := c.BindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.roleUsecase.CreateRole(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, role)
}

func (h *RoleHandler) GetRoleByID(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	role, err := h.roleUsecase.GetRoleByID(uint(uid))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, role)
}

func (h *RoleHandler) GetAllRoles(c *gin.Context) {
	roles, err := h.roleUsecase.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

func (h *RoleHandler) UpdateRole(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	var role models.Role

	if err := c.BindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	role.UID = uint(uid)

	if err := h.roleUsecase.UpdateRole(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}

func (h *RoleHandler) DeleteRole(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	if err := h.roleUsecase.DeleteRole(uint(uid)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}
