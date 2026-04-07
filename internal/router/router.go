package router

import (
	"github.com/WhosAnder/nei-api/internal/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(r *gin.Engine) {
	// ── Swagger UI (/swagger/index.html) ───────────────────────────────────────
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// ── Rutas Públicas (lectura del catálogo) ──────────────────────────────────
	v1 := r.Group("/api/v1")
	{
		v1.GET("/categorias", handlers.GetCategorias)
		v1.GET("/categorias/:id/maquinaria", handlers.GetMaquinariaByCat)
		v1.GET("/maquinaria/:id/neumaticos", handlers.GetNeumaticosByMaq)
		v1.GET("/marcas", handlers.GetMarcas)
		v1.GET("/servicios", handlers.GetServicios)
	}

	// ── Rutas Admin (requieren sesión Better Auth válida) ──────────────────────
	admin := r.Group("/api/v1/admin")
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
