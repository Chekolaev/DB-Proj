package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookReq struct {
	UUID      string `json:"uuid"`
	NewStatus string `json:"status"`
}

func (h *Handler) RequestPost(ctx *gin.Context) {
	resp, ok := ctx.Get("userUUID")

	if !ok {
		log.Println("No users with this uuid")
		return
	}

	respNew, err := h.services.GetUserByInterface(resp)

	if err != nil {
		log.Println(err)
		return
	}

	if respNew["role"] != "2" {
		log.Println("No permission!")
		return
	}

	var input BookReq

	if err := ctx.BindJSON(&input); err != nil {
		log.Println("json wrong format:", err)
	}

	resp, err = h.services.ChangeRequest(input.UUID, input.NewStatus)

	if err != nil {
		log.Println(err)
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) RequestGet(ctx *gin.Context) {
	resp, ok := ctx.Get("userUUID")

	if !ok {
		log.Println("No users with this uuid")
		return
	}

	respNew, err := h.services.GetUserByInterface(resp)

	if err != nil {
		log.Println(err)
		return
	}

	if respNew["role"] != "2" {
		log.Println("No permission!")
		return
	}

	resp, err = h.services.GetRequests()

	if err != nil {
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
