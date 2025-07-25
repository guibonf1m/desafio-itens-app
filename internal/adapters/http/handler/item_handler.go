package handler

import (
	"desafio-itens-app/internal/adapters/http/dto"
	"desafio-itens-app/internal/application/ports/services"
	entity "desafio-itens-app/internal/domain/item" // Domain entities
	"fmt"
	"github.com/gin-gonic/gin" // HTTP framework
	"net/http"                 // HTTP status codes
	"strconv"                  // String conversions
)

type ResponseInfo struct { // Padronização de resposta HTTP
	TotalItens int         `json:"totalItens,omitempty"`
	TotalPages int         `json:"totalPages,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Error      bool        `json:"error,omitempty"`
	Result     any         `json:"result,omitempty"`
}
type PageInfo struct {
	Page       int `json:"pagina"`
	PageSize   int `json:"tamanhoPagina"`
	TotalItems int `json:"totalItens"`
	TotalPages int `json:"totalPaginas"`
}

type ItemHandler struct { // Handler para operações de Item
	service services.ItemService // Dependência: service layer
}

func NewItemHandler(service services.ItemService) *ItemHandler { // Factory function
	return &ItemHandler{service: service} // Injeta dependência
}

func (h *ItemHandler) AddItem(c *gin.Context) {
	// PASSO 1: EXTRAIR userID do context
	userID, exists := c.Get("userID")

	// 🔍 DEBUG - adicionar logs
	fmt.Printf("🔍 DEBUG - userID exists: %v\n", exists)
	fmt.Printf("🔍 DEBUG - userID value: %v\n", userID)

	if !exists {
		c.JSON(http.StatusUnauthorized, ResponseInfo{
			Error:  true,
			Result: "Usuário não autenticado",
		})
		return
	}

	// PASSO 2: CONVERTER para int
	userIDInt, ok := userID.(int)
	fmt.Printf("🔍 DEBUG - userIDInt: %v, ok: %v\n", userIDInt, ok) // ← Mais um log

	if !ok {
		c.JSON(http.StatusInternalServerError, ResponseInfo{
			Error:  true,
			Result: "Erro interno: userID inválido",
		})
		return
	}

	// PASSO 3: RECEBER e VALIDAR JSON
	var req dto.CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("erro: %v", err)
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	// PASSO 4: CONVERTER para Entity e DEFINIR auditoria
	item := req.ToEntity()
	item.CreatedBy = &userIDInt
	fmt.Printf("🔍 DEBUG - item.CreatedBy: %v\n", item.CreatedBy) // ← Mais um log

	// PASSO 5: CHAMAR Service
	createdItem, err := h.service.AddItem(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	// PASSO 6: RETORNAR resposta
	c.JSON(http.StatusCreated, ResponseInfo{
		Error:  false,
		Result: dto.FromEntity(createdItem),
	})
}

func (h *ItemHandler) GetItem(c *gin.Context) {
	idParam := c.Param("id") // Extrai parâmetro da URL

	id, err := strconv.Atoi(idParam) // 🔄 TRANSFORMATION: string → int
	if err != nil {                  // 🛡️ VALIDATION GUARD
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: "o parametro não é um número, tente novamente.",
		})
		return
	}

	item, err := h.service.GetItem(id) // 🌐 EXTERNAL CALL: busca no service
	if item == nil || item.ID == 0 {   // ⚙️ BUSINESS RULE: verifica se encontrou
		c.JSON(http.StatusNotFound, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseInfo{ // Retorna item encontrado
		TotalPages: 1,
		Data:       item,
	})
}

func (h *ItemHandler) GetItens(c *gin.Context) {
	// 🔍 PARÂMETROS DE FILTRO (já existentes)
	statusParam := c.Query("status") // ?status=active

	// 📄 PARÂMETROS DE PAGINAÇÃO (novos)
	pageParam := c.DefaultQuery("page", "1")          // ?page=2
	pageSizeParam := c.DefaultQuery("pageSize", "10") // ?pageSize=5

	// 🔄 PROCESSAR filtro de status (já existente)
	var status *entity.Status
	if statusParam != "" {
		s := entity.Status(statusParam)
		status = &s
	}

	// 🔄 PROCESSAR parâmetros de paginação
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeParam)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100 // Limite máximo
	}

	// 📞 CHAMAR Service com paginação E filtros
	itens, totalItens, err := h.service.GetItensFiltradosPaginados(status, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	// 🔄 TRANSFORMAR para Response (igual ao seu código)
	resp := make([]dto.ItemResponse, 0, len(itens))
	for _, it := range itens {
		resp = append(resp, dto.FromEntity(it))
	}

	// 🧮 CALCULAR total de páginas
	totalPages := (totalItens + pageSize - 1) / pageSize

	// ✅ RESPOSTA com sua struct ResponseInfo + paginação
	c.JSON(http.StatusOK, ResponseInfo{
		TotalItens: totalItens,
		TotalPages: totalPages,
		Data:       resp,
		Error:      false,
	})
}

func (h *ItemHandler) UpdateItem(c *gin.Context) {
	// PASSO 1: EXTRAIR userID do context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, ResponseInfo{
			Error:  true,
			Result: "Usuário não autenticado",
		})
		return
	}

	// PASSO 2: CONVERTER para int
	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, ResponseInfo{
			Error:  true,
			Result: "Erro interno: userID inválido",
		})
		return
	}

	// PASSO 3: PEGAR ID da URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: "ID inválido",
		})
		return
	}

	// PASSO 4: RECEBER e VALIDAR JSON
	var req dto.UpdateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	// PASSO 5: BUSCAR item existente
	existingItem, err := h.service.GetItem(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	// PASSO 6: VERIFICAR AUTORIZAÇÃO
	userRole, exists := c.Get("userRole")
	if !exists {
		c.JSON(http.StatusForbidden, ResponseInfo{
			Error:  true,
			Result: "Role não encontrado",
		})
		return
	}

	roleStr := userRole.(string)

	if roleStr != "admin" {
		if existingItem.CreatedBy == nil || *existingItem.CreatedBy != userIDInt {
			c.JSON(http.StatusForbidden, ResponseInfo{
				Error:  true,
				Result: "Você só pode editar itens que criou",
			})
			return
		}
	}

	// PASSO 6: APLICAR mudanças e DEFINIR auditoria
	updatedItem := *existingItem
	req.ApplyTo(&updatedItem)         // ← Usando SEU método
	updatedItem.UpdateBy = &userIDInt // ← AUDITORIA: quem atualizou

	// PASSO 7: CHAMAR Service
	err = h.service.UpdateItem(updatedItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	// PASSO 8: RETORNAR resposta
	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: dto.FromEntity(updatedItem),
	})
}

func (h *ItemHandler) DeleteItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: "ID inválido",
		})
		return
	}

	// 🔑 CORREÇÃO: Lógica simples e clara
	err = h.service.DeleteItem(id)
	if err != nil {
		// ✅ QUALQUER erro = resposta de erro
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(), // "item com ID 999 não encontrado"
		})
		return
	}

	// ✅ SÓ chega aqui se NÃO teve erro
	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: "Item deletado com sucesso!",
	})
}
