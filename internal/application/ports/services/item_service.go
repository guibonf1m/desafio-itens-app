package services

import entity "desafio-itens-app/internal/domain/item"

type ItemService interface {
	GetItem(id int) (*entity.Item, error)
	AddItem(item entity.Item) (entity.Item, error)
	GetItens() ([]entity.Item, error)
	GetItensFiltrados(status *entity.Status, limit int) ([]entity.Item, int, int, error)
	GetItensPaginados(page, pageSize int) ([]entity.Item, int, error)
	GetItensFiltradosPaginados(status *entity.Status, page, pageSize int) ([]entity.Item, int, error)
	UpdateItem(item entity.Item) error
	DeleteItem(id int) error
}
