package handler

import (
	"desafio-itens-app/internal/adapters/http/auth"
	"desafio-itens-app/internal/adapters/http/dto"
	"desafio-itens-app/internal/application/ports/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	service    services.UserService // ← Dependência: UserService
	jwtService *auth.JWTService
}

// NewUserHandler - Factory function (cria instância do handler)
func NewUserHandler(service services.UserService, jwtService *auth.JWTService) *UserHandler {
	return &UserHandler{
		service:    service,
		jwtService: jwtService, // ← Injetar dependência
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	// PASSO 1: RECEBER e VALIDAR JSON
	var req dto.CreateUserRequest
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
		Result: dto.FromUserEntity(createdUser), // ← Entity → DTO Response (sem senha)
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
		Result: dto.FromUserEntity(*user), //*user porque Service retorna ponteiro
	})

}

func (h *UserHandler) ListUsers(c *gin.Context) {
	// PASSO 1: Pegar parâmetros de paginação
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// PASSO 2: Chamar service
	result, err := h.service.ListUsers(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	// PASSO 3: Retornar lista
	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: result,
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
		Result: dto.FromUserEntity(*user),
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

	var req dto.UpdateUserRequest
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
		Result: dto.FromUserEntity(updateUser),
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

func (h *UserHandler) Register(c *gin.Context) {
	var req dto.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	user := req.ToEntity()

	createdUser, err := h.service.CreateUser(user)
	if err != nil {
		if err.Error() == "username já está em uso" {
			c.JSON(http.StatusBadRequest, ResponseInfo{
				Error:  true,
				Result: "username já existe",
			})
			return
		}
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, ResponseInfo{
		Error:  false,
		Result: dto.FromUserEntity(createdUser),
	})

}

func (h *UserHandler) Login(c *gin.Context) {

	var req dto.LoginRequest
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

	token, err := h.jwtService.GenerateToken(user.ID, user.Username, string(user.Role))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseInfo{
			Error:  true,
			Result: "erro ao gerar o token",
		})
		return
	}

	LoginResponse := dto.LoginResponse{
		Token:     token,
		ExpiresIn: 3600,
		User:      dto.FromUserEntity(*user),
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: LoginResponse,
	})

}

func (h *UserHandler) ValidateCredentials(c *gin.Context) {
	var req dto.LoginRequest

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

	token, err := h.jwtService.GenerateToken(user.ID, user.Username, string(user.Role))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseInfo{
			Error:  true,
			Result: "erro ao gerar o token",
		})
		return
	}

	LoginResponse := dto.LoginResponse{
		Token:     token,
		ExpiresIn: 3600,
		User:      dto.FromUserEntity(*user),
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: LoginResponse,
	})
}
