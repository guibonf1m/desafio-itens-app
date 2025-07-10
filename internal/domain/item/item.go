package item

import "time"

type Status string

const (
	StatusAtivo   Status = "active"
	StatusInativo Status = "inactive"
)

type Item struct {
	ID        int
	Code      string
	Nome      string
	Descricao string
	Preco     float64
	Estoque   int
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy *int
	UpdateBy  *int
}
