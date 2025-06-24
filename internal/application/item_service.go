package application

import (
	"desafio-itens-app/internal/domain"
	"errors"
)

type ItemService struct {
	repo domain.ItemRepository
}

func NewItemService(repo domain.ItemRepository) *ItemService {
	return &ItemService{
		repo: repo,
	}
}

func (s *ItemService) AddItem(item domain.Item) (domain.Item, error) {

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

func (s *ItemService) GetItemByID(id int) (domain.Item, error) {

}

//func (s *ItemService) ListItens() ([]domain.Item, error) {

//}

//func (s *ItemService) UpdateItem(item domain.Item) error {

//}

//func (s *ItemService) DeleteItem(id int) error {

//}
