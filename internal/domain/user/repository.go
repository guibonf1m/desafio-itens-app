package user

import "context"

type UserRepository interface {
	Create(user User) (User, error)
	GetById(id int) (*User, error)
	List(ctx context.Context, limit, offset int) ([]*User, int64, error)
	GetByUsername(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user User) error
	Delete(id int) error
	UserNameExists(username string) (bool, error)
	EmailExists(email string) (bool, error)
}
