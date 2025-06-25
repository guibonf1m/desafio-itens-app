package http

import (
	"desafio-itens-app/internal/application"
	entity "desafio-itens-app/internal/domain/item"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResponseInfo struct {
	TotalPages int         `json:"totalPages,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Error      bool        `json:"error,omitempty"`
	Result     any         `json:"result,omitempty"`
}
type ItemHandler struct {
	service application.ItemService
}

func NovoItemResponse(item entity.Item) ItemResponse {

	respostaItem := ItemResponse{
		ID:        item.ID,
		Code:      item.Code,
		Nome:      item.Nome,
		Descricao: item.Descricao,
		Preco:     item.Preco,
		Estoque:   item.Estoque,
		Status:    item.Status,
	}
	return respostaItem
}

func (h *ItemHandler) AddItem(c *gin.Context) {
	var item entity.Item

	err := json.NewDecoder(c.Request.Body).Decode(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	itemCriado, err := h.service.AddItem(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	itemResponse := NovoItemResponse(itemCriado)

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
	if item.ID == 0 {
		c.JSON(http.StatusNotFound, ResponseInfo{
			Error:  true,
			Result: "produto não existe, tente novamente.",
		})
		return
	}

	c.JSON(http.StatusOK, ResponseInfo{
		TotalPages: 1,
		Data:       item,
	})
}

func (h *ItemHandler) GetItens(c *gin.Context) {
	// Captura os parâmetros de query 'page' e 'size', com valores padrão "1" e "10" se não forem fornecidos
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "10")

	// Converte os parâmetros 'page' e 'size' de string para int
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: "Número da página inválido.",
		})
		return
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: "Tamanho da página inválido",
		})
		return
	}

	// Uso do serviço para listar os itens com paginação
	items, err := h.service.GetItens(page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseInfo{
			Error:  true,
			Result: "Erro ao buscar itens",
		})
		return
	}

	c.JSON(http.StatusOK, ResponseInfo{
		TotalPages: 1,
		Data:       items,
	})
}

func (h *ItemHandler) UpdateItem(c *gin.Context) {

}

func (h *ItemHandler) DeleteItem(c *gin.Context) {

}
