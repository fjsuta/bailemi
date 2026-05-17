package middleware

import (
	"strings"

	"github.com/bailemi/gateway/pkg/errors"
	"github.com/bailemi/gateway/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtManager *jwt.JWTManager
}

func NewAuthMiddleware(jwtManager *jwt.JWTManager) *AuthMiddleware {
	return &AuthMiddleware{jwtManager: jwtManager}
}

func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			errors.Error(c, 401, errors.ErrUnauthorized.Code, errors.ErrUnauthorized.Message)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			errors.Error(c, 401, errors.ErrInvalidToken.Code, errors.ErrInvalidToken.Message)
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := m.jwtManager.ValidateAccessToken(tokenString)
		if err != nil {
			errors.Error(c, 401, errors.ErrInvalidToken.Code, errors.ErrInvalidToken.Message)
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func (m *AuthMiddleware) RequireRole(minRole uint8) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			errors.Error(c, 403, errors.ErrUnauthorized.Code, errors.ErrUnauthorized.Message)
			c.Abort()
			return
		}

		userRole := role.(uint8)
		if userRole < minRole {
			errors.Error(c, 403, errors.ErrNoPermission.Code, errors.ErrNoPermission.Message)
			c.Abort()
			return
		}

		c.Next()
	}
}

func (m *AuthMiddleware) OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.Next()
			return
		}

		tokenString := parts[1]
		claims, err := m.jwtManager.ValidateAccessToken(tokenString)
		if err == nil {
			c.Set("user_id", claims.UserID)
			c.Set("username", claims.Username)
			c.Set("role", claims.Role)
		}

		c.Next()
	}
}
