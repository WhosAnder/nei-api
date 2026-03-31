package handlers

import (
	"net/http"

	"github.com/WhosAnder/nei-api/internal/database"
	"github.com/WhosAnder/nei-api/internal/models"
	"github.com/gin-gonic/gin"
)

// ─── Categorías ───────────────────────────────────────────────────────────────

// GetCategorias godoc
// @Summary      Listar categorías
// @Description  Retorna todas las categorías del catálogo (Agrícola, Industrial)
// @Tags         categorias
// @Produce      json
// @Success      200  {array}   CategoriaResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /categorias [get]
func GetCategorias(c *gin.Context) {
	var categorias []models.Categoria
	if err := database.DB.Find(&categorias).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categorias)
}

// CreateCategoria godoc
// @Summary      Crear categoría (Admin)
// @Tags         categorias
// @Accept       json
// @Produce      json
// @Param        categoria  body      CategoriaResponse  true  "Datos de la categoría"
// @Success      201  {object}  CategoriaResponse
// @Failure      400  {object}  ErrorResponse
// @Security     BetterAuthSession
// @Router       /admin/categorias [post]
func CreateCategoria(c *gin.Context) {
	var cat models.Categoria
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&cat)
	c.JSON(http.StatusCreated, cat)
}

// UpdateCategoria godoc
// @Summary      Actualizar categoría (Admin)
// @Tags         categorias
// @Accept       json
// @Produce      json
// @Param        id         path      int               true  "ID de la categoría"
// @Param        categoria  body      CategoriaResponse  true  "Datos actualizados"
// @Success      200  {object}  CategoriaResponse
// @Failure      404  {object}  ErrorResponse
// @Security     BetterAuthSession
// @Router       /admin/categorias/{id} [put]
func UpdateCategoria(c *gin.Context) {
	var cat models.Categoria
	if err := database.DB.First(&cat, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoría no encontrada"})
		return
	}
	c.ShouldBindJSON(&cat)
	database.DB.Save(&cat)
	c.JSON(http.StatusOK, cat)
}

