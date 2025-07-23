package service

import (
	"desafio-itens-app/internal/application/service/mocks"
	domain "desafio-itens-app/internal/domain/user"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestUserService_CreateUser_Success(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	mockRepo.On("UserNameExists", "userexistente").Return(false, nil)
	mockRepo.On("Create", mock.MatchedBy(func(user domain.User) bool {
		return user.Username == "userexistente" && user.Email == "teste@email.com"
	})).Return(domain.User{
		ID:       1,
		Username: "userexistente",
		Email:    "teste@email.com",
		Password: "hashed_password",
		Role:     domain.RoleUser,
	}, nil)

	service := NewUserService(mockRepo)

	testUser := domain.User{
		Username: "userexistente",
		Email:    "teste@email.com",
		Password: "123456",
		Role:     domain.RoleUser,
	}

	//ACT
	result, err := service.CreateUser(testUser)

	// ASSERT
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.NotEqual(t, "123456", result.Password)
	mockRepo.AssertExpectations(t)
}

func TestUserService_CreateUser_InvalidUser(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	service := NewUserService(mockRepo)

	testUser := domain.User{
		Username: "",
		Email:    "test@email.com",
		Password: "123456",
		Role:     domain.RoleUser,
	}

	//ACT
	result, err := service.CreateUser(testUser)

	//ASSERT
	assert.Error(t, err)
	assert.Equal(t, domain.User{}, result)
	mockRepo.AssertExpectations(t)
}

func TestUserService_CreateUser_UsernameExists(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	mockRepo.On("UserNameExists", "userexistente").Return(true, nil)

	service := NewUserService(mockRepo)

	testUser := domain.User{
		Username: "userexistente",
		Email:    "testuser@email.com",
		Password: "123456",
		Role:     domain.RoleUser,
	}

	//ACT
	result, err := service.CreateUser(testUser)

	//ASSERT
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "username já está em uso")
	assert.Equal(t, domain.User{}, result)
	mockRepo.AssertExpectations(t)
}

func TestUserService_CreateUser_UserNameExistsError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	mockRepo.On("UserNameExists", "Bonfim").Return(false, errors.New("erro ao verificar username"))

	testUser := domain.User{
		Username: "Bonfim",
		Email:    "test@email.com",
		Password: "123456",
		Role:     domain.RoleUser,
	}

	service := NewUserService(mockRepo)

	//ACT
	result, err := service.CreateUser(testUser)

	//ASSERT
	assert.Error(t, err)
	assert.Equal(t, domain.User{}, result)
	assert.Contains(t, err.Error(), "erro ao verificar username")
	mockRepo.AssertExpectations(t)
}

func TestUserService_CreateUser_RepositoryCreateError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	mockRepo.On("UserNameExists", "Bonfim").Return(false, nil)
	mockRepo.On("Create", mock.Anything).Return(domain.User{}, errors.New("erro ao criar usuário no banco"))

	service := NewUserService(mockRepo)

	testUser := domain.User{
		Username: "Bonfim",
		Email:    "teste@email.com",
		Password: "123456",
		Role:     domain.RoleUser,
	}

	//ACT
	result, err := service.CreateUser(testUser)

	//ASSERT
	assert.Error(t, err)
	assert.Equal(t, domain.User{}, result)
	assert.Contains(t, err.Error(), "erro ao criar o usuário")
	mockRepo.AssertExpectations(t)
}
func TestUserService_GetUser_Success(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	expectedUser := &domain.User{
		ID:       1,
		Username: "testeuser",
		Email:    "test@email.com",
		Role:     domain.RoleUser,
	}
	mockRepo.On("GetById", 1).Return(expectedUser, nil)
	service := NewUserService(mockRepo)

	//ACT
	result, err := service.GetUser(1)

	//ASSERT
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUser_InvalidID(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	service := NewUserService(mockRepo)

	//ACT
	result, err := service.GetUser(0)

	//ASSERT
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "ID deve ser maior que zero")
	mockRepo.AssertExpectations(t)

}

