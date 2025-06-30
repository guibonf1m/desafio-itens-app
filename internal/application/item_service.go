package application

import (
	entity "desafio-itens-app/internal/domain/item"
	"errors"
	"fmt"
	"math"
)

type itemService struct {
	repo entity.ItemRepository
}

type ItemService interface {
	GetItem(id int) (*entity.Item, error)
	AddItem(item entity.Item) (entity.Item, error)
	GetItens() ([]entity.Item, error)
	GetItensFiltrados(status *entity.Status, limit int) ([]entity.Item, int, int, error)
	UpdateItem(item entity.Item) error
	DeleteItem(id int) error
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

func (s *itemService) GetItens() ([]entity.Item, error) {

	itens, err := s.repo.GetItens()
	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar os itens: %w", err)
	}

	return itens, nil
}

func (s *itemService) GetItensFiltrados(status *entity.Status, limit int) (itens []entity.Item, totalItens int, totalPages int, err error) {
	// normaliza o limit
	if limit <= 0 {
		limit = 10
	}

	if limit >= 20 {
		limit = 20
	}

	// 1) conta quantos itens
	totalItens, err = s.repo.CountItens(status)
	if err != nil {
		err = fmt.Errorf("erro ao contar itens: %w", err)
		return
	}

	// 2) busca os itens
	itens, err = s.repo.GetItensFiltrados(status, limit)
	if err != nil {
		err = fmt.Errorf("erro ao buscar itens: %w", err)
		return
	}

	// 3) calcula totalPages
	totalPages = int(math.Ceil(float64(totalItens) / float64(limit)))
	return
}

func (s *itemService) UpdateItem(item entity.Item) error {

	itemExistente, err := s.repo.GetItem(item.ID)

	if err != nil {
		return fmt.Errorf("Erro ao buscar o item: %w", err)
	}

	if itemExistente == nil {
		return fmt.Errorf("Item não encontrado para atualização.")
	}

	if item.Preco <= 0 {
		return fmt.Errorf("O produto tem preço inválido.")
	}

	if item.Estoque < 0 {
		return fmt.Errorf("O produto tem estoque inválido.")
	}

	if item.Estoque == 0 {
		item.Status = entity.StatusInativo
	} else {
		item.Status = entity.StatusAtivo
	}

	if err := s.repo.UpdateItem(item); err != nil {
		return fmt.Errorf("Erro ao atualizar o item: %w", err)
	}
	return nil
}

func (s *itemService) DeleteItem(id int) error {

	if id <= 0 {
		return fmt.Errorf("ID inválido para a exclusão: %d", id)
	}

	err := s.repo.DeleteItem(id)

	if err != nil {
		fmt.Errorf("Erro ao deletar item: %w", err)
	}

	return nil

}
