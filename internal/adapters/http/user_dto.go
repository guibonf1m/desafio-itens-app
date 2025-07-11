package http

import (
	userDomain "desafio-itens-app/internal/domain/user"
	"time"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
type UpdateUserRequest struct {
	Username *string `json:"username,omitempty" binding:"omitempty,min=3,max=50"`
	Email    *string `json:"email,omitempty" binding:"omitempty,email"`
	Password *string `json:"password,omitempty" binding:"omitempty,min=6"`
}

type UserResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ToEntity converte CreateUserRequest → User
func (r *CreateUserRequest) ToEntity() userDomain.User {
	return userDomain.User{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}
}

func FromUserEntity(user userDomain.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// ApplyTo aplica UpdateUserRequest em User existente
func (r *UpdateUserRequest) ApplyTo(user *userDomain.User) {
	if r.Username != nil {
		user.Username = *r.Username
	}
	if r.Email != nil {
		user.Email = *r.Email
	}
	if r.Password != nil {
		user.Password = *r.Password
	}
}