// DeleteCategoria godoc
// @Summary      Eliminar categoría (Admin)
// @Tags         categorias
// @Param        id  path  int  true  "ID de la categoría"
// @Success      200  {object}  MessageResponse
// @Security     BetterAuthSession
// @Router       /admin/categorias/{id} [delete]
func DeleteCategoria(c *gin.Context) {
	database.DB.Delete(&models.Categoria{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Eliminado"})
}

// ─── Maquinaria por Categoría ─────────────────────────────────────────────────

// GetMaquinariaByCat godoc
// @Summary      Maquinaria por categoría
// @Description  Retorna la lista de maquinaria asociada a una categoría por su slug
// @Tags         maquinaria
// @Produce      json
// @Param        id  path  string  true  "Slug de la categoría (ej: agricola, industrial)"
// @Success      200  {array}   MaquinariaResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /categorias/{id}/maquinaria [get]
func GetMaquinariaByCat(c *gin.Context) {
	var categoria models.Categoria
	if err := database.DB.Where("slug = ?", c.Param("id")).
		Preload("Maquinaria").First(&categoria).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoría no encontrada"})
		return
	}
	c.JSON(http.StatusOK, categoria.Maquinaria)
}

// CreateMaquinaria godoc
// @Summary      Crear maquinaria (Admin)
// @Tags         maquinaria
// @Accept       json
// @Produce      json
// @Param        maquinaria  body      MaquinariaResponse  true  "Datos de la maquinaria"
// @Success      201  {object}  MaquinariaResponse
// @Failure      400  {object}  ErrorResponse
// @Security     BetterAuthSession
// @Router       /admin/maquinaria [post]
func CreateMaquinaria(c *gin.Context) {
	var maq models.Maquinaria
	if err := c.ShouldBindJSON(&maq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&maq)
	c.JSON(http.StatusCreated, maq)
}

// UpdateMaquinaria godoc
// @Summary      Actualizar maquinaria (Admin)
// @Tags         maquinaria
// @Accept       json
// @Produce      json
// @Param        id          path  int               true  "ID"
// @Param        maquinaria  body  MaquinariaResponse true  "Datos actualizados"
// @Success      200  {object}  MaquinariaResponse
// @Failure      404  {object}  ErrorResponse
// @Security     BetterAuthSession
// @Router       /admin/maquinaria/{id} [put]
func UpdateMaquinaria(c *gin.Context) {
	var maq models.Maquinaria
	if err := database.DB.First(&maq, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Maquinaria no encontrada"})
		return
	}
	c.ShouldBindJSON(&maq)
	database.DB.Save(&maq)
	c.JSON(http.StatusOK, maq)
}

// DeleteMaquinaria godoc
// @Summary      Eliminar maquinaria (Admin)
// @Tags         maquinaria
// @Param        id  path  int  true  "ID"
// @Success      200  {object}  MessageResponse
// @Security     BetterAuthSession
// @Router       /admin/maquinaria/{id} [delete]
func DeleteMaquinaria(c *gin.Context) {
	database.DB.Delete(&models.Maquinaria{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Eliminado"})
}

// ─── Neumáticos por Maquinaria ────────────────────────────────────────────────

// GetNeumaticosByMaq godoc
// @Summary      Neumáticos por maquinaria
// @Description  Retorna los neumáticos disponibles para una máquina por su slug
// @Tags         neumaticos
// @Produce      json
// @Param        id  path  string  true  "Slug de la maquinaria (ej: tractor, retroexcavadora)"
// @Success      200  {array}   NeumaticoResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /maquinaria/{id}/neumaticos [get]
func GetNeumaticosByMaq(c *gin.Context) {
	var maq models.Maquinaria
	if err := database.DB.Where("slug = ?", c.Param("id")).
		Preload("Neumaticos.Marca").First(&maq).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Maquinaria no encontrada"})
		return
	}
	c.JSON(http.StatusOK, maq.Neumaticos)
}

// CreateNeumatico godoc
// @Summary      Crear neumático (Admin)
// @Tags         neumaticos
// @Accept       json
// @Produce      json
// @Param        neumatico  body      NeumaticoResponse  true  "Datos del neumático"
// @Success      201  {object}  NeumaticoResponse
// @Failure      400  {object}  ErrorResponse
// @Security     BetterAuthSession
// @Router       /admin/neumaticos [post]
func CreateNeumatico(c *gin.Context) {
	var neu models.Neumatico
	if err := c.ShouldBindJSON(&neu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&neu)
	c.JSON(http.StatusCreated, neu)
}

// UpdateNeumatico godoc
// @Summary      Actualizar neumático (Admin)
// @Tags         neumaticos
// @Accept       json
// @Produce      json
// @Param        id         path  int               true  "ID"
// @Param        neumatico  body  NeumaticoResponse  true  "Datos actualizados"
// @Success      200  {object}  NeumaticoResponse
// @Failure      404  {object}  ErrorResponse
// @Security     BetterAuthSession
// @Router       /admin/neumaticos/{id} [put]
func UpdateNeumatico(c *gin.Context) {
	var neu models.Neumatico
	if err := database.DB.First(&neu, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Neumático no encontrado"})
		return
	}
	c.ShouldBindJSON(&neu)
	database.DB.Save(&neu)
	c.JSON(http.StatusOK, neu)
}

// DeleteNeumatico godoc
// @Summary      Eliminar neumático (Admin)
// @Tags         neumaticos
// @Param        id  path  int  true  "ID"
// @Success      200  {object}  MessageResponse
// @Security     BetterAuthSession
// @Router       /admin/neumaticos/{id} [delete]
func DeleteNeumatico(c *gin.Context) {
	database.DB.Delete(&models.Neumatico{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Eliminado"})
}

// ─── Marcas ───────────────────────────────────────────────────────────────────

// GetMarcas godoc
// @Summary      Listar marcas
// @Description  Retorna todas las marcas de neumáticos disponibles
// @Tags         marcas
// @Produce      json
// @Success      200  {array}  MarcaResponse
// @Router       /marcas [get]
func GetMarcas(c *gin.Context) {
	var marcas []models.Marca
	database.DB.Find(&marcas)
	c.JSON(http.StatusOK, marcas)
}

// CreateMarca godoc
// @Summary      Crear marca (Admin)
// @Tags         marcas
// @Accept       json
// @Produce      json
// @Param        marca  body      MarcaResponse  true  "Datos de la marca"
// @Success      201  {object}  MarcaResponse
// @Failure      400  {object}  ErrorResponse
// @Security     BetterAuthSession
// @Router       /admin/marcas [post]
func CreateMarca(c *gin.Context) {
	var marca models.Marca
	if err := c.ShouldBindJSON(&marca); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&marca)
	c.JSON(http.StatusCreated, marca)
}

// UpdateMarca godoc
// @Summary      Actualizar marca (Admin)
// @Tags         marcas
// @Accept       json
// @Produce      json
// @Param        id     path  int           true  "ID"
// @Param        marca  body  MarcaResponse  true  "Datos actualizados"
// @Success      200  {object}  MarcaResponse
// @Failure      404  {object}  ErrorResponse
// @Security     BetterAuthSession
// @Router       /admin/marcas/{id} [put]
func UpdateMarca(c *gin.Context) {
	var marca models.Marca
	if err := database.DB.First(&marca, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Marca no encontrada"})
		return
	}
	c.ShouldBindJSON(&marca)
	database.DB.Save(&marca)
	c.JSON(http.StatusOK, marca)
}

// DeleteMarca godoc
// @Summary      Eliminar marca (Admin)
// @Tags         marcas
// @Param        id  path  int  true  "ID"
// @Success      200  {object}  MessageResponse
// @Security     BetterAuthSession
// @Router       /admin/marcas/{id} [delete]
func DeleteMarca(c *gin.Context) {
	database.DB.Delete(&models.Marca{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Eliminado"})
}
