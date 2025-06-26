package item

import "time"

type Status string

const (
	StatusAtivo   Status = "active"
	StatusInativo Status = "inactive"
)

type Item struct {
	ID            int
	Code          string
	Nome          string
	Descricao     string
	Preco         float64
	Estoque       int
	Status        Status
	Creado_em     *time.Time `db:"created_at"`
	Atualizado_em *time.Time `db:"updated_at"`
}
