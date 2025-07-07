package http

import (
	"desafio-itens-app/internal/application"        // Service layer
	entity "desafio-itens-app/internal/domain/item" // Domain entities
	"encoding/json"                                 // JSON parsing
	"github.com/gin-gonic/gin"                      // HTTP framework
	"net/http"                                      // HTTP status codes
	"strconv"                                       // String conversions
	"strings"                                       // String manipulation
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
	service application.ItemService // Dependência: service layer
}

func NewItemHandler(service application.ItemService) *ItemHandler { // Factory function
	return &ItemHandler{service: service} // Injeta dependência
}

func (h *ItemHandler) AddItem(c *gin.Context) {
	var req CreateItemRequest // DTO para request

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil { // 🔄 TRANSFORMATION: JSON → struct
		c.JSON(http.StatusBadRequest, ResponseInfo{ // 🛡️ VALIDATION GUARD
			Error:  true,
			Result: "Erro ao decodificar JSON",
		})
		return
	}

	item := req.ToEntity() // 🔄 TRANSFORMATION: DTO → Entity

	itemCriado, err := h.service.AddItem(item) // 🌐 EXTERNAL CALL
	if err != nil {
		msg := err.Error()

		switch { // ⚙️ BUSINESS RULE: mapeia erro → status HTTP
		case strings.Contains(msg, "já existe"): // 409 - Recurso duplicado
			c.JSON(http.StatusConflict, ResponseInfo{
				Error:  true,
				Result: msg,
			})
		case strings.Contains(msg, "inválido"): // 400 - Dados inválidos
			c.JSON(http.StatusBadRequest, ResponseInfo{
				Error:  true,
				Result: msg,
			})
		default: // 500 - Erro interno
			c.JSON(http.StatusInternalServerError, ResponseInfo{
				Error:  true,
				Result: "Erro interno: " + msg,
			})
		}
		return
	}

	itemResponse := FromEntity(itemCriado) // 🔄 TRANSFORMATION: Entity → DTO

	c.JSON(http.StatusCreated, ResponseInfo{ // 201 - Recurso criado
		TotalPages: 1,
		Data:       []ItemResponse{itemResponse},
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
	resp := make([]ItemResponse, 0, len(itens))
	for _, it := range itens {
		resp = append(resp, FromEntity(it))
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
	// 🔄 TRANSFORMATION: string → int
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: "Id inválido",
		})
		return
	}

	// 🔄 TRANSFORMATION: JSON → DTO
	var req UpdateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	existingItem, err := h.service.GetItem(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ResponseInfo{
			Error:  true,
			Result: "Item não encontrado",
		})
		return
	}

	// 🔄 APLICA MUDANÇAS DE FORMA SEGURA
	req.ApplyTo(existingItem)

	// 💾 SALVAR no banco
	if err := h.service.UpdateItem(*existingItem); err != nil {
		c.JSON(http.StatusInternalServerError, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	// ✅ RETORNAR item atualizado
	c.JSON(http.StatusOK, ResponseInfo{
		TotalPages: 1,
		Error:      false,
		Result:     "Item atualizado com sucesso!",
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
