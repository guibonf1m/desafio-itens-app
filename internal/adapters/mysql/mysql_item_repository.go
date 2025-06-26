package mysql

import (
	"database/sql"
	entity "desafio-itens-app/internal/domain/item"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLItemRepository struct {
	db *sqlx.DB
}

func NewMySQLItemRepository(db *sqlx.DB) *MySQLItemRepository {
	return &MySQLItemRepository{db: db}
}

func Conectar() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/desafio_itens")
	if err != nil {
		return nil, err
	}
	return db, nil
}
func (r *MySQLItemRepository) AddItem(item entity.Item) (entity.Item, error) {
	result, err := r.db.Exec("INSERT INTO itens (code, nome, preco, descricao, estoque, status) VALUES (?, ?, ?, ?, ?, ?)",
		item.Code, item.Nome, item.Preco, item.Descricao, item.Estoque, item.Status)
	if err != nil {
		return entity.Item{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.Item{}, err
	}

	item.ID = int(id)
	return item, nil
}

func (r *MySQLItemRepository) GetItem(id int) (*entity.Item, error) {

	//Antes de qualquer operação, verificamos se o ID está vazio.
	if id == 0 {
		return nil, errors.New("Id não pode ser 0.")
	}

	// Esta consulta SQL busca um item pelo seu ID.
	// O '?' é um espaço reservado para o ID do item, ajudando a prevenir ataques de injeção SQL
	// e permitindo que o valor seja passado de forma segura.
	// Ela seleciona as colunas 'id', 'code', 'name' e 'description' da tabela 'items'.
	query := "SELECT * FROM itens WHERE id = ?"

	// Executa a consulta preparada, passando o 'id' real.
	// O valor de 'id' substitui o '?' na consulta, garantindo que apenas o item específico seja retornado.
	row := r.db.QueryRow(query, id)

	var item entity.Item
	err := row.Scan(&item.ID, &item.Code, &item.Nome, &item.Descricao, &item.Preco, &item.Estoque,
		&item.Status, &item.Creado_em, &item.Atualizado_em)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Item não encontrado com id %s", id)
		}
		return nil, err
	}

	return &item, nil
}

//func (r *MySQLItemRepository) GetItens(item entity.Item) (entity.Item, error) {
//
//}
//func (r *MySQLItemRepository) UpdateItem(id int) (entity.Item, error) {
//
//}
//func (r *MySQLItemRepository) DeleteItem(id int) (entity.Item, error) {
//
//}
