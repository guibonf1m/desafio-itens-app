package http

import (
	userDomain "desafio-itens-app/internal/domain/user"
	"time"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
}
type UpdateUserRequest struct {
	Username *string `json:"username,omitempty" binding:"omitempty,min=3,max=50"`
	Password *string `json:"password,omitempty" binding:"omitempty,min=6"`
}

type UserResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ToEntity converte CreateUserRequest → User
func (r *CreateUserRequest) ToEntity() userDomain.User {
	return userDomain.User{
		Username: r.Username,
		Password: r.Password,
	}
}

func FromUserEntity(user userDomain.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdatedAt,
	}
}

// ApplyTo aplica UpdateUserRequest em User existente
func (r *UpdateUserRequest) ApplyTo(user *userDomain.User) {
	if r.Username != nil {
		user.Username = *r.Username
	}
	if r.Password != nil {
		user.Password = *r.Password
	}
}
