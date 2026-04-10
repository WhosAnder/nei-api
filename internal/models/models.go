package models

import "gorm.io/gorm"

// ─── Categoria ────────────────────────────────────────────────────────────────

type Categoria struct {
	gorm.Model
	Slug        string      `gorm:"uniqueIndex;not null" json:"slug"`
	Nombre      string      `gorm:"not null"             json:"nombre"`
	Descripcion string      `json:"descripcion"`
	ImagenURL   string      `json:"imagen_url"`
	Maquinaria  []Maquinaria `gorm:"foreignKey:CategoriaID" json:"maquinaria,omitempty"`
}

// ─── Maquinaria ───────────────────────────────────────────────────────────────

type Maquinaria struct {
	gorm.Model
	Slug        string      `gorm:"uniqueIndex;not null" json:"slug"`
	Nombre      string      `gorm:"not null"             json:"nombre"`
	IconoNombre string      `json:"icono_nombre"`
	ImagenURL   string      `json:"imagen_url"`
	Descripcion string      `json:"descripcion"`
	CategoriaID uint        `gorm:"not null"             json:"categoria_id"`
	Neumaticos  []Neumatico `gorm:"foreignKey:MaquinariaID" json:"neumaticos,omitempty"`
}

// ─── Neumatico ────────────────────────────────────────────────────────────────

type Neumatico struct {
	gorm.Model
	Nombre      string `gorm:"not null" json:"nombre"`
	Medida      string `json:"medida"`
	Patron      string `json:"patron"`
	Precio      string `json:"precio"`
	ImagenURL   string `json:"imagen_url"`
	MarcaID     *uint  `json:"marca_id"`
	MaquinariaID uint  `gorm:"not null" json:"maquinaria_id"`
	Marca       Marca  `gorm:"foreignKey:MarcaID" json:"marca,omitempty"`
}

// ─── Marca ────────────────────────────────────────────────────────────────────

type Marca struct {
	gorm.Model
	Slug    string `gorm:"uniqueIndex;not null" json:"slug"`
	Nombre  string `gorm:"not null"             json:"nombre"`
	LogoURL string `json:"logo_url"`
}

// ─── Servicio ─────────────────────────────────────────────────────────────────

type Servicio struct {
	gorm.Model
	Titulo      string `gorm:"not null" json:"titulo"`
	Descripcion string `json:"descripcion"`
	IconoNombre string `json:"icono_nombre"`
	TextoBoton  string `json:"texto_boton"`
}
