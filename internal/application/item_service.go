package application

import (
	entity "desafio-itens-app/internal/domain/item" // Importa entidades do domínio
	"errors"                                        // Para criar erros simples
	"fmt"                                           // Para formatar erros
	"math"                                          // Para cálculos (Ceil)
)

type itemService struct { // Struct que implementa as regras de negócio
	repo entity.ItemRepository // Dependência: interface do repositório
}

type ItemService interface { // Interface que define operações do service
	GetItem(id int) (*entity.Item, error)
	AddItem(item entity.Item) (entity.Item, error)
	GetItens() ([]entity.Item, error)
	GetItensFiltrados(status *entity.Status, limit int) ([]entity.Item, int, int, error)
	UpdateItem(item entity.Item) error
	DeleteItem(id int) error
}

func NewItemService(repo entity.ItemRepository) *itemService { // Factory: cria nova instância do service
	return &itemService{ // Injeta dependência do repositório
		repo: repo,
	}
}

func (s *itemService) AddItem(item entity.Item) (entity.Item, error) {
	if item.Preco <= 0 { // Valida preço positivo
		return entity.Item{}, errors.New("O produto tem preço inválido.")
	}

	if item.Estoque < 0 { // Valida estoque não-negativo
		return entity.Item{}, errors.New("O produto tem estoque inválido.")
	}

	if item.Estoque == 0 { // Regra: sem estoque = inativo
		item.Status = entity.StatusInativo
	} else {
		item.Status = entity.StatusAtivo // Com estoque = ativo
	}

	code, err := s.generateUniqueCode(item.Nome) // Gera código único
	if err != nil {
		return entity.Item{}, err
	}
	item.Code = code // Atribui código gerado

	itemCriado, err := s.repo.AddItem(item) // Persiste no banco
	if err != nil {
		return entity.Item{}, err
	}

	return itemCriado, nil // Retorna item com ID do banco
}

func (s *itemService) GetItem(id int) (*entity.Item, error) {
	if id <= 0 {
		return nil, fmt.Errorf("O id não pode ser 0.")
	}

	item, err := s.repo.GetItem(id) // Busca no repositório
	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar o item: %w", err)
	}

	return item, nil // Retorna item encontrado
}

func (s *itemService) GetItens() ([]entity.Item, error) {
	itens, err := s.repo.GetItens() // Busca todos os itens
	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar os itens: %w", err)
	}

	return itens, nil // Retorna lista completa
}

func (s *itemService) generateUniqueCode(nome string) (string, error) {
	maxTentativas := 5

	for tentativa := 0; tentativa < maxTentativas; tentativa++ {
		code, err := entity.GenerateItemCode(nome)
		if err != nil {
			return "", err
		}

		exists, err := s.repo.CodeExists(code)
		if err != nil {
			return "", fmt.Errorf("erro ao verificar código: %w", err)
		}

		if !exists {
			return code, nil
		}
	}

	return "", errors.New("não foi possível gerar código único")
}

func (s *itemService) GetItensFiltrados(status *entity.Status, limit int) (itens []entity.Item, totalItens int, totalPages int, err error) {
	if limit <= 0 { // Normaliza limit mínimo
		limit = 10
	}

	if limit >= 20 { // Normaliza limit máximo
		limit = 20
	}

	totalItens, err = s.repo.CountItens(status) // Conta total para paginação
	if err != nil {
		err = fmt.Errorf("erro ao contar itens: %w", err)
		return // Named return
	}

	itens, err = s.repo.GetItensFiltrados(status, limit) // ❌ BUG: falta OFFSET para paginação
	if err != nil {
		err = fmt.Errorf("erro ao buscar itens: %w", err)
		return
	}

	totalPages = int(math.Ceil(float64(totalItens) / float64(limit))) // Calcula total de páginas
	return                                                            // Retorna tudo via named return
}

func (s *itemService) UpdateItem(item entity.Item) error {
	itemExistente, err := s.repo.GetItem(item.ID) // Verifica se item existe

	if err != nil {
		return fmt.Errorf("Erro ao buscar o item: %w", err)
	}

	if itemExistente == nil { // Se não encontrou
		return fmt.Errorf("Item não encontrado para atualização.")
	}

	if item.Preco <= 0 { // Valida preço
		return fmt.Errorf("O produto tem preço inválido.")
	}

	if item.Estoque < 0 { // Valida estoque
		return fmt.Errorf("O produto tem estoque inválido.")
	}

	if item.Estoque == 0 { // Recalcula status
		item.Status = entity.StatusInativo
	} else {
		item.Status = entity.StatusAtivo
	}

	if err := s.repo.UpdateItem(item); err != nil { // ❌ BUG: sobrescreve campos vazios
		return fmt.Errorf("Erro ao atualizar o item: %w", err)
	}
	return nil
}

func (s *itemService) DeleteItem(id int) error {
	if id <= 0 { // Valida ID positivo
		return fmt.Errorf("ID inválido para a exclusão: %d", id)
	}

	err := s.repo.DeleteItem(id) // Deleta do banco

	if err != nil { // ✅ CORRIGIDO: agora retorna erro
		return fmt.Errorf("Erro ao deletar item: %w", err)
	}

	return nil // Sucesso
}
