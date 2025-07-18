package user

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestUser_IsValid_Sucess(t *testing.T) {
	//ARRANGE - Prepara Dados
	user := User{
		Username: "testuser",
		Email:    "test@email.com",
		Password: "123456",
		Role:     RoleUser,
	}

	//ACT - Executa a função
	err := user.IsValid()

	//ASSERT - Verificar resultado
	assert.NoError(t, err)
}

func TestUser_IsValid_UsernameEmpty(t *testing.T) {
	//ARRANGE - Prepara Dados
	user := User{
		Username: "",
		Email:    "test@email.com",
		Password: "123456",
		Role:     RoleUser,
	}

	//ACT - Executa a função
	err := user.IsValid()

	//ASSERT - Verificar resultado
	assert.Error(t, err)
	assert.Equal(t, "Username é obrigatório", err.Error())
}

func TestUser_IsValid_PasswordEmpty(t *testing.T) {
	//ARRANGE - Prepara Dados
	user := User{
		Username: "bonfim",
		Email:    "test@email.com",
		Password: "",
		Role:     RoleUser,
	}

	//ACT - Executa a função
	err := user.IsValid()

	//ASSERT - Verificar resultado
	assert.Error(t, err)
	assert.Equal(t, "Password é obrigatório.", err.Error())
}

func TestUser_IsValid_UsernameTooShort(t *testing.T) {
	//ARRANGE - Prepara Dados
	user := User{
		Username: "ab",
		Email:    "test@email.com",
		Password: "123456",
		Role:     RoleUser,
	}

	//ACT - Executa a função
	err := user.IsValid()

	//ASSERT - Verificar resultado
	assert.Error(t, err)
	assert.Equal(t, "Username deve ter pelo menos 3 letras.", err.Error())
}

func TestUser_IsValid_UsernameTooLong(t *testing.T) {
	//ARRANGE - Prepara Dados
	user := User{
		Username: "a" + strings.Repeat("b", 50),
		Email:    "email@test.com",
		Password: "123456",
		Role:     RoleUser,
	}

	//ACT - Executa a função
	err := user.IsValid()

	//ASSERT - Verifica resultado
	assert.Error(t, err)
	assert.Equal(t, "username deve ter no máximo 50 letras", err.Error())

}

func TestUser_IsValid_InvalidRole(t *testing.T) {
	// ARRANGE - Prepara Dados
	user := User{
		Username: "testeuser",
		Email:    "email@test.com",
		Password: "123456",
		Role:     "invalid_role",
	}

	//ACT - Executa a função
	err := user.IsValid()

	//ASSERT - Verifica resultado
	assert.Error(t, err)
	assert.Equal(t, "role deve ser 'admin' ou 'user'", err.Error())
}
