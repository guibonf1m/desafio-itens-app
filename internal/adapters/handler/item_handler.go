package handler

import (
	"desafio-itens-app/internal/application"
	"desafio-itens-app/internal/domain"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseInfo struct {
	TotalPages int         `json:"totalPages"`
	Data       interface{} `json:"data"`
	Error      bool        `json:"error"`
	Result     any         `json:"result"`
}
type ItemHandler struct {
	service *application.ItemService
}

func NovoItemResponse(item domain.Item) ItemResponse {

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
	var item domain.Item

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
