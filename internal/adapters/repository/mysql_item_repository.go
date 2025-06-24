package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"

	"desafio-itens-app/internal/domain"
)

type MySQLItemRepository struct {
	db *sqlx.DB
}

func NewMySQLItemRepository(db *sqlx.DB) *MySQLItemRepository {
	return &MySQLItemRepository{db: db}
}

func (r *MySQLItemRepository) AddItem(item domain.Item) (domain.Item, error) {
	result, err := r.db.Exec("INSERT INTO itens (nome, preco, estoque, status) VALUES (?, ?, ?, ?)",
		item.Nome, item.Preco, item.Estoque, item.Status)
	if err != nil {
		return domain.Item{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Item{}, err
	}

	item.ID = int(id)
	return item, nil
}

func (r *MySQLItemRepository) GetItem(id int) (*domain.Item, error) {

	//Antes de qualquer operação, verificamos se o ID está vazio.
	if id == 0 {
		return nil, errors.New("Id não pode ser 0.")
	}

	// Esta consulta SQL busca um item pelo seu ID.
	// O '?' é um espaço reservado para o ID do item, ajudando a prevenir ataques de injeção SQL
	// e permitindo que o valor seja passado de forma segura.
	// Ela seleciona as colunas 'id', 'code', 'name' e 'description' da tabela 'items'.
	query := "SELECT id, code, name, description FROM items WHERE id = ?"

	// Executa a consulta preparada, passando o 'id' real.
	// O valor de 'id' substitui o '?' na consulta, garantindo que apenas o item específico seja retornado.
	row := r.db.QueryRow(query, id)

	var item domain.Item
	err := row.Scan(&item.ID, &item.Code, &item.Nome, &item.Descricao)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Item não encontrado com id %s", id)
		}
		return nil, err
	}
	
	return &item, nil
}

func (r *MySQLItemRepository) GetItens(item domain.Item) (domain.Item, error) {

}
func (r *MySQLItemRepository) UpdateItem(id int) (domain.Item, error) {

}
func (r *MySQLItemRepository) DeleteItem(id int) (domain.Item, error) {

}
