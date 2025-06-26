package main

import (
	"desafio-itens-app/internal/adapters/http"
	"desafio-itens-app/internal/adapters/mysql" // <- Aqui está sua função Conectar
	"desafio-itens-app/internal/application"

	"log"
)

func main() {
	// 1. Conectar ao banco MySQL
	db, err := mysql.Conectar()
	if err != nil {
		log.Fatal("Erro ao conectar com o banco:", err)
	}
	defer db.Close()

	// 2. Criar o repositório
	repo := mysql.NewMySQLItemRepository(db)

	// 3. Criar o service
	service := application.NewItemService(repo)

	// 4. Criar o handler
	handler := http.NewItemHandler(service)

	// 5. Registrar as rotas
	router := RegistrarRotas(handler)

	// 6. Subir o servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erro ao subir o servidor:", err)
	}
}
