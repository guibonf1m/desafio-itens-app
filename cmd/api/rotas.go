package main // se ele estiver sendo usado só no main.go, o pacote deve ser main

import (
	"desafio-itens-app/internal/adapters/http/handler"
	"desafio-itens-app/internal/adapters/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegistrarRotas(itemHandler *handler.ItemHandler, userHandler *handler.UserHandler, authMiddleware *middlewares.AuthMiddleware) *gin.Engine {
	router := gin.Default()

	// 🌍 ROTAS PÚBLICAS (sem autenticação)
	public := router.Group("v1")
	{
		public.POST("/register", userHandler.Register)
		public.POST("/login", userHandler.Login)
	}

	// 🔐 ROTAS PARA USUÁRIOS LOGADOS (qualquer role)
	authenticated := router.Group("v1")
	authenticated.Use(authMiddleware.RequireAuth()) // ← 1º segurança
	{
		// Qualquer usuário logado pode VER itens
		authenticated.GET("/itens", itemHandler.GetItens)
		authenticated.GET("/itens/:id", itemHandler.GetItem)
	}

	// 👤 ROTAS PARA USUÁRIOS (user ou admin)
	userRoutes := router.Group("v1")
	userRoutes.Use(authMiddleware.RequireAuth())                // ← 1º segurança
	userRoutes.Use(authMiddleware.RequireRole("user", "admin")) // ← 2º segurança
	{
		userRoutes.POST("/itens", itemHandler.AddItem)       // Criar item
		userRoutes.PUT("/itens/:id", itemHandler.UpdateItem) // Editar item
	}

	// 👑 ROTAS SÓ PARA ADMIN
	adminRoutes := router.Group("v1")
	adminRoutes.Use(authMiddleware.RequireAuth())        // ← 1º segurança
	adminRoutes.Use(authMiddleware.RequireRole("admin")) // ← 2º segurança
	{
		adminRoutes.DELETE("/itens/:id", itemHandler.DeleteItem) // Só admin deleta
		adminRoutes.GET("/users", userHandler.ListUsers)         // Gerenciar usuários
		adminRoutes.POST("/users", userHandler.CreateUser)       // Criar usuários
	}

	return router
}
