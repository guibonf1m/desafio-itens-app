# Desafio Itens App

## Sobre o Projeto
O **Desafio Itens App** é uma API REST desenvolvida em Go com o framework Gin, utilizando arquitetura hexagonal e banco de dados MySQL. O projeto visa gerenciar itens, permitindo operações de CRUD.

## Funcionalidades
- **Gerenciamento de Itens**: criar, ler, atualizar e excluir itens
- **Validações**: verificações de preço, estoque e status
- **Código Único**: geração automática de código único para cada item

## Tecnologias Utilizadas
- **Go**: linguagem de programação
- **Gin**: framework para desenvolvimento web
- **MySQL**: banco de dados
- **Docker**: containerização
- **Colima**: ambiente de desenvolvimento

## Como Executar o Projeto

1. **Clone o repositório:**
   ```sh
   git clone https://github.com/guibonf1m/desafio-itens-app.git
   ```

2. **Instale o [Colima](https://github.com/abiosoft/colima)** (caso ainda não tenha).
   ```sh
   brew install colima
   ```

3. **Inicie o Colima:**
   ```sh
   colima start
   ```

4. **Instale o [Docker](https://www.docker.com/get-started/)** (caso ainda não tenha).

5. **Suba um container MySQL com Docker:**
   ```sh
   docker run --name mysql-desafio-itens -e MYSQL_ROOT_PASSWORD=senha123 -e MYSQL_DATABASE=itens_db -p 3306:3306 -d mysql:8
   ```
   > **Obs:**  
   > - Usuário: `root`  
   > - Senha: `senha123`  
   > - Banco: `itens_db`  
   > Ajuste variáveis de ambiente conforme necessário.

6. **Atualize as configurações de conexão do seu projeto (`.env` ou arquivo de config, conforme implementado) para apontar para:**  
   ```
   host: localhost  
   user: root  
   password: senha123  
   dbname: itens_db  
   port: 3306
   ```

7. **Instale as dependências do projeto (se houver):**
   ```sh
   go mod tidy
   ```

8. **Execute o projeto:**
   ```sh
   go run main.go
   ```

---

## Contribuição

Contribuições são bem-vindas! Se você deseja contribuir, por favor, faça um fork do repositório e envie um pull request.

---

## Licença

O Desafio Itens App é licenciado sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.f1m/desafio-itens-app.git
