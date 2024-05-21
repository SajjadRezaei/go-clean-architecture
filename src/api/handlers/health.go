package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "Working!")
	return
}

func (h *HealthHandler) HealthPost(c *gin.Context) {
	c.JSON(http.StatusOK, "Working! post method")
	return
}
