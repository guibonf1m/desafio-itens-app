package main // se ele estiver sendo usado só no main.go, o pacote deve ser main

import (
	"desafio-itens-app/internal/adapters/http"
	"github.com/gin-gonic/gin"
)

func RegistrarRotas(handler *http.ItemHandler) *gin.Engine {
	router := gin.Default()

	itens := router.Group("/itens")
	{
		itens.POST("", handler.AddItem)
		itens.GET("/:id", handler.GetItem)
	}

	return router
}
