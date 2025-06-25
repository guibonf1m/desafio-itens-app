package main

import (
	"desafio-itens-app/internal/adapters/http"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func main() {
	r := gin.Default()

	db, err := sqlx.Open()

	r.GET("", index)
	r.POST("/itens/:categoria", http.ItemHandler.AddItem)
	r.GET("/item/:id", http.ItemHandler.GetItem)
	r.GET("/itens", http.ItemHandler.GetAllItens)
	r.PUT("/item/:id", http.ItemHandler.UpdateItem)
	r.DELETE("/item/:id", http.ItemHandler.DeleteItem)

	r.Run(":8080")

}
func index(c *gin.Context) {
	c.JSON(http.StatusOK, "Bem vindo a minha segunda API!")
}
