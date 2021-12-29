package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) userIdentify(ctx *gin.Context) {
	header := ctx.GetHeader("autherisationHeader")

	if header == "" {
		log.Println("authentificate header empty")
		return
	}

	userUUID, err := h.services.ParseToken(header)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.Set("userUUID", userUUID)
}
