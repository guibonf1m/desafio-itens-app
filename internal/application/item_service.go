package application

import (
	"desafio-itens-app/internal/domain"
	"errors"
	"fmt"
	"go/doc/comment"
	"strconv"
)

type itemService struct {
	repo domain.ItemRepository
}

type ItemService interface {
	GetItem(id int) (*domain.Item, error)
	AddItem(item domain.Item) (domain.Item, error)
	GetItens(page, size int) ([]domain.Item, error)
	UpdateItem(item domain.Item) error
	DeleteItem(id int) error
}

func NewItemService(repo domain.ItemRepository) *itemService {
	return &itemService{
		repo: repo,
	}
}

func (s *itemService) AddItem(item domain.Item) (domain.Item, error) {

	if item.Preco <= 0 {
		return domain.Item{}, errors.New("O produto tem preço inválido.")
	}

	if item.Estoque <= 0 {
		return domain.Item{}, errors.New("O produto tem estoque inválido.")
	}

	switch item.Status {
	case domain.StatusAtivo:
		if item.Estoque <= 0 {
			return domain.Item{}, errors.New("Item ativo deve ter estoque maior do que zero")
		}
	case domain.StatusInativo:
		if item.Estoque > 0 {
			return domain.Item{}, errors.New("Item inativo não pode ter estoque disponível")
		}
	default:
		return domain.Item{}, errors.New("status inválido")
	}

	code, err := domain.GenerateItemCode(item.Nome)
	if err != nil {
		return domain.Item{}, err
	}
	item.Code = code

	err = s.repo.AddItem(item)
	if err != nil {
		return domain.Item{}, err
	}
	return item, nil
}

func (s *itemService) GetItem(id int) (*domain.Item, error) {

	if id == 0 {
		return nil, fmt.Errorf("O id não pode ser 0.")
	}

	item, err := s.repo.GetItem(id)
	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar o item", err)
	}

	return item, nil
}

func (s *itemService) GetItens() ([]domain.Item, error) {

}

func (s *itemService) UpdateItem(item domain.Item) error {

}

func (s *itemService) DeleteItem(id int) error {

}
