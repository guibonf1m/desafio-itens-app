package http

import (
	entity "desafio-itens-app/internal/domain/item"
	"strings"
	"time"
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
	CreatedAt *time.Time    `json:"created_at"`
	UpdatedAt *time.Time    `json:"updated_at"`
}

type UpdateItemRequest struct {
	Preco     *float64 `json:"preco,omitempty"`
	Estoque   *int     `json:"estoque,omitempty"`
	Descricao *string  `json:"descricao,omitempty"`
}

func (r *CreateItemRequest) ToEntity() entity.Item {
	return entity.Item{
		Nome:      r.Nome,
		Descricao: r.Descricao,
		Preco:     r.Preco,
		Estoque:   r.Estoque,
	}
}

func FromEntity(item entity.Item) ItemResponse {
	return ItemResponse{
		ID:        item.ID,
		Code:      item.Code,
		Nome:      item.Nome,
		Descricao: item.Descricao,
		Preco:     item.Preco,
		Estoque:   item.Estoque,
		Status:    item.Status,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

func (r *UpdateItemRequest) ApplyTo(item *entity.Item) {

	if r.Descricao != nil {
		item.Descricao = strings.TrimSpace(*r.Descricao)
	}
	if r.Preco != nil && *r.Preco > 0 {
		item.Preco = *r.Preco
	}
	if r.Estoque != nil && *r.Estoque >= 0 {
		item.Estoque = *r.Estoque
	}
}
