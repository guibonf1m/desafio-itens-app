package service

import (
	"desafio-itens-app/internal/application/service/mocks"
	entity "desafio-itens-app/internal/domain/item"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddItem_WhenValidationFails_ReturnsError(t *testing.T) {
	// Teste focado só na validação

	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	invalidItem := entity.Item{
		Nome:  "",
		Preco: -1,
	}

	//ACT
	result, err := service.AddItem(invalidItem)

	//ASSERT
	assert.Error(t, err)
	assert.Equal(t, entity.Item{}, result)
	mockRepo.AssertExpectations(t)
}

func TestAddItem_WhenEstoqueZero_SetsStatusInativo(t *testing.T) {
	// ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	itemComEstoqueZero := entity.Item{
		Nome:    "Produto Válido",
		Preco:   100.0,
		Estoque: 0, // Estoque zero deve gerar StatusInativo
	}

	// Mock: CodeExists retorna false (código disponível)
	mockRepo.On("CodeExists", "PROD123").Return(false, nil)

	// Mock: AddItem deve receber item com StatusInativo
	expectedItem := itemComEstoqueZero
	expectedItem.Status = entity.StatusInativo
	expectedItem.Code = "PROD123"
	mockRepo.On("AddItem", expectedItem).Return(expectedItem, nil)

	// ACT
	result, err := service.AddItem(itemComEstoqueZero)

	// ASSERT
	assert.NoError(t, err)
	assert.Equal(t, entity.StatusInativo, result.Status)
	mockRepo.AssertExpectations(t)
}

func TestAddItem_WhenEstoquePositivo_SetsStatusAtivo(t *testing.T) {
	// Teste focado só no estoque > 0
}

func TestAddItem_WhenGenerateCodeFails_ReturnsError(t *testing.T) {
	// Teste focado só no erro de código
}

func TestAddItem_WhenRepositoryFails_ReturnsError(t *testing.T) {
	// Teste focado só no erro do repository
}

func TestAddItem_WhenAllValid_ReturnsItemWithCorrectStatus(t *testing.T) {
	// Teste do caminho feliz completo
}
