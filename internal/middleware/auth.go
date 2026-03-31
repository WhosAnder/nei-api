package middleware

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// AuthRequired valida que la petición venga con una sesión activa de Better Auth
// preguntándole directamente al servidor de Next.js.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Leer cookie de sesión enviada por el navegador / admin panel
		cookie, err := c.Cookie("better-auth.session_token")
		if err != nil || cookie == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Sin sesión activa"})
			return
		}

		// Consultar a Next.js si la sesión es válida
		nextjsURL := os.Getenv("NEXTJS_URL") // ej: http://localhost:3000
		req, _ := http.NewRequest("GET", nextjsURL+"/api/auth/get-session", nil)
		req.Header.Set("Cookie", "better-auth.session_token="+cookie)

		resp, err := http.DefaultClient.Do(req)
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
