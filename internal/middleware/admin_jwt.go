package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AdminClaims struct {
	Role  string `json:"role"`
	Email string `json:"email,omitempty"`
	jwt.RegisteredClaims
}

func AdminJWTRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		appEnv := os.Getenv("APP_ENV")
		if appEnv != "production" && appEnv != "staging" {
			c.Set("user", gin.H{"dev": true, "role": "admin"})
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token no enviado"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		secret := os.Getenv("ADMIN_API_JWT_SECRET")
		if secret == "" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "ADMIN_API_JWT_SECRET no configurado"})
			return
		}

		claims := &AdminClaims{}
		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			},
			jwt.WithValidMethods([]string{"HS256"}),
			jwt.WithIssuer("next-admin"),
			jwt.WithAudience("nei-api-admin"),
			jwt.WithLeeway(15*time.Second),
		)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Sesión inválida"})
			return
		}

		if claims.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Sin permisos"})
			return
		}

		c.Set("user", gin.H{
			"id":    claims.Subject,
			"email": claims.Email,
			"role":  claims.Role,
		})
		c.Next()
	}
}
