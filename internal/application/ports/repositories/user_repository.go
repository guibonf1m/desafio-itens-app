package repositories

import (
	"context"
	"desafio-itens-app/internal/domain/user"
)

type UserRepository interface {
	Create(user user.User) (user.User, error)
	GetById(id int) (*user.User, error)
	List(ctx context.Context, limit, offset int) ([]*user.User, int64, error)
	GetByUsername(username string) (*user.User, error)
	GetByEmail(email string) (*user.User, error)
	Update(user user.User) error
	Delete(id int) error
	UserNameExists(username string) (bool, error)
	EmailExists(email string) (bool, error)
}
