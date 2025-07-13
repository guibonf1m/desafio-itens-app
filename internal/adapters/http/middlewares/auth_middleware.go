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

		c.Next()
	}
}
