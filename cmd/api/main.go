package main

import (
	"desafio-itens-app/internal/adapters/http/auth"
	"desafio-itens-app/internal/adapters/http/handler"
	"desafio-itens-app/internal/adapters/http/middlewares"
	"desafio-itens-app/internal/adapters/mysql"
	"desafio-itens-app/internal/application/service"

	"log"
)

func main() {
	db, err := mysql.ConectarGORM()
	if err != nil {
		log.Fatal("Erro ao conectar com o banco:", err)
	}

	itemRepo := mysql.NewMySQLItemRepository(db)
	userRepo := mysql.NewMySQLUserRepository(db)

	itemService := service.NewItemService(itemRepo)
	userService := service.NewUserService(userRepo)

	jwtService := auth.NewJWTService("minha-secret-key-super-secreta")
	authMiddleware := middlewares.NewAuthMiddleware(jwtService)

	itemHandler := handler.NewItemHandler(itemService)
	userHandler := handler.NewUserHandler(userService, jwtService)

	router := RegistrarRotas(itemHandler, userHandler, authMiddleware)
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erro ao subir o servidor:", err)
	}
}
