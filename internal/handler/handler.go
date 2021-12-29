package handler

import (
	"Library/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	auth := router.Group("/authentificate")
	{
		auth.POST("/singUp", h.singUp)
		auth.POST("/singIn", h.singIn)
	}

	store := router.Group("/storage", h.userIdentify)
	{
		store.GET("/", h.ShowStore)
		store.GET("/:book_id", h.ShowBook)
		store.POST("/:book_id", h.RentBook)
	}

	admin := router.Group("/admin", h.userIdentify)
	{
		admin.GET("/requests", h.RequestGet)
		admin.POST("/requests", h.RequestPost)
	}

	return router
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}
