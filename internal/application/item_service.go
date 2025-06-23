package application

import "desafio-itens-app/internal/domain"

type ItemService struct {
	repo domain.ItemRepository
}

func NewItemService(repo domain.ItemRepository) *ItemService {
	return &ItemService{
		repo: repo,
	}
}
