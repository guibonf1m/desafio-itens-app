package http

import (
	"desafio-itens-app/internal/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	service application.UserService // ← Dependência: UserService
}

// NewUserHandler - Factory function (cria instância do handler)
func NewUserHandler(service application.UserService) *UserHandler {
	return &UserHandler{service: service} // ← Injeta dependência
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	// PASSO 1: RECEBER e VALIDAR JSON
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	// PASSO 2: CONVERTER DTO → Entity
	user := req.ToEntity()

	//PASSO 3: CHAMAR Service (toda lógica está lá)
	createdUser, err := h.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, ResponseInfo{
		Error:  false,
		Result: FromUserEntity(createdUser), // ← Entity → DTO Response (sem senha)
	})

}

func (h *UserHandler) GetUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil && id <= 0 {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: "ID inválido",
		})
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: FromUserEntity(*user), //*user porque Service retorna ponteiro
	})

}

func (h *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: "username é obrigatório",
		})
		return
	}

	user, err := h.service.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: FromUserEntity(*user),
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: "ID inválido",
		})
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	existingUser, err := h.service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	updateUser := *existingUser
	req.ApplyTo(&updateUser)

	err = h.service.UpdateUser(updateUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: FromUserEntity(updateUser),
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: "ID Inválido",
		})
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: "usuário deletado com sucesso",
	})

}

func (h *UserHandler) ValidateCredentials(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	user, err := h.service.ValidateCredentials(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error: false,
		Result: map[string]interface{}{
			"message": "login realizado com sucesso",
			"user":    FromUserEntity(*user),
		},
	})
}
