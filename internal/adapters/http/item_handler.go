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

func NewItemHandler(service application.ItemService) *ItemHandler {
	return &ItemHandler{
		service: service,
	}
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

	//declarandoDTO -> Instância da struct que representa os dados do cliente
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

//
//func (h *ItemHandler) GetItens(c *gin.Context) {
//	// Captura os parâmetros de query 'page' e 'size', com valores padrão "1" e "10" se não forem fornecidos
//	pageStr := c.DefaultQuery("page", "1")
//	sizeStr := c.DefaultQuery("size", "10")
//
//	// Converte os parâmetros 'page' e 'size' de string para int
//	page, err := strconv.Atoi(pageStr)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, ResponseInfo{
//			Error:  true,
//			Result: "Número da página inválido.",
//		})
//		return
//	}
//
//	size, err := strconv.Atoi(sizeStr)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, ResponseInfo{
//			Error:  true,
//			Result: "Tamanho da página inválido",
//		})
//		return
//	}
//
//	// Uso do serviço para listar os itens com paginação
//	items, err := h.service.GetItens(page, size)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, ResponseInfo{
//			Error:  true,
//			Result: "Erro ao buscar itens",
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, ResponseInfo{
//		TotalPages: 1,
//		Data:       items,
//	})
//}

func (h *ItemHandler) UpdateItem(c *gin.Context) {

	idParam := c.Param("id")
	var item entity.Item

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	item.ID = id
	produtoAtualizado, err := h.service.UpdateItem(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}
//
//	produtoResponse := h.service.NovoProdutoResponse(produtoAtualizado)
//	c.JSON(http.StatusOK, ResponseInfo{
//		Error:  false,
//		Result: produtoResponse,
//	})
//
//}

//func (h *ItemHandler) DeleteItem(c *gin.Context) {
//
//		idParam := c.Param("id")
//
//		id, err := strconv.Atoi(idParam)
//		if err != nil {
//			c.JSON(http.StatusBadRequest, ResponseInfo{
//				Error:  true,
//				Result: err.Error(),
//			})
//			return
//		}
//
//		h.service.DeleteProduto(id)
//
//		c.JSON(http.StatusOK, ResponseInfo{
//			Error:  false,
//			Result: "deletado com sucesso",
//		})
//	}
