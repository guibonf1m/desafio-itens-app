package item

import (
	"errors"
	"time"
)

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

func (i *Item) IsValid() error {
	if i.Nome == "" {
		return errors.New("Nome é obrigatório")
	}
	if i.Preco <= 0 {
		return errors.New("Preço deve ser maior que zero")
	}
	if i.Estoque <= 0 {
		return errors.New("Estoque não pode ser negativo")
	}

	if i.Status != StatusAtivo && i.Status != StatusInativo {
		return errors.New("status deve ser 'active' ou 'inative'")
	}

	return nil
}
