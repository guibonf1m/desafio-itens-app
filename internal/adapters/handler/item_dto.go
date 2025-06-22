package handler

import "time"

type ItemRequest struct {
	ID            int       `json:"id"`
	Code          int       `json:"code"`
	Nome          string    `json:"nome"`
	Descricao     string    `json:"descricao"`
	Preco         float64   `json:"preco"`
	Estoque       int       `json:"estoque"`
	Status        int       `json:"status"`
	Creado_em     time.Time `json:"creado_em"`
	Atualizado_em time.Time `json:"atualizado_em"`
}
type ItemResponse struct {
	ID        int     `json:"id"`
	Code      int     `json:"code"`
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Preco     float64 `json:"preco"`
}
