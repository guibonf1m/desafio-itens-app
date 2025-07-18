package services

import (
	"context"
	"desafio-itens-app/internal/adapters/http/dto"
	userDomain "desafio-itens-app/internal/domain/user"
)

type UserService interface {
	CreateUser(user userDomain.User) (userDomain.User, error)
	GetUser(id int) (*userDomain.User, error)
	ListUsers(ctx context.Context, page, limit int) (*dto.ListUsersResponse, error)
	GetUserByUsername(username string) (*userDomain.User, error)
	UpdateUser(user userDomain.User) error
	DeleteUser(id int) error
	ValidateCredentials(username, password string) (*userDomain.User, error)
}
