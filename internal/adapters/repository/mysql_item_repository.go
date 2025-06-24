package repository

import (
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
