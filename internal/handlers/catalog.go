package handlers

import (
	"net/http"

	"github.com/WhosAnder/nei-api/internal/database"
	"github.com/WhosAnder/nei-api/internal/models"
	"github.com/gin-gonic/gin"
)

// ─── Categorías ───────────────────────────────────────────────────────────────

func GetCategorias(c *gin.Context) {
	var categorias []models.Categoria
	if err := database.DB.Find(&categorias).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categorias)
}

func CreateCategoria(c *gin.Context) {
	var cat models.Categoria
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&cat)
	c.JSON(http.StatusCreated, cat)
}

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

func DeleteCategoria(c *gin.Context) {
	database.DB.Delete(&models.Categoria{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Eliminado"})
}

// ─── Maquinaria por Categoría ─────────────────────────────────────────────────

func GetMaquinariaByCat(c *gin.Context) {
	var categoria models.Categoria
	if err := database.DB.Where("slug = ?", c.Param("id")).
		Preload("Maquinaria").First(&categoria).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoría no encontrada"})
		return
	}
	c.JSON(http.StatusOK, categoria.Maquinaria)
}

func CreateMaquinaria(c *gin.Context) {
	var maq models.Maquinaria
	if err := c.ShouldBindJSON(&maq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&maq)
	c.JSON(http.StatusCreated, maq)
}

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

func DeleteMaquinaria(c *gin.Context) {
	database.DB.Delete(&models.Maquinaria{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Eliminado"})
}

// ─── Neumáticos por Maquinaria ────────────────────────────────────────────────

func GetNeumaticosByMaq(c *gin.Context) {
	var maq models.Maquinaria
	if err := database.DB.Where("slug = ?", c.Param("id")).
		Preload("Neumaticos.Marca").First(&maq).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Maquinaria no encontrada"})
		return
	}
	c.JSON(http.StatusOK, maq.Neumaticos)
}

func CreateNeumatico(c *gin.Context) {
	var neu models.Neumatico
	if err := c.ShouldBindJSON(&neu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&neu)
	c.JSON(http.StatusCreated, neu)
}

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

func DeleteNeumatico(c *gin.Context) {
	database.DB.Delete(&models.Neumatico{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Eliminado"})
}

// ─── Marcas ───────────────────────────────────────────────────────────────────

func GetMarcas(c *gin.Context) {
	var marcas []models.Marca
	database.DB.Find(&marcas)
	c.JSON(http.StatusOK, marcas)
}

func CreateMarca(c *gin.Context) {
	var marca models.Marca
	if err := c.ShouldBindJSON(&marca); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&marca)
	c.JSON(http.StatusCreated, marca)
}

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

func DeleteMarca(c *gin.Context) {
	database.DB.Delete(&models.Marca{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Eliminado"})
}
