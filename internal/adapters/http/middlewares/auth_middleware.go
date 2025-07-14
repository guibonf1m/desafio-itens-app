package middlewares

import (
	"desafio-itens-app/internal/adapters/http/auth"
	"desafio-itens-app/internal/adapters/http/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	jwtService *auth.JWTService
}

func NewAuthMiddleware(jwtService *auth.JWTService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, handler.ResponseInfo{
				Error:  true,
				Result: "Token de autorização é obrigatório",
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, handler.ResponseInfo{
				Error:  true,
				Result: "Formato do token inválido. Use: Bearer <token>",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		claims, err := m.jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, handler.ResponseInfo{
				Error:  true,
				Result: "Token inválido ou expirado",
			})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("userRole", claims.Role)

		c.Next()
	}
}

// RequireRole - segundo segurança que verifica o "cargo"
func (m *AuthMiddleware) RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusForbidden, handler.ResponseInfo{
				Error:  true,
				Result: "Acesso negado: role não encontrado",
			})
			c.Abort()
			return
		}

		roleStr, ok := userRole.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, handler.ResponseInfo{
				Error:  true,
				Result: "Erro interno: role inválido",
			})
			c.Abort()
			return
		}

		hasPermission := false
		for _, allowedRole := range allowedRoles {
			if roleStr == allowedRole {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, handler.ResponseInfo{
				Error:  true,
				Result: "Acesso negado: permissão insuficiente",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
