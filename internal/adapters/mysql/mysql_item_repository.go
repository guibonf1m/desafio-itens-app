package mysql

import (
	"database/sql"                                  // Para sql.ErrNoRows
	entity "desafio-itens-app/internal/domain/item" // Importa entidades do domain
	"errors"                                        // Para criar erros simples
	"fmt"                                           // Para formatar erros/strings
	_ "github.com/go-sql-driver/mysql"              // Driver MySQL (blank import)
	"github.com/jmoiron/sqlx"                       // Extensão do database/sql
)

type MySQLItemRepository struct { // Struct que implementa ItemRepository
	db *sqlx.DB // Conexão com o banco
}

func Conectar() (*sqlx.DB, error) { // ❌ CRÍTICO: credenciais hardcoded
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/desafio_itens?parseTime=true")
	if err != nil { // 🛡️ VALIDATION GUARD
		return nil, err
	}
	return db, nil
}

func NewMySQLItemRepository(db *sqlx.DB) *MySQLItemRepository { // Factory function
	return &MySQLItemRepository{db: db} // Injeta dependência da conexão
}

func (r *MySQLItemRepository) AddItem(item entity.Item) (entity.Item, error) {
	result, err := r.db.Exec("INSERT INTO itens (code, nome, preco, descricao, estoque, status) VALUES (?, ?, ?, ?, ?, ?)",
		item.Code, item.Nome, item.Preco, item.Descricao, item.Estoque, item.Status) // 🌐 EXTERNAL CALL + prepared statement
	if err != nil { // 🛡️ VALIDATION GUARD
		return entity.Item{}, err
	}

	id, err := result.LastInsertId() // Pega ID auto-increment do banco
	if err != nil {
		return entity.Item{}, err
	}

	item.ID = int(id) // 🔄 TRANSFORMATION: int64 → int
	return item, nil  // Retorna item com ID preenchido
}

func (r *MySQLItemRepository) GetItem(id int) (*entity.Item, error) {
	if id <= 0 { // 🛡️ VALIDATION GUARD
		return nil, errors.New("Id não pode ser 0.")
	}

	query := "SELECT * FROM itens WHERE id = ?" // Prepared statement
	row := r.db.QueryRow(query, id)             // 🌐 EXTERNAL CALL - busca uma linha

	var item entity.Item
	err := row.Scan(&item.ID, &item.Code, &item.Nome, &item.Descricao, &item.Preco, &item.Estoque,
		&item.Status, &item.Creado_em, &item.Atualizado_em) // 🔄 TRANSFORMATION: SQL → struct
	if err != nil {
		if err == sql.ErrNoRows { // ⚙️ BUSINESS RULE: trata "não encontrado"
			return nil, fmt.Errorf("Item não encontrado com id %d", id)
		}
		return nil, err
	}

	return &item, nil // Retorna ponteiro para o item
}

func (r *MySQLItemRepository) GetItens() ([]entity.Item, error) {
	var itens []entity.Item // Slice para múltiplos resultados

	const query = "SELECT id, code, nome, descricao, preco, estoque, status, created_at AS created_at, updated_at AS updated_at FROM itens"

	err := r.db.Select(&itens, query) // 🌐 EXTERNAL CALL + 🔄 TRANSFORMATION automática
	if err != nil {
		return nil, err
	}
	return itens, nil
}

func (r *MySQLItemRepository) GetItensFiltrados(status *entity.Status, limit int) ([]entity.Item, error) {
	var (
		query = `SELECT id, code, nome, descricao, preco, estoque, status, created_at, updated_at FROM itens`
		args  []interface{} // Slice para argumentos dinâmicos
	)

	if status != nil { // ⚙️ BUSINESS RULE: filtro condicional
		query += " WHERE status = ?" // Query dinâmica
		args = append(args, *status) // Desreferencia ponteiro
	}
	query += " ORDER BY updated_at DESC LIMIT ?" // Ordenação + paginação
	args = append(args, limit)

	var itens []entity.Item
	if err := r.db.Select(&itens, query, args...); err != nil { // 🌐 EXTERNAL CALL com args dinâmicos
		return nil, err
	}
	return itens, nil
}

func (r *MySQLItemRepository) CountItens(status *entity.Status) (int, error) {
	var query string
	var args []interface{}

	if status != nil { // ⚙️ BUSINESS RULE: contagem condicional
		query = "SELECT COUNT(*) FROM itens WHERE status = ?"
		args = append(args, *status)
	} else {
		query = "SELECT COUNT(*) FROM itens" // Conta todos
	}

	var total int
	if err := r.db.Get(&total, query, args...); err != nil { // 🌐 EXTERNAL CALL - pega valor único
		return 0, err
	}
	return total, nil
}

func (r *MySQLItemRepository) CodeExists(code string) (bool, error) {

	query := "SELECT EXISTS(SELECT 1 FROM itens WHERE code = ?)"

	var exists bool
	err := r.db.QueryRow(query, code).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("erro ao verificar o código: %w", err)
	}
	return exists, nil
}

func (r *MySQLItemRepository) UpdateItem(item entity.Item) error {
	query := "UPDATE itens SET nome = ?, descricao = ?, preco = ?, estoque = ?, status = ? WHERE id = ?"

	result, err := r.db.Exec(query, item.Nome, item.Descricao, item.Preco, item.Estoque, item.Status, item.ID) // 🌐 EXTERNAL CALL
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected() // Verifica quantas linhas foram alteradas
	if err != nil {
		return err
	}

	if rows == 0 { // ⚙️ BUSINESS RULE: item não existe
		return fmt.Errorf("nenhum item encontrado com id %d", item.ID)
	}

	return nil
}

func (r *MySQLItemRepository) DeleteItem(id int) error {
	result, err := r.db.Exec("DELETE FROM itens WHERE id = ?", id) // 🌐 EXTERNAL CALL
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected() // Verifica se deletou algo
	if err != nil {
		return err
	}

	if rows == 0 { // ⚙️ BUSINESS RULE: item não existia
		return fmt.Errorf("Nenhum item encontrado para o id %d", id)
	}

	return nil
}
