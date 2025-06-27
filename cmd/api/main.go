package main

import (
	"desafio-itens-app/internal/adapters/http"
	"desafio-itens-app/internal/adapters/mysql"
	"desafio-itens-app/internal/application"

	"log"
)

func main() {
	// 1. Conectar ao banco MySQL - > Todas as camadas dependem do banco
	db, err := mysql.Conectar()
	if err != nil {
		log.Fatal("Erro ao conectar com o banco:", err)
	}
	defer db.Close()

	// 2. Criar o repositório - Essa camada sabe como acessar o banco, aqui moram todos os métodos
	//INSERT, SELECT, UPDATE etc. Ele recebe o db, ou seja a conexão que eu abri no passo 1.
	repo := mysql.NewMySQLItemRepository(db)

	// 3. Criar o service - Camada onde fica as regras de negócio, recebo repo por que preciso
	//dos dados para aplicar as regras.
	service := application.NewItemService(repo)

	// 4. Criar o handler - Handler é quem recebe a requisição Postman,
	//ela faz a ponte entre o mundo HTTP (requisições) e o mundo GO (service)
	handler := http.NewItemHandler(service)

	// 5. Registrar as rotas - Rotas são os caminhos da API -> /itens, /itens/:id, etc.
	//Você centraliza a configuração das rotas num arquivo separado.
	router := RegistrarRotas(handler)

	// 6. Subir o servidor - Aqui você sobe sua API e ela começa a escutar na porta 8080.
	//Ou seja: se você for no Postman e fizer um request pra http://localhost:8080/itens,
	//a requisição vai chegar na sua aplicação.
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erro ao subir o servidor:", err)
	}
}
