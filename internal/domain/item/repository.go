package item

type ItemRepository interface {
	GetItem(id int) (*Item, error)
	GetItens() ([]Item, error)
	GetItensFiltrados(status *Status, limit int) ([]Item, error)
	GetItensPaginados(ofsset, limit int) ([]Item, int, error)
	GetItensFiltradosPaginados(status *Status, page, pageSize int) ([]Item, int, error)
	CountItens(status *Status) (int, error)
	CodeExists(code string) (bool, error)
	AddItem(item Item) (Item, error)
	UpdateItem(item Item) error
	DeleteItem(id int) error
}
