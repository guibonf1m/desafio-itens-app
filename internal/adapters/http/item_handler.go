package http

import (
	"desafio-itens-app/internal/application"
	entity "desafio-itens-app/internal/domain/item"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ResponseInfo struct {
	TotalItens int         `json:"totalItens,omitempty"`
	TotalPages int         `json:"totalPages,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Error      bool        `json:"error,omitempty"`
	Result     any         `json:"result,omitempty"`
}
type ItemHandler struct {
	service application.ItemService
}

func NewItemHandler(service application.ItemService) *ItemHandler {
	return &ItemHandler{
		service: service,
	}
}

func (h *ItemHandler) AddItem(c *gin.Context) {

	var req CreateItemRequest

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {

		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: "Erro ao decodificar JSON",
		})
		return
	}

	item := req.ToEntity()

	itemCriado, err := h.service.AddItem(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	itemResponse := FromEntity(itemCriado)

	c.JSON(http.StatusCreated, ResponseInfo{
		TotalPages: 1,
		Data:       []ItemResponse{itemResponse},
	})
}

func (h *ItemHandler) GetItem(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: "o parametro não é um número, tente novamente.",
		})
		return
	}

	item, err := h.service.GetItem(id)
	if item == nil || item.ID == 0 {
		c.JSON(http.StatusNotFound, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseInfo{
		TotalPages: 1,
		Data:       item,
	})
}

func (h *ItemHandler) GetItens(c *gin.Context) {
	// lê query params…
	statusParam := c.Query("status")
	limitParam := c.Query("limit")

	var status *entity.Status
	if statusParam != "" {
		s := entity.Status(statusParam)
		status = &s
	}

	limit := 10
	if l, err := strconv.Atoi(limitParam); err == nil {
		limit = l
	}

	// **chamada UNIFICADA** que retorna 4 valores
	itens, totalItens, totalPages, err := h.service.GetItensFiltrados(status, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	// mapeia e devolve
	resp := make([]ItemResponse, 0, len(itens))
	for _, it := range itens {
		resp = append(resp, FromEntity(it))
	}

	c.JSON(http.StatusOK, ResponseInfo{
		TotalItens: totalItens,
		TotalPages: totalPages,
		Data:       resp,
		Error:      false,
	})
}

func (h *ItemHandler) UpdateItem(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	var req UpdateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	itemToUpdate := req.ToEntity(id)

	if err := h.service.UpdateItem(itemToUpdate); err != nil {
		msg := err.Error()

		switch {
		case strings.Contains(msg, "não encontrado"):
			c.JSON(404, ResponseInfo{Error: true, Result: msg})
		case strings.Contains(msg, "inválido"):
			c.JSON(400, ResponseInfo{Error: true, Result: msg})
		default:
			c.JSON(500, ResponseInfo{Error: true, Result: "Erro interno: " + msg})
		}
		return
	}

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

	if err := h.service.DeleteItem(id); err != nil {
		msg := err.Error()
		switch {
		case strings.Contains(msg, "Nenhum item encontrado"):
			c.JSON(http.StatusNotFound, ResponseInfo{
				Error:  true,
				Result: msg,
			})
		default:
			c.JSON(http.StatusOK, ResponseInfo{
				Error:  false,
				Result: "Item deletado com sucesso!",
			})
		}
	}
}
