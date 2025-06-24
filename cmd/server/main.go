package main

import (
	"desafio-itens-app/internal/adapters/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("", index)
	r.POST("/itens/:categoria", handler.ItemHandler.AddItem)
	r.GET("/item/:id", handler.ItemHandler.GetItem)
	r.GET("/itens", handler.ItemHandler.GetAllItens)
	r.PUT("/item/:id", handler.ItemHandler.UpdateItem)
	r.DELETE("/item/:id", handler.ItemHandler.DeleteItem)

	r.Run(":8080")

}
func index(c *gin.Context) {
	c.JSON(http.StatusOK, "Bem vindo a minha segunda API!")
}
