package http

import (
	entity "desafio-itens-app/internal/domain/item"
)

type CreateItemRequest struct {
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Preco     float64 `json:"preco"`
	Estoque   int     `json:"estoque"`
}
type ItemResponse struct {
	ID        int           `json:"id"`
	Code      string        `json:"code"`
	Nome      string        `json:"nome"`
	Descricao string        `json:"descricao"`
	Preco     float64       `json:"preco"`
	Estoque   int           `json:"estoque"`
	Status    entity.Status `json:"status"`
}
