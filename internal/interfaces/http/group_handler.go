package http

import (
	"net/http"
	"strconv"

	"project/internal/domain/models"
	"project/internal/usecase"

	"github.com/gin-gonic/gin"
)

type GroupHandler struct {
	groupUsecase usecase.GroupUsecase
}

func NewGroupHandler(r *gin.Engine, gUsecase usecase.GroupUsecase) {
	handler := &GroupHandler{
		groupUsecase: gUsecase,
	}

	r.POST("/groups", handler.CreateGroup)
	r.GET("/groups/:uid", handler.GetGroupByID)
	r.GET("/groups", handler.GetAllGroups)
	r.PUT("/groups/:uid", handler.UpdateGroup)
	r.DELETE("/groups/:uid", handler.DeleteGroup)
}

func (h *GroupHandler) CreateGroup(c *gin.Context) {
	var group models.Group

	if err := c.BindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.groupUsecase.CreateGroup(&group); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, group)
}

func (h *GroupHandler) GetGroupByID(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	group, err := h.groupUsecase.GetGroupByID(uint(uid))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
		return
	}

	c.JSON(http.StatusOK, group)
}

func (h *GroupHandler) GetAllGroups(c *gin.Context) {
	groups, err := h.groupUsecase.GetAllGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, groups)
}

func (h *GroupHandler) UpdateGroup(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	var group models.Group

	if err := c.BindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	group.UID = uint(uid)

	if err := h.groupUsecase.UpdateGroup(&group); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, group)
}

func (h *GroupHandler) DeleteGroup(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	if err := h.groupUsecase.DeleteGroup(uint(uid)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Group deleted successfully"})
}
