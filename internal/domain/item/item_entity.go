package item

import "time"

type Status string

const (
	StatusAtivo   Status = "active"
	StatusInativo Status = "inactive"
)

type Item struct {
	ID            int        `db:"id"`
	Code          string     `db:"code"`
	Nome          string     `db:"nome"`
	Descricao     string     `db:"descricao"`
	Preco         float64    `db:"preco"`
	Estoque       int        `db:"estoque"`
	Status        Status     `db:"status"`
	Creado_em     *time.Time `db:"created_at"`
	Atualizado_em *time.Time `db:"updated_at"`
}
