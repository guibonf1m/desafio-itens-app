package domain

import "time"

type Item struct {
	ID            int
	Code          int
	Nome          string
	Descricao     string
	Preco         float64
	Estoque       int
	Status        int
	Creado_em     time.Time
	Atualizado_em time.Time
}
