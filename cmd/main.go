package main

import (
	"log"

	"Library/internal/handler"
	"Library/internal/repository"
	"Library/internal/server"
	"Library/internal/service"

	_ "github.com/lib/pq"
)

func main() {

	//ConStr := "host=ip_your_db(127.0.0.1) port=port_your_db(5432) user=your_db_user(root) password=your_db_user_password(12345) dbname=your_db_name sslmode=disable"
	ConStr := "host=127.0.0.1 port=5432 user=postgres password=1369 dbname=postgres sslmode=disable"
	repos := repository.Init(ConStr)
	service := service.Init(repos)
	handlers := handler.NewHandler(service)

	var srv server.Server

	err := srv.Run("8080", handlers.InitRouter())

	if err != nil {
		log.Fatal(err)
	}
}
