package dto

import (
	userDomain "desafio-itens-app/internal/domain/user"
	"time"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"omitempty,oneof=admin user"`
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
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token     string       `json:"token"`
	ExpiresIn int64        `json:"expires_in"`
	User      UserResponse `json:"user"`
}

type ListUsersResponse struct {
	Users      []UserResponse `json:"users"`
	Total      int64          `json:"total"`
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
	TotalPages int            `json:"total_pages"`
}

// ToEntity converte CreateUserRequest → User
func (r *CreateUserRequest) ToEntity() userDomain.User {
	role := userDomain.RoleUser
	if r.Role == "admin" {
		role = userDomain.RoleAdmin
	}

	return userDomain.User{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
		Role:     role,
	}
}

func FromUserEntity(user userDomain.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      string(user.Role),
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
