package application

import (
	entity "desafio-itens-app/internal/domain/item"
	"errors"
	"fmt"
)

type itemService struct {
	repo entity.ItemRepository
}

type ItemService interface {
	GetItem(id int) (*entity.Item, error)
	AddItem(item entity.Item) (entity.Item, error)
	//GetItens(page, size int) ([]entity.Item, error)
	//UpdateItem(item entity.Item) error
	//DeleteItem(id int) error
}

func NewItemService(repo entity.ItemRepository) *itemService {
	return &itemService{
		repo: repo,
	}
}

func (s *itemService) AddItem(item entity.Item) (entity.Item, error) {

	if item.Preco <= 0 {
		return entity.Item{}, errors.New("O produto tem preço inválido.")
	}

	if item.Estoque < 0 {
		return entity.Item{}, errors.New("O produto tem estoque inválido.")
	}

	if item.Estoque == 0 {
		item.Status = entity.StatusInativo
	} else {
		item.Status = entity.StatusAtivo
	}

	code, err := entity.GenerateItemCode(item.Nome)
	if err != nil {
		return entity.Item{}, err
	}
	item.Code = code

	itemCriado, err := s.repo.AddItem(item)
	if err != nil {
		return entity.Item{}, err
	}

	return itemCriado, nil

}

func (s *itemService) GetItem(id int) (*entity.Item, error) {

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
