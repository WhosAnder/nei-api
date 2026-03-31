package router

import (
	"github.com/WhosAnder/nei-api/internal/handlers"
	"github.com/WhosAnder/nei-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	// ── Rutas Públicas (lectura del catálogo) ──────────────────────────────────
	v1 := r.Group("/api/v1")
	{
		v1.GET("/categorias", handlers.GetCategorias)
		v1.GET("/categorias/:id/maquinaria", handlers.GetMaquinariaByCat)
		v1.GET("/maquinaria/:id/neumaticos", handlers.GetNeumaticosByMaq)
		v1.GET("/marcas", handlers.GetMarcas)
	}

	// ── Rutas Admin (requieren sesión Better Auth válida) ──────────────────────
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AuthRequired())
	{
		// Categorías
		admin.POST("/categorias", handlers.CreateCategoria)
		admin.PUT("/categorias/:id", handlers.UpdateCategoria)
		admin.DELETE("/categorias/:id", handlers.DeleteCategoria)

		// Maquinaria
		admin.POST("/maquinaria", handlers.CreateMaquinaria)
		admin.PUT("/maquinaria/:id", handlers.UpdateMaquinaria)
		admin.DELETE("/maquinaria/:id", handlers.DeleteMaquinaria)

		// Neumáticos
		admin.POST("/neumaticos", handlers.CreateNeumatico)
		admin.PUT("/neumaticos/:id", handlers.UpdateNeumatico)
		admin.DELETE("/neumaticos/:id", handlers.DeleteNeumatico)

		// Marcas
		admin.POST("/marcas", handlers.CreateMarca)
		admin.PUT("/marcas/:id", handlers.UpdateMarca)
		admin.DELETE("/marcas/:id", handlers.DeleteMarca)
	}
}