func TestUserService_GetUser_RepositoryError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	mockRepo.On("GetById", 1).Return((*domain.User)(nil),
		errors.New("erro ao buscar usuário no banco"))
	service := NewUserService(mockRepo)

	//ACT
	result, err := service.GetUser(1)

	//ASSERT
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "erro ao buscar usuário no banco")
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserByUsername_Success(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	expectedUser := &domain.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@email.com",
		Role:     domain.RoleUser,
	}
	mockRepo.On("GetByUsername", "testuser").Return(expectedUser, nil)
	service := NewUserService(mockRepo)

	//ACT
	result, err := service.GetUserByUsername("testuser")

	//ASSERT
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserByUsername_EmptyUsername(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	service := NewUserService(mockRepo)

	//ACT
	result, err := service.GetUserByUsername("")

	//ASSERT
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "username não pode está vazio")
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserByUsername_RepositoryError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	mockRepo.On("GetByUsername", "testuser").Return((*domain.User)(nil),
		errors.New("erro ao buscar usuário"))
	service := NewUserService(mockRepo)

	//ACT
	result, err := service.GetUserByUsername("testuser")

	//ASSERT
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUserService_UpdateUser_Success(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	existingUser := &domain.User{
		ID:       1,
		Username: "olduser",
		Email:    "old@email.com",
		Password: "hashed_password",
		Role:     domain.RoleUser,
	}

	mockRepo.On("GetById", 1).Return(existingUser, nil)
	mockRepo.On("UserNameExists", "newuser").Return(false, nil)
	mockRepo.On("Update", mock.MatchedBy(func(user domain.User) bool {
		return user.ID == 1 && user.Username == "newuser" && user.Email == "new@email.com"
	})).Return(nil)

	service := NewUserService(mockRepo)

	updateUser := domain.User{
		ID:       1,
		Username: "newuser",
		Email:    "new@email.com",
		Password: "newpassword",
		Role:     domain.RoleUser,
	}

	//ACT
	err := service.UpdateUser(updateUser)

	//ASSERT
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_UpdateUser_InvalidUser(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	service := NewUserService(mockRepo)

	invalidUser := domain.User{
		ID:       1,
		Username: "", // username vazio é inválido
		Email:    "test@email.com",
		Password: "123456",
		Role:     domain.RoleUser,
	}

	//ACT
	err := service.UpdateUser(invalidUser)

	//ASSERT
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_UpdateUser_UserNotFound(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	mockRepo.On("GetById", 1).Return((*domain.User)(nil), errors.New("usuário não encontrado"))

	service := NewUserService(mockRepo)

	updateUser := domain.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@email.com",
		Password: "123456",
		Role:     domain.RoleUser,
	}

	//ACT
	err := service.UpdateUser(updateUser)

	//ASSERT
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_UpdateUser_UsernameExists(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	existingUser := &domain.User{
		ID:       1,
		Username: "olduser",
		Email:    "old@email.com",
		Password: "hashed_password",
		Role:     domain.RoleUser,
	}

	mockRepo.On("GetById", 1).Return(existingUser, nil)
	mockRepo.On("UserNameExists", "newuser").Return(true, nil)

	service := NewUserService(mockRepo)

	updateUser := domain.User{
		ID:       1,
		Username: "newuser", // username diferente que já existe
		Email:    "new@email.com",
		Password: "123456",
		Role:     domain.RoleUser,
	}

	//ACT
	err := service.UpdateUser(updateUser)

	//ASSERT
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "username já está em uso")
	mockRepo.AssertExpectations(t)
}

func TestUserService_UpdateUser_RepositoryError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	existingUser := &domain.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@email.com",
		Password: "hashed_password",
		Role:     domain.RoleUser,
	}

	mockRepo.On("GetById", 1).Return(existingUser, nil)
	mockRepo.On("Update", mock.Anything).Return(errors.New("erro ao atualizar usuário no banco"))

	service := NewUserService(mockRepo)

	updateUser := domain.User{
		ID:       1,
		Username: "testuser", // mesmo username, não precisa verificar se existe
		Email:    "newemail@email.com",
		Password: "newpassword",
		Role:     domain.RoleUser,
	}

	//ACT
	err := service.UpdateUser(updateUser)

	//ASSERT
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_DeleteUser_Success(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	mockRepo.On("Delete", 1).Return(nil)
	service := NewUserService(mockRepo)

	//ACT
	err := service.DeleteUser(1)

	//ASSERT
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_DeleteUser_InvalidID(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	service := NewUserService(mockRepo)

	//ACT
	err := service.DeleteUser(0)

	//ASSERT
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ID deve ser maior que zero")
	mockRepo.AssertExpectations(t)
}

func TestUserService_DeleteUser_RepositoryError(t *testing.T) {

	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	mockRepo.On("Delete", 1).Return(errors.New("erro ao deletar usuário"))
	service := NewUserService(mockRepo)

	//ACT
	err := service.DeleteUser(1)

	//ASSERT
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_ValidateCredentials_Success(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	expectedUser := &domain.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@email.com",
		Password: string(hashedPassword),
		Role:     domain.RoleUser,
	}
	mockRepo.On("GetByUsername", "testuser").Return(expectedUser, nil)
	service := NewUserService(mockRepo)

	//ACT
	result, err := service.ValidateCredentials("testuser", "123456")

	//ASSERT
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)
}

