package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	UUID     string `json:"-"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

func (h *Handler) singUp(ctx *gin.Context) {
	var input User

	if err := ctx.BindJSON(&input); err != nil {
		log.Println("registration failed:", err)
		return
	}
	err := h.services.AddNewUser(input.Name, input.Surname, input.Password, input.Login)

	if err != nil {
		log.Panicln(err)
	}

	log.Print("Registration: 200")
}

func (h *Handler) singIn(ctx *gin.Context) {
	var input User

	if err := ctx.BindJSON(&input); err != nil {
		log.Println("registration failed:", err)
		return
	}
	token, err := h.services.GenerateJWT(input.Login, input.Password)

	if err != nil {
		log.Panicln(err)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
