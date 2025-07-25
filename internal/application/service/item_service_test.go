package service

import (
	"desafio-itens-app/internal/application/service/mocks"
	entity "desafio-itens-app/internal/domain/item"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAddItem_WhenSuccess_ReturnsCreatedItem(t *testing.T) {
	// ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	validItem := entity.Item{
		Nome:    "Produto Válido",
		Preco:   100.0,
		Estoque: 10,
	}

	expectedItem := entity.Item{
		ID:      1,
		Nome:    "Produto Válido",
		Preco:   100.0,
		Estoque: 10,
		Status:  entity.StatusAtivo,
		Code:    "PR12345678",
	}

	mockRepo.On("CodeExists", mock.AnythingOfType("string")).Return(false, nil)
	mockRepo.On("AddItem", mock.Anything).Return(expectedItem, nil)

	// ACT
	result, err := service.AddItem(validItem)

	// ASSERT
	assert.NoError(t, err)
	assert.Equal(t, expectedItem.ID, result.ID)
	assert.Equal(t, expectedItem.Nome, result.Nome)
	assert.Equal(t, expectedItem.Preco, result.Preco)
	assert.Equal(t, expectedItem.Estoque, result.Estoque)
	assert.Equal(t, entity.StatusAtivo, result.Status)
	assert.NotEmpty(t, result.Code)
	mockRepo.AssertExpectations(t)
}

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

func TestAddItem_WhenEstoquePositivo_SetsStatusAtivo(t *testing.T) {
	// ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	itemComEstoque := entity.Item{
		Nome:    "Produto Válido",
		Preco:   100.0,
		Estoque: 5,
		Status:  entity.StatusAtivo,
	}

	mockRepo.On("CodeExists", mock.AnythingOfType("string")).Return(false, nil)

	mockRepo.On("AddItem", mock.MatchedBy(func(item entity.Item) bool {
		return item.Nome == "Produto Válido" &&
			item.Preco == 100.0 &&
			item.Estoque == 5 &&
			item.Status == entity.StatusAtivo &&
			item.Code != ""
	})).Return(entity.Item{
		Nome:    "Produto Válido",
		Preco:   100.0,
		Estoque: 5,
		Status:  entity.StatusAtivo,
		Code:    "PR71619235",
	}, nil)

	// ACT
	result, err := service.AddItem(itemComEstoque)

	// ASSERT
	assert.NoError(t, err)
	assert.Equal(t, entity.StatusAtivo, result.Status)
	assert.NotEmpty(t, result.Code)
	mockRepo.AssertExpectations(t)
}

func TestAddItem_WhenEstoqueZero_SetsStatusInativo(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	itemSemEstoque := entity.Item{
		Nome:    "Produto Válido",
		Preco:   100.0,
		Estoque: 0,
		Status:  entity.StatusInativo,
	}

	mockRepo.On("CodeExists", mock.AnythingOfType("string")).Return(false, nil)

	mockRepo.On("AddItem", mock.MatchedBy(func(item entity.Item) bool {
		return item.Nome == "Produto Válido" &&
			item.Preco == 100.0 &&
			item.Estoque == 0 &&
			item.Status == entity.StatusInativo &&
			item.Code != ""
	})).Return(entity.Item{
		Nome:    "Produto Válido",
		Preco:   100.0,
		Estoque: 0,
		Status:  entity.StatusInativo,
		Code:    "PR71619235",
	}, nil)

	//ACT
	result, err := service.AddItem(itemSemEstoque)

	//ASSERT
	assert.NoError(t, err)
	assert.Equal(t, entity.StatusInativo, result.Status)
	assert.NotEmpty(t, result.Code)
	mockRepo.AssertExpectations(t)
}