func TestUserService_ValidateCredentials_UserNotFound(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	mockRepo.On("GetByUsername", "testuser").Return((*domain.User)(nil), errors.New("usuário não encontrado"))
	service := NewUserService(mockRepo)

	//ACT
	result, err := service.ValidateCredentials("testuser", "123456")

	//ASSERT
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "credenciais inválidas")
	mockRepo.AssertExpectations(t)
}

func TestUserService_ValidateCredentials_WrongPassword(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	user := &domain.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@email.com",
		Password: string(hashedPassword),
		Role:     domain.RoleUser,
	}
	mockRepo.On("GetByUsername", "testuser").Return(user, nil)
	service := NewUserService(mockRepo)

	//ACT
	result, err := service.ValidateCredentials("testuser", "wrongpassword")

	//ASSERT
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "credenciais inválidas")
	mockRepo.AssertExpectations(t)
}

func TestUserService_ListUsers_Success(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	users := []*domain.User{
		{ID: 1, Username: "user1", Email: "user1@test.com", Role: domain.RoleUser},
		{ID: 2, Username: "user2", Email: "user2@test.com", Role: domain.RoleAdmin},
	}
	mockRepo.On("List", mock.Anything, 10, 0).Return(users, int64(2), nil)
	service := NewUserService(mockRepo)

	//ACT
	result, err := service.ListUsers(nil, 1, 10)

	//ASSERT
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Users, 2)
	assert.Equal(t, int64(2), result.Total)
	assert.Equal(t, 1, result.Page)
	assert.Equal(t, 10, result.Limit)
	assert.Equal(t, 1, result.TotalPages)
	mockRepo.AssertExpectations(t)
}

func TestUserService_ListUsers_DefaultPagination(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	users := []*domain.User{}
	mockRepo.On("List", mock.Anything, 10, 0).Return(users, int64(0), nil)
	service := NewUserService(mockRepo)

	//ACT
	result, err := service.ListUsers(nil, 0, 0)

	//ASSERT
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.Page)
	assert.Equal(t, 10, result.Limit)
	mockRepo.AssertExpectations(t)
}

func TestUserService_ListUsers_RepositoryError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewUserRepository(t)
	mockRepo.On("List", mock.Anything, 10, 0).Return([]*domain.User{}, int64(0), errors.New("erro ao listar usuários"))

	service := NewUserService(mockRepo)

	//ACT
	result, err := service.ListUsers(nil, 1, 10)

	//ASSERT
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "erro ao listar usuários")
	mockRepo.AssertExpectations(t)
}
