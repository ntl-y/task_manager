package handler

import (
	"net/http"
	taskmanager "task_manager"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTask(c *gin.Context) {
	var input taskmanager.Task

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var output taskmanager.TaskValidated
	output, err := h.service.Task.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, output)
}
