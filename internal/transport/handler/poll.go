package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"polling-system/internal/service"
)

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

	uuid, err := h.poll.Create(&newPoll)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{"data": uuid})
}

func (h *Handler) GetPoll(ctx *gin.Context) {
	uuid := ctx.Query("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "uuid parameter is required"})
		return
	}

	pollInfo, err := h.poll.Get(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"data": pollInfo})
}

func (h *Handler) SaveVote(ctx *gin.Context) {
	uuid := ctx.Query("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Uuid parameter is required"})
		return
	}

	answer := ctx.Query("answer")
	if answer == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Answer parameter is required"})
		return
	}

	err := h.poll.SaveVote(uuid, answer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}
