desafio-itens-app
Este projeto é uma API simples para gestão de itens (produtos), desenvolvida em Go, utilizando o framework Gin Gonic para as rotas HTTP e MySQL como banco de dados.

🚀 Funcionalidades
Listagem de itens com paginação básica (parâmetro limit).

Estrutura de projeto modular e organizada.

Tratamento de erros robusto para requisições HTTP.

Integração com Docker para facilitar o ambiente de desenvolvimento.

📦 Tecnologias Utilizadas
Go Lang: Linguagem de programação principal.

Gin Gonic: Um framework web rápido para Go.

MySQL: Banco de dados relacional para persistência dos dados dos itens.

Docker / Docker Compose: Para orquestração e gerenciamento dos contêineres da aplicação e do banco de dados.

🛠️ Começando
Siga as instruções abaixo para configurar e executar o projeto em sua máquina local.

Pré-requisitos
Certifique-se de ter as seguintes ferramentas instaladas:

Go (versão 1.20 ou superior)

Docker

Docker Compose

Instalação
Clone o repositório:

git clone https://github.com/guibonf1m/desafio-itens-app.git
cd desafio-itens-app

Baixe as dependências do Go:

go mod tidy

Configuração do Banco de Dados (MySQL)
As configurações do banco de dados (usuário, senha, nome do DB) podem ser definidas no arquivo docker-compose.yml e/ou via variáveis de ambiente.

Exemplo de variáveis de ambiente esperadas (você pode criar um arquivo .env na raiz do projeto):

DB_HOST=mysql
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=desafio_itens_db

Executando a Aplicação
A maneira mais fácil de executar a aplicação junto com o banco de dados é usando Docker Compose:

docker-compose up --build

Este comando irá:

Construir a imagem Docker da sua aplicação Go.

Subir um contêiner MySQL.

Subir o contêiner da sua aplicação e conectá-lo ao MySQL.

A API estará disponível em http://localhost:8080.

Executando apenas a aplicação Go (sem Docker Compose para a aplicação)
Se você preferir executar apenas a aplicação Go localmente, garantindo que o MySQL esteja rodando (via Docker Compose ou localmente):

Certifique-se de que o MySQL esteja acessível:

Se estiver usando o MySQL via Docker Compose, comente o serviço app no docker-compose.yml e execute docker-compose up -d mysql.

Ou certifique-se de ter uma instância MySQL rodando localmente com as configurações correspondentes às suas variáveis de ambiente.

Execute a aplicação Go:

go run cmd/api/main.go

🗺️ Endpoints da API
GET /items
Retorna uma lista de itens.

Parâmetros de Query
limit (opcional): Um inteiro que especifica o número máximo de itens a serem retornados. Ex: ?limit=5.

Exemplo de Resposta
[
  {
    "id": "1",
    "name": "Item A",
    "description": "Descrição do Item A",
    "price": 10.50
  },
  {
    "id": "2",
    "name": "Item B",
    "description": "Descrição do Item B",
    "price": 20.00
  }
]

🏗️ Estrutura do Projeto
O projeto segue uma estrutura modular para separar as preocupações:

desafio-itens-app/
├── cmd/
│   └── api/             # Ponto de entrada principal da aplicação (main.go)
├── internal/
│   ├── adapters/
│   │   ├── http/        # Handlers HTTP (Gin Gonic)
│   │   └── repository/  # Implementações de repositório (ex: MySQL)
│   ├── application/
│   │   └── service/     # Camada de serviços e lógica de negócio
│   └── domain/
│       └── item/        # Definições da entidade Item
├── docker-compose.yml   # Configuração do Docker Compose
├── go.mod               # Módulos Go e dependências
├── go.sum               # Checksums das dependências
├── .gitignore           # Arquivos e diretórios a serem ignorados pelo Git
└── README.md            # Este arquivo

🤝 Contribuição
Contribuições são bem-vindas! Se você deseja contribuir, por favor:

Faça um fork do projeto.

Crie uma branch para sua feature (git checkout -b feature/minha-feature).

Faça suas alterações e commit (git commit -m 'feat: adiciona nova feature').

Envie para a branch (git push origin feature/minha-feature).

Abra um Pull Request.

📄 Licença
Este projeto está licenciado sob a Licença MIT. Veja o arquivo LICENSE para mais detalhes.
