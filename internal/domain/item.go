package domain

import "time"

type Status string

const (
	StatusAtivo   Status = "ativo"
	StatusInativo Status = "inativo"
)

type Item struct {
	ID            int
	Code          string
	Nome          string
	Descricao     string
	Preco         float64
	Estoque       int
	Status        Status
	Creado_em     time.Time
	Atualizado_em time.Time
}
