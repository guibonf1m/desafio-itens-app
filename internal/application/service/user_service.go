package service

import (
	"context"
	"desafio-itens-app/internal/adapters/http/dto"
	"desafio-itens-app/internal/application/ports/repositories"
	"desafio-itens-app/internal/application/ports/services"
	userDomain "desafio-itens-app/internal/domain/user"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math"
	"strings"
)

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) services.UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user userDomain.User) (userDomain.User, error) {
	// PASSO 1: VALIDAR dados básicos
	if err := user.IsValid(); err != nil {
		return userDomain.User{}, err
	}

	// PASSO 2: VERIFICAR se username já existe
	exists, err := s.repo.UserNameExists(user.Username)
	if err != nil {
		return userDomain.User{}, fmt.Errorf("erro ao verificar username: %w", err)
	}
	if exists {
		return userDomain.User{}, errors.New("username já está em uso")
	}

	hashedPassword, err := s.hashPassword(user.Password)
	if err != nil {
		return userDomain.User{}, fmt.Errorf("erro ao criptografar a senha: %w", err)
	}
	user.Password = hashedPassword

	createdUser, err := s.repo.Create(user)
	if err != nil {
		return userDomain.User{}, fmt.Errorf("erro ao criar o usuário: %w", err)
	}

	return createdUser, nil
}

func (s *userService) GetUser(id int) (*userDomain.User, error) {
	if id <= 0 {
		return nil, errors.New("ID deve ser maior que zero")
	}

	return s.repo.GetById(id)
}

func (s *userService) ListUsers(ctx context.Context, page, limit int) (*dto.ListUsersResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	users, total, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar usuários: %w", err)
	}

	userResponse := make([]dto.UserResponse, len(users))
	for i, user := range users {
		userResponse[i] = dto.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Role:      string(user.Role),
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &dto.ListUsersResponse{
		Users:      userResponse,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}

func (s *userService) GetUserByUsername(username string) (*userDomain.User, error) {
	username = strings.TrimSpace(username)

	if username == "" {
		return nil, errors.New("username não pode está vazio")
	}

	return s.repo.GetByUsername(username)
}

func (s *userService) UpdateUser(user userDomain.User) error {
	if err := user.IsValid(); err != nil {
		return err
	}

	existing, err := s.repo.GetById(user.ID)
	if err != nil {
		return err
	}

	if existing.Username != user.Username {
		exists, err := s.repo.UserNameExists(user.Username)
		if err != nil {
			return fmt.Errorf("erro ao verificar username: %w", err)
		}
		if exists {
			return errors.New("username já está em uso")
		}
	}

	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id int) error {
	if id <= 0 {
		return errors.New("ID deve ser maior que zero")
	}

	return s.repo.Delete(id)
}

func (s *userService) ValidateCredentials(username, password string) (*userDomain.User, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		// SEGURANÇA: Não revela se usuário existe ou não
		return nil, errors.New("credenciais inválidas")
	}

	// PASSO 2: VERIFICAR se senha está correta
	if !s.checkPassword(password, user.Password) {
		return nil, errors.New("credenciais inválidas")
	}
	// PASSO 3: CREDENCIAIS CORRETAS - retorna usuário
	return user, nil
}

func (s *userService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *userService) checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
