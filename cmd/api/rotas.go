package main // se ele estiver sendo usado só no main.go, o pacote deve ser main

import (
	"desafio-itens-app/internal/adapters/http/handler"
	"desafio-itens-app/internal/adapters/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegistrarRotas(itemHandler *handler.ItemHandler, userHandler *handler.UserHandler, authMiddleware *middlewares.AuthMiddleware) *gin.Engine {
	router := gin.Default()

	public := router.Group("v1")
	{
		public.POST("/register", userHandler.Register)
		public.POST("/login", userHandler.Login)
	}

	protected := router.Group("v1")
	protected.Use(authMiddleware.RequireAuth())
	{

		itens := protected.Group("/itens")
		{
			itens.POST("", itemHandler.AddItem)
			itens.GET("/:id", itemHandler.GetItem)
			itens.GET("", itemHandler.GetItens)
			itens.PUT("/:id", itemHandler.UpdateItem)
			itens.DELETE("/:id", itemHandler.DeleteItem)
		}

		users := protected.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/: id", userHandler.DeleteUser)
			users.GET("/username/:username", userHandler.GetUserByUsername)
		}
	}

	return router
}
