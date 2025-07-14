package ports

import (
	"context"
	"desafio-itens-app/internal/adapters/http/dto"
	entity "desafio-itens-app/internal/domain/item"
	userDomain "desafio-itens-app/internal/domain/user"
)

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

type UserService interface {
	CreateUser(user userDomain.User) (userDomain.User, error)
	GetUser(id int) (*userDomain.User, error)
	ListUsers(ctx context.Context, page, limit int) (*dto.ListUsersResponse, error)
	GetUserByUsername(username string) (*userDomain.User, error)
	UpdateUser(user userDomain.User) error
	DeleteUser(id int) error
	ValidateCredentials(username, password string) (*userDomain.User, error)
}
