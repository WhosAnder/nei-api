package handlers

// ─── DTOs para documentación Swagger ─────────────────────────────────────────
// Estas estructuras planas representan las respuestas JSON de la API.
// No contienen gorm.Model para que swaggo pueda parsearlas correctamente.

// CategoriaResponse es la respuesta JSON de una categoría
type CategoriaResponse struct {
	ID          uint   `json:"id"`
	Slug        string `json:"slug"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	ImagenURL   string `json:"imagen_url"`
}

// MaquinariaResponse es la respuesta JSON de una maquinaria
type MaquinariaResponse struct {
	ID          uint   `json:"id"`
	Slug        string `json:"slug"`
	Nombre      string `json:"nombre"`
	IconoNombre string `json:"icono_nombre"`
	ImagenURL   string `json:"imagen_url"`
	Descripcion string `json:"descripcion"`
	CategoriaID uint   `json:"categoria_id"`
}

// NeumaticoResponse es la respuesta JSON de un neumático
type NeumaticoResponse struct {
	ID           uint   `json:"id"`
	Nombre       string `json:"nombre"`
	Medida       string `json:"medida"`
	Patron       string `json:"patron"`
	Precio       string `json:"precio"`
	ImagenURL    string `json:"imagen_url"`
	MarcaID      uint   `json:"marca_id"`
	MaquinariaID uint   `json:"maquinaria_id"`
}

// MarcaResponse es la respuesta JSON de una marca
type MarcaResponse struct {
	ID      uint   `json:"id"`
	Slug    string `json:"slug"`
	Nombre  string `json:"nombre"`
	LogoURL string `json:"logo_url"`
}

// ErrorResponse es la respuesta de error estándar
type ErrorResponse struct {
	Error string `json:"error"`
}

// MessageResponse es la respuesta de operaciones simples
type MessageResponse struct {
	Message string `json:"message"`
}

// ServicioResponse es la respuesta JSON de un servicio
type ServicioResponse struct {
	ID          uint   `json:"id"`
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	IconoNombre string `json:"icono_nombre"`
	TextoBoton  string `json:"texto_boton"`
}
