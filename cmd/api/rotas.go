package main // se ele estiver sendo usado só no main.go, o pacote deve ser main

import (
	"desafio-itens-app/internal/adapters/http"
	"github.com/gin-gonic/gin"
)

func RegistrarRotas(handler *http.ItemHandler) *gin.Engine {
	router := gin.Default()

	itens := router.Group("v1/itens")
	{
		itens.POST("", handler.AddItem)
		itens.GET("/:id", handler.GetItem)
		itens.GET("", handler.GetItens)
		itens.PUT("/:id", handler.UpdateItem)
		itens.DELETE("/:id", handler.DeleteItem)
	}

	return router
}
