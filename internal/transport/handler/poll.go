package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"polling-system/internal/service"
)

type CreatePollRequest struct {
	Name string `json:"name"`
}

type CreatePollResponse struct {
	UUID string `json:"uuid"`
}

func (h *Handler) CreatePoll(ctx *gin.Context) {
	uuid, err := h.poll.Create(ctx.Request.Context(), &service.PollInfo{
		Name: "",
	})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{"data": uuid})
}
