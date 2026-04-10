package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthRequired valida que la petición venga con una sesión activa de Better Auth
// preguntándole directamente al servidor de Next.js.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		appEnv := os.Getenv("APP_ENV")
		if appEnv != "production" && appEnv != "staging" {
			c.Set("user", gin.H{"dev": true})
			c.Next()
			return
		}

		cookieHeader := c.Request.Header.Get("Cookie")
		if cookieHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Sin sesión activa"})
			return
		}

		// Consultar a Next.js si la sesión es válida
		nextjsURL := os.Getenv("NEXTJS_URL") // ej: http://localhost:3000
		if nextjsURL == "" {
			nextjsURL = "http://127.0.0.1:3000"
		}

		requestURLs := []string{nextjsURL + "/api/auth/get-session"}
		if strings.Contains(nextjsURL, "localhost") {
			requestURLs = append(requestURLs, strings.Replace(nextjsURL, "localhost", "127.0.0.1", 1)+"/api/auth/get-session")
		}

		var resp *http.Response
		var err error
		for _, url := range requestURLs {
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("Cookie", cookieHeader)
			resp, err = http.DefaultClient.Do(req)
			if err == nil {
				break
			}
		}

		if err != nil || resp.StatusCode != http.StatusOK {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Sesión inválida"})
			return
		}
		defer resp.Body.Close()

		var session map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&session); err != nil || session["user"] == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
			return
		}

		c.Set("user", session["user"])
		c.Next()
	}
}