func TestAddItem_WhenRepositoryFails_ReturnsError(t *testing.T) {
	// Teste focado só no erro do repository

	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	validItem := entity.Item{
		Nome:    "Produto Válido",
		Preco:   100.0,
		Estoque: 10,
		Status:  entity.StatusAtivo,
	}

	mockRepo.On("CodeExists", mock.AnythingOfType("string")).Return(false, nil)
	mockRepo.On("AddItem", mock.Anything).Return(entity.Item{}, assert.AnError)

	//ACT
	result, err := service.AddItem(validItem)

	//ASSERT
	assert.Error(t, err)
	assert.Equal(t, entity.Item{}, result)
	mockRepo.AssertExpectations(t)

}

func TestAddItem_WhenCodeExistsCheckFails_ReturnsError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	validItem := entity.Item{
		Nome:    "Produto Valido",
		Preco:   100.0,
		Estoque: 10,
		Status:  entity.StatusAtivo,
	}

	mockRepo.On("CodeExists", mock.AnythingOfType("string")).Return(false, assert.AnError)

	//ACT
	result, err := service.AddItem(validItem)

	//ASSERT
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "erro ao verificar código")
	assert.Equal(t, entity.Item{}, result)
	mockRepo.AssertExpectations(t)
}

func TestGetItem_WhenIdIsZero_ReturnsError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	//ACT
	result, err := service.GetItem(0)

	//ASSERT
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "O id não pode ser 0.")
	mockRepo.AssertExpectations(t)
}

func TestGetItem_WhenIdIsNegative_ReturnsError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	//ACT
	result, err := service.GetItem(-1)

	//ASSERT
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "O id não pode ser negativo.")
	mockRepo.AssertExpectations(t)
}

func TestGetItem_WhenRepositoryFails_ReturnsError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	mockRepo.On("GetItem", 1).Return((*entity.Item)(nil), assert.AnError)

	//ACT
	result, err := service.GetItem(1)

	//ASSERT
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Erro ao buscar o item")
	mockRepo.AssertExpectations(t)
}

func TestGetItem_WhenSuccess_ReturnsItem(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	expectedItem := &entity.Item{
		ID:      1,
		Nome:    "Produto Teste",
		Preco:   50,
		Estoque: 10,
		Status:  entity.StatusAtivo,
		Code:    "PR12345678",
	}

	mockRepo.On("GetItem", 1).Return(expectedItem, nil)

	//ACT
	result, err := service.GetItem(1)

	//ASSERT
	assert.NoError(t, err)
	assert.Equal(t, expectedItem, result)
	mockRepo.AssertExpectations(t)
}

func TestGetItens_WhenRepositoryFails_ReturnsError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	mockRepo.On("GetItens").Return(nil, assert.AnError)

	//ACT
	result, err := service.GetItens()

	//ASSERT
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Erro ao buscar os itens")
	mockRepo.AssertExpectations(t)
}

