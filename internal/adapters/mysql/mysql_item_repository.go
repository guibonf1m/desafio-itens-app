package mysql

import (
	"database/sql"
	entity "desafio-itens-app/internal/domain/item"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLItemRepository struct {
	db *sqlx.DB
}

func Conectar() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/desafio_itens?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewMySQLItemRepository(db *sqlx.DB) *MySQLItemRepository {
	//Cria uma nova struct preenchida e devolve o endereço (ponteiro) dessa struct.
	//Retorna o repositório prontinho para ser usado
	return &MySQLItemRepository{db: db}

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

func (r *MySQLItemRepository) GetItens() ([]entity.Item, error) {

	var itens []entity.Item

	const query = "  SELECT \n id, code, nome, descricao, preco, estoque, status,\n    " +
		"created_at AS created_at, updated_at AS updated_at\n  FROM itens"

	err := r.db.Select(&itens, query)

	if err != nil {
		return nil, err
	}
	return itens, nil
}

func (r *MySQLItemRepository) GetItensFiltrados(status *entity.Status, limit int) ([]entity.Item, error) {
	var (
		query = `SELECT id, code, nome, descricao, preco, estoque, status, created_at, updated_at FROM itens`
		args  []interface{}
	)

	if status != nil {
		query += " WHERE status = ?"
		args = append(args, *status)
	}
	query += " ORDER BY updated_at DESC LIMIT ?"
	args = append(args, limit)

	var itens []entity.Item
	if err := r.db.Select(&itens, query, args...); err != nil {
		return nil, err
	}
	return itens, nil
}

func (r *MySQLItemRepository) CountItens(status *entity.Status) (int, error) {
	var query string
	var args []interface{}

	if status != nil {
		query = "SELECT COUNT(*) FROM itens WHERE status = ?"
		args = append(args, *status)
	} else {
		query = "SELECT COUNT(*) FROM itens"
	}

	var total int
	if err := r.db.Get(&total, query, args...); err != nil {
		return 0, err
	}
	return total, nil
}

func (r *MySQLItemRepository) UpdateItem(item entity.Item) error {

	query := "UPDATE itens SET nome = ?, descricao = ?, preco = ?, estoque = ?, status = ?  WHERE id = ?"

	result, err := r.db.Exec(
		query,
		item.Nome,
		item.Descricao,
		item.Preco,
		item.Estoque,
		item.Status,
		item.ID,
	)

	if err != nil {
		return err
	}

	//RowsAffected(): consulta no result quantas linhas efetivamente mudaram de valor.
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	//Se rows == 0, significa que nenhum registro com aquele id foi encontrado (e, portanto, nada foi atualizado).
	if rows == 0 {
		return fmt.Errorf("nenhum item encontrado com id %d", item.ID)
	}

	return nil
}

func (r *MySQLItemRepository) DeleteItem(id int) error {

	result, err := r.db.Exec("DELETE FROM itens WHERE id = ?", id)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("Nenhum item encontrado para o id %d", id)
	}

	return nil

}
