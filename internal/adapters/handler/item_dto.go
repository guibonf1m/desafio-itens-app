package handler

import (
	"desafio-itens-app/internal/domain"
)

type CreateItemRequest struct {
	Nome      string        `json:"nome"`
	Descricao string        `json:"descricao"`
	Preco     float64       `json:"preco"`
	Estoque   int           `json:"estoque"`
	Status    domain.Status `json:"status"`
}
type ItemResponse struct {
	ID        int           `json:"id"`
	Code      int           `json:"code"`
	Nome      string        `json:"nome"`
	Descricao string        `json:"descricao"`
	Preco     float64       `json:"preco"`
	Estoque   int           `json:"estoque"`
	Status    domain.Status `json:"status"`
}