func TestGetItens_WhenSuccess_ReturnsItems(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	expectedItems := []entity.Item{
		{ID: 1, Nome: "Item 1", Preco: 10.0, Estoque: 5, Status: entity.StatusAtivo},
		{ID: 2, Nome: "Item 2", Preco: 20.0, Estoque: 0, Status: entity.StatusInativo},
	}

	mockRepo.On("GetItens").Return(expectedItems, nil)

	//ACT
	result, err := service.GetItens()

	//ASSERT
	assert.NoError(t, err)
	assert.Equal(t, expectedItems, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateItem_WhenPrecoInvalido_ReturnsError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	invalidItem := entity.Item{
		ID:      1,
		Nome:    "Produto Teste",
		Preco:   0,
		Estoque: 10,
	}

	//ACT
	err := service.UpdateItem(invalidItem)

	//ASSERT
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Preço deve ser maior que zero")
	mockRepo.AssertExpectations(t)
}

func TestUpdateItem_WhenEstoqueNegativo_ReturnsError(t *testing.T) {

	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	invalidItem := entity.Item{
		ID:      2,
		Nome:    "Produto Teste",
		Preco:   100.0,
		Estoque: -5,
	}

	//ACT

	err := service.UpdateItem(invalidItem)

	//ASSERT
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Estoque não pode ser negativo")
	mockRepo.AssertExpectations(t)
}

func TestUpdateItem_WhenEstoqueZero_SetsStatusInativo(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	item := entity.Item{
		ID:      1,
		Nome:    "Produto Teste",
		Preco:   100.0,
		Estoque: 0,
		Status:  entity.StatusAtivo,
	}

	mockRepo.On("UpdateItem", mock.MatchedBy(func(item entity.Item) bool {
		return item.Status == entity.StatusInativo
	})).Return(nil)

	//ACT
	err := service.UpdateItem(item)

	//ASSERT
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateItem_WhenEstoquePositivo_SetsStatusAtivo(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	item := entity.Item{
		ID:      1,
		Nome:    "Produto Teste",
		Preco:   100.0,
		Estoque: 10,
		Status:  entity.StatusInativo,
	}

	mockRepo.On("UpdateItem", mock.MatchedBy(func(item entity.Item) bool {
		return item.Status == entity.StatusAtivo
	})).Return(nil)

	//ACT
	err := service.UpdateItem(item)

	//ASSERT
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateItem_WhenRepositoryFails_ReturnError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	validItem := entity.Item{
		ID:      1,
		Nome:    "Produto Teste",
		Preco:   100.0,
		Estoque: 10,
	}

	mockRepo.On("UpdateItem", mock.Anything).Return(assert.AnError)

	//ACT
	err := service.UpdateItem(validItem)

	//ASSERT
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Erro ao atualizar o item")
	mockRepo.AssertExpectations(t)
}

func TestUpdateItem_WhenSucess_UpdateItens(t *testing.T) {
	// ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	validItem := entity.Item{
		ID:      1,
		Nome:    "Produto Teste",
		Preco:   100.0,
		Estoque: 15,
	}

	mockRepo.On("UpdateItem", mock.MatchedBy(func(item entity.Item) bool {
		return item.ID == 1 &&
			item.Nome == "Produto Teste" &&
			item.Preco == 100.0 &&
			item.Estoque == 15 &&
			item.Status == entity.StatusAtivo
	})).Return(nil)

	// ACT
	err := service.UpdateItem(validItem)

	// ASSERT
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateItem_WhenIdIsNegative_ReturnsError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	invalidItem := entity.Item{
		ID:      0,
		Nome:    "Produto Teste",
		Preco:   100.0,
		Estoque: 10,
	}

	//ACT
	err := service.UpdateItem(invalidItem)

	//ASSERT
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "O id deve ser maior que zero")
	mockRepo.AssertExpectations(t)
}

func TestDeleteItem_WhenIdIsNegative_ReturnError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	//ACT
	err := service.DeleteItem(-1)

	//ASSERT
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ID inválido para a exclusão")
	mockRepo.AssertExpectations(t)
}

func TestDeleteItem_WhenRepositoryFails_ReturnsError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	mockRepo.On("DeleteItem", 1).Return(assert.AnError)

	//ACT
	err := service.DeleteItem(1)

	//ASSERT
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Erro ao deletar item")
	mockRepo.AssertExpectations(t)
}

func TestDeleteItem_WhenSucess_DeleteItem(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	mockRepo.On("DeleteItem", 1).Return(nil)

	//ACT
	err := service.DeleteItem(1)

	//ASSERT
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
func TestGetItensPaginados_WhenSuccess_ReturnsItems(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	expectedItens := []entity.Item{
		{ID: 1, Nome: "Item 1", Preco: 10, Estoque: 5},
		{ID: 2, Nome: "Item 2", Preco: 20, Estoque: 10},
	}

	mockRepo.On("GetItensPaginados", 0, 10).Return(expectedItens, 2, nil)

	//ACT
	result, totalItens, err := service.GetItensPaginados(1, 10)

	//ASSERT
	assert.NoError(t, err)
	assert.Equal(t, expectedItens, result)
	assert.Equal(t, 2, totalItens)
	mockRepo.AssertExpectations(t)
}

func TestGetItensPaginados_WhenRepositoryFails_ReturnsError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	mockRepo.On("GetItensPaginados", 0, 10).Return(nil, 0, assert.AnError)

	//ACT
	result, TotalItens, err := service.GetItensPaginados(1, 10)

	//ASSERT
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, 0, TotalItens)
	assert.Contains(t, err.Error(), "Erro ao buscar itens paginados")
	mockRepo.AssertExpectations(t)
}

