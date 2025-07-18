package item

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItem_IsValid_Sucess(t *testing.T) {
	//ARRANGE - Prepara Dados
	item := Item{
		Nome:    "itemtest",
		Preco:   99.99,
		Estoque: 10.00,
		Status:  StatusAtivo,
	}

	//ACT - Executa a função
	err := item.IsValid()

	//ASSERT - Verificar resultado
	assert.NoError(t, err)
}

func TestItem_IsValid_NomeVazio(t *testing.T) {
	//ARRANGE - Prepara Dados
	item := Item{
		Nome:    "",
		Preco:   99.99,
		Estoque: 10.00,
		Status:  StatusAtivo,
	}

	//ACT - Executa a função
	err := item.IsValid()

	//ASSERT
	assert.Error(t, err)
	assert.Equal(t, "Nome é obrigatório", err.Error())
}

func TestItem_IsValid_PrecoInvalido(t *testing.T) {
	//ARRANGE
	item := Item{
		Nome:    "Produto Teste",
		Preco:   0,
		Estoque: 10,
		Status:  StatusAtivo,
	}

	//ACT
	err := item.IsValid()

	//ASSERT
	assert.Error(t, err)
	assert.Equal(t, "Preço deve ser maior que zero", err.Error())
}

func TestItem_IsValid_StatusInvalido(t *testing.T) {
	//ARRANGE
	item := Item{
		Nome:    "Produto Teste",
		Preco:   100,
		Estoque: 10,
		Status:  "Outro Status",
	}

	//ACT
	err := item.IsValid()

	//ASSERT
	assert.Error(t, err)
	assert.Equal(t, "status deve ser 'active' ou 'inative'", err.Error())
}

func TestItem_IsValid_EstoqueInvalido(t *testing.T) {
	//ARRANGE
	item := Item{
		Nome:    "Produto Teste",
		Preco:   100.00,
		Estoque: -10,
		Status:  StatusAtivo,
	}

	//ACT
	err := item.IsValid()

	//ASSERT
	assert.Error(t, err)
	assert.Equal(t, "Estoque não pode ser negativo", err.Error())
}
