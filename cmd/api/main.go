package main

import (
	"desafio-itens-app/internal/adapters/http"
	"desafio-itens-app/internal/adapters/mysql"
	"desafio-itens-app/internal/application"

	"log"
)

func main() {
	db, err := mysql.Conectar()
	if err != nil {
		log.Fatal("Erro ao conectar com o banco:", err)
	}
	defer db.Close()

	repo := mysql.NewMySQLItemRepository(db)

	service := application.NewItemService(repo)

	handler := http.NewItemHandler(service)

	router := RegistrarRotas(handler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erro ao subir o servidor:", err)
	}
}
