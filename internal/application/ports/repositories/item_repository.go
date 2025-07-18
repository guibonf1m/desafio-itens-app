package repositories

import "desafio-itens-app/internal/domain/item"

type ItemRepository interface {
	GetItem(id int) (*item.Item, error)
	GetItens() ([]item.Item, error)
	GetItensFiltrados(status *item.Status, limit int) ([]item.Item, error)
	GetItensPaginados(ofsset, limit int) ([]item.Item, int, error)
	GetItensFiltradosPaginados(status *item.Status, page, pageSize int) ([]item.Item, int, error)
	CountItens(status *item.Status) (int, error)
	CodeExists(code string) (bool, error)
	AddItem(item item.Item) (item.Item, error)
	UpdateItem(item item.Item) error
	DeleteItem(id int) error
}
