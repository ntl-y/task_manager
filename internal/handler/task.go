package handler

import (
	"net/http"
	taskmanager "task_manager"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTask(c *gin.Context) {
	start := time.Now()
	defer func() {
		observeRequest(time.Since(start), c.Writer.Status())
	}()
	var input taskmanager.Task

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if input.Title == "" || input.Describtion == "" {
		newErrorResponse(c, http.StatusUnprocessableEntity, "empty title or describtion")
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
