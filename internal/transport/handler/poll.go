package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"polling-system/internal/service"
)

type CreatePollRequest struct {
	Name string `json:"name"`
}

type CreatePollResponse struct {
	UUID string `json:"uuid"`
}

func (h *Handler) CreatePoll(ctx *gin.Context) {
	var newPoll service.PollInfo
	if err := ctx.BindJSON(&newPoll); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newPoll.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	if len(newPoll.Options) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Options are required"})
		return
	}

	uuid, err := h.poll.Create(ctx.Request.Context(), &newPoll)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{"data": uuid})
}

func (h *Handler) GetPoll(ctx *gin.Context) {
	uuid := ctx.Query("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "UUID parameter is required"})
		return
	}

	pollInfo, err := h.poll.Get(ctx.Request.Context(), uuid)
	if err != nil {
		ctx.Error(fmt.Errorf("failed to get poll: %w", err))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"data": pollInfo})
}
