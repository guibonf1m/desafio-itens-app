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
	statusParam := c.Query("status") // Query parameter opcional
	limitParam := c.Query("limit")   // Query parameter opcional

	var status *entity.Status // Ponteiro para permitir nil
	if statusParam != "" {    // ⚙️ BUSINESS RULE: filtro condicional
		s := entity.Status(statusParam) // 🔄 TRANSFORMATION: string → Status
		status = &s
	}

	limit := 10                                         // Default limit
	if l, err := strconv.Atoi(limitParam); err == nil { // 🔄 TRANSFORMATION: string → int
		limit = l
	}

	itens, totalItens, totalPages, err := h.service.GetItensFiltrados(status, limit) // 🌐 EXTERNAL CALL
	if err != nil {                                                                  // 🛡️ VALIDATION GUARD
		c.JSON(http.StatusInternalServerError, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	resp := make([]ItemResponse, 0, len(itens)) // Pre-aloca slice para performance
	for _, it := range itens {
		resp = append(resp, FromEntity(it)) // 🔄 TRANSFORMATION: Entity → DTO
	}

	c.JSON(http.StatusOK, ResponseInfo{ // Resposta com paginação
		TotalItens: totalItens,
		TotalPages: totalPages,
		Data:       resp,
		Error:      false,
	})
}

func (h *ItemHandler) UpdateItem(c *gin.Context) {
	idParam := c.Param("id") // Parâmetro da URL

	id, err := strconv.Atoi(idParam) // 🔄 TRANSFORMATION: string → int
	if err != nil || id <= 0 {       // 🛡️ VALIDATION GUARD: ID válido
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	var req UpdateItemRequest                      // DTO para update
	if err := c.ShouldBindJSON(&req); err != nil { // 🔄 TRANSFORMATION: JSON → struct
		c.JSON(http.StatusBadRequest, ResponseInfo{ // 🛡️ VALIDATION GUARD
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	itemToUpdate := req.ToEntity(id) // 🔄 TRANSFORMATION: DTO → Entity

	if err := h.service.UpdateItem(itemToUpdate); err != nil { // 🌐 EXTERNAL CALL
		msg := err.Error()

		switch { // ⚙️ BUSINESS RULE: mapeia erros para HTTP status
		case strings.Contains(msg, "não encontrado"):
			c.JSON(404, ResponseInfo{Error: true, Result: msg})
		case strings.Contains(msg, "inválido"):
			c.JSON(400, ResponseInfo{Error: true, Result: msg})
		default:
			c.JSON(500, ResponseInfo{Error: true, Result: "Erro interno: " + msg})
		}
		return
	}

	c.JSON(http.StatusOK, ResponseInfo{ // Sucesso
		TotalPages: 1,
		Error:      false,
		Result:     "Item atualizado com sucesso!",
	})
}

func (h *ItemHandler) DeleteItem(c *gin.Context) {
	idParam := c.Param("id")         // Parâmetro da URL
	id, err := strconv.Atoi(idParam) // 🔄 TRANSFORMATION: string → int
	if err != nil || id <= 0 {       // 🛡️ VALIDATION GUARD
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: "ID inválido",
		})
		return
	}

	if err := h.service.DeleteItem(id); err != nil { // 🌐 EXTERNAL CALL
		msg := err.Error()
		switch { // ⚙️ BUSINESS RULE: mapeia erros
		case strings.Contains(msg, "Nenhum item encontrado"):
			c.JSON(http.StatusNotFound, ResponseInfo{
				Error:  true,
				Result: msg,
			})
		default: // ❌ BUG: sempre retorna sucesso
			c.JSON(http.StatusOK, ResponseInfo{
				Error:  false,
				Result: "Item deletado com sucesso!",
			})
		}
	}
}
