package application

import (
	"desafio-itens-app/internal/domain/item"
	"errors"
	"fmt"
)

type itemService struct {
	repo item.ItemRepository
}

type ItemService interface {
	GetItem(id int) (*item.Item, error)
	AddItem(item item.Item) (item.Item, error)
	GetItens(page, size int) ([]item.Item, error)
	UpdateItem(item item.Item) error
	DeleteItem(id int) error
}

func NewItemService(repo item.ItemRepository) *itemService {
	return &itemService{
		repo: repo,
	}
}

func (s *itemService) AddItem(item item.Item) (item.Item, error) {

	if item.Preco <= 0 {
		return item.Item{}, errors.New("O produto tem preço inválido.")
	}

	if item.Estoque <= 0 {
		return item.Item{}, errors.New("O produto tem estoque inválido.")
	}

	switch item.Status {
	case item.StatusAtivo:
		if item.Estoque <= 0 {
			return item.Item{}, errors.New("Item ativo deve ter estoque maior do que zero")
		}
	case item.StatusInativo:
		if item.Estoque > 0 {
			return item.Item{}, errors.New("Item inativo não pode ter estoque disponível")
		}
	default:
		return item.Item{}, errors.New("status inválido")
	}

	code, err := item.GenerateItemCode(item.Nome)
	if err != nil {
		return item.Item{}, err
	}
	item.Code = code

	err = s.repo.AddItem(item)
	if err != nil {
		return item.Item{}, err
	}
	return item, nil
}

func (s *itemService) GetItem(id int) (*item.Item, error) {

	if id == 0 {
		return nil, fmt.Errorf("O id não pode ser 0.")
	}

	item, err := s.repo.GetItem(id)
	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar o item", err)
	}

	return item, nil
}

//func (s *itemService) GetItens() ([]entity.Item, error) {
//
//}
//
//func (s *itemService) UpdateItem(item entity.Item) error {
//
//}
//
//func (s *itemService) DeleteItem(id int) error {
//
//}
