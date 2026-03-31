// @title           NEI API
// @version         1.0
// @description     API REST para gestión del catálogo de neumáticos agrícolas e industriales.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Soporte NEI
// @contact.email  contacto@nei.com

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BetterAuthSession
// @in cookie
// @name better-auth.session_token

package main

import (
	"log"
	"os"

	_ "github.com/WhosAnder/nei-api/docs" // generado por swag
	"github.com/WhosAnder/nei-api/internal/database"
	"github.com/WhosAnder/nei-api/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// En desarrollo local cargamos el .env.
	// En Railway (production/staging) las variables ya están inyectadas,
	// godotenv no se ejecuta para no sobreescribir nada.
	appEnv := os.Getenv("APP_ENV")
	if appEnv != "production" && appEnv != "staging" {
		if err := godotenv.Load(); err != nil {
			log.Println("⚠️  .env no encontrado — asegúrate de copiar .env.example a .env")
		} else {
			log.Println("✅ Variables cargadas desde .env (modo local)")
		}
	} else {
		log.Printf("☁️  Entorno %s — usando variables inyectadas por Railway", appEnv)
	}

	// Conectar y migrar base de datos
	database.Connect()
	database.Migrate()

	// Configurar modo Gin
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Crear servidor
	r := gin.Default()

	// Configurar CORS para permitir peticiones desde Next.js
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Cookie")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "app": "nei-api"})
	})

	// Registrar rutas
	router.Setup(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 NEI API corriendo en http://localhost:%s", port)
	r.Run(":" + port)
}
