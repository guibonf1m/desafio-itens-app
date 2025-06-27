package item

type ItemRepository interface {
	GetItem(id int) (*Item, error)
	//GetItens() ([]Item, error)
	AddItem(item Item) (Item, error)
	UpdateItem(item Item) error
	//DeleteItem(id int) error
}
