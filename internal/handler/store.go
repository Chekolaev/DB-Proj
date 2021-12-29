package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ShowStore(ctx *gin.Context) {
	resp, ok := ctx.Get("userUUID")

	if !ok {
		log.Println("No user with this uuid!")
		return
	}

	respNew, err := h.services.GetUserByInterface(resp)

	if err != nil {
		log.Println(err)
		return
	}

	if respNew["role"] != "1" && respNew["role"] != "2" {
		log.Println("chet ne to!")
		return
	}

	resp, err = h.services.GetAllBooks()

	if err != nil {
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) ShowBook(ctx *gin.Context) {
	resp, ok := ctx.Get("userUUID")

	if !ok {
		log.Println("No user with this uuid!")
		return
	}

	respNew, err := h.services.GetUserByInterface(resp)

	if err != nil {
		log.Println(err)
		return
	}

	if respNew["role"] != "1" && respNew["role"] != "2" {
		log.Println("chet ne to!")
		return
	}

	uuidBook := ctx.Param("book_id")

	resp, err = h.services.GetBookByUUID(uuidBook)

	if err != nil {
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) RentBook(ctx *gin.Context) {
	resp, ok := ctx.Get("userUUID")

	if !ok {
		log.Println("No user with this uuid!")
		return
	}

	respNew, err := h.services.GetUserByInterface(resp)

	if err != nil {
		log.Println(err)
		return
	}

	if respNew["role"] != "1" && respNew["role"] != "2" {
		log.Println("Store no permission!")
		return
	}

	uuidBook := ctx.Param("book_id")

	resp, err = h.services.RentBookByUUID(uuidBook, resp)

	if err != nil {
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