func TestGetItensPaginados_WhenInvalidParams_NormalizesValues(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	mockRepo.On("GetItensPaginados", 0, 10).Return([]entity.Item{}, 0, nil)

	//ACT
	_, _, err := service.GetItensPaginados(0, 0)

	//ASSERT
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetItensFiltradosPaginados_WhenSuccess_ReturnsItems(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	status := entity.StatusAtivo
	expectedItens := []entity.Item{
		{ID: 1, Nome: "Item 1", Status: entity.StatusAtivo},
	}

	mockRepo.On("GetItensFiltradosPaginados", &status, 0, 10).Return(expectedItens, 1, nil)

	//ACT
	result, totalItens, err := service.GetItensFiltradosPaginados(&status, 1, 10)

	//ASSERT
	assert.NoError(t, err)
	assert.Equal(t, expectedItens, result)
	assert.Equal(t, 1, totalItens)
	mockRepo.AssertExpectations(t)
}

func TestGetItensFiltradosPaginados_WhenRepositoryFails_ReturnsError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	status := entity.StatusAtivo
	mockRepo.On("GetItensFiltradosPaginados", &status, 0, 10).Return(nil, 0, assert.AnError)

	//ACT
	result, totalItens, err := service.GetItensFiltradosPaginados(&status, 1, 10)

	//ASSERT
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, 0, totalItens)
	mockRepo.AssertExpectations(t)
}

func TestGetItensFiltrados_WhenSuccess_ReturnsItems(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	status := entity.StatusAtivo
	expectedItems := []entity.Item{
		{ID: 1, Nome: "Item 1", Status: entity.StatusAtivo},
	}

	mockRepo.On("CountItens", &status).Return(1, nil)
	mockRepo.On("GetItensFiltrados", &status, 10).Return(expectedItems, nil)

	//ACT
	result, totalItens, totalPages, err := service.GetItensFiltrados(&status, 10)

	//ASSERT
	assert.NoError(t, err)
	assert.Equal(t, expectedItems, result)
	assert.Equal(t, 1, totalItens)
	assert.Equal(t, 1, totalPages)
	mockRepo.AssertExpectations(t)
}
func TestGetItensFiltrados_WhenCountFails_ReturnsError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	status := entity.StatusAtivo
	mockRepo.On("CountItens", &status).Return(0, assert.AnError)

	//ACT
	result, totalItens, totalPages, err := service.GetItensFiltrados(&status, 10)

	//ASSERT
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, 0, totalItens)
	assert.Equal(t, 0, totalPages)
	assert.Contains(t, err.Error(), "erro ao contar itens")
	mockRepo.AssertExpectations(t)
}

func TestGetItensFiltrados_WhenGetFiltradosFails_ReturnsError(t *testing.T) {
	//ARRANGE
	mockRepo := mocks.NewItemRepository(t)
	service := NewItemService(mockRepo)

	status := entity.StatusAtivo
	mockRepo.On("CountItens", &status).Return(5, nil)
	mockRepo.On("GetItensFiltrados", &status, 10).Return(nil, assert.AnError)

	//ACT
	result, totalItens, totalPages, err := service.GetItensFiltrados(&status, 10)

	//ASSERT
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, 5, totalItens)
	assert.Equal(t, 0, totalPages)
	assert.Contains(t, err.Error(), "erro ao buscar itens")
	mockRepo.AssertExpectations(t)
}
