package main

import (
	"log"
	"os"

	"github.com/WhosAnder/nei-api/internal/database"
	"github.com/WhosAnder/nei-api/internal/models"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar .env en local; en Railway Railway inyecta las vars
	appEnv := os.Getenv("APP_ENV")
	if appEnv != "production" && appEnv != "staging" {
		godotenv.Load()
	}

	database.Connect()
	database.Migrate()

	db := database.DB

	log.Println("🌱 Iniciando seed de datos...")

	// ─── Marcas ───────────────────────────────────────────────────────────────
	marcas := []models.Marca{
		{Slug: "pirelli", Nombre: "Pirelli", LogoURL: "/images/marcas/pirelli.png"},
		{Slug: "seba", Nombre: "SEBA", LogoURL: "/images/marcas/seba.png"},
		{Slug: "goodyear", Nombre: "Goodyear", LogoURL: "/images/marcas/goodyear.png"},
		{Slug: "eurogrip", Nombre: "Eurogrip", LogoURL: "/images/marcas/eurogrip.png"},
		{Slug: "samson", Nombre: "Samson", LogoURL: "/images/marcas/samson.png"},
		{Slug: "galaxy", Nombre: "Galaxy", LogoURL: "/images/marcas/galaxy.png"},
		{Slug: "numa", Nombre: "Numa", LogoURL: "/images/marcas/numa.png"},
		// Marcas de los neumáticos de ejemplo (para relaciones)
		{Slug: "seba-ultra", Nombre: "Seba Ultra", LogoURL: ""},
		{Slug: "firestone", Nombre: "Firestone", LogoURL: ""},
		{Slug: "michelin", Nombre: "Michelin", LogoURL: ""},
		{Slug: "alliance", Nombre: "Alliance", LogoURL: ""},
		{Slug: "camso", Nombre: "Camso", LogoURL: ""},
		{Slug: "titan", Nombre: "Titan", LogoURL: ""},
	}

	for i := range marcas {
		db.FirstOrCreate(&marcas[i], models.Marca{Slug: marcas[i].Slug})
	}
	log.Printf("✅ %d marcas insertadas", len(marcas))

	// ─── Categorías ───────────────────────────────────────────────────────────
	catAgricola := models.Categoria{
		Slug:        "agricola",
		Nombre:      "Agrícola",
		Descripcion: "Neumáticos especializados para maquinaria agrícola",
		ImagenURL:   "/images/agricultural-machinery-new.png",
	}
	catIndustrial := models.Categoria{
		Slug:        "industrial",
		Nombre:      "Industrial",
		Descripcion: "Neumáticos para maquinaria industrial y construcción",
		ImagenURL:   "/images/industrial-machinery.png",
	}
	db.FirstOrCreate(&catAgricola, models.Categoria{Slug: "agricola"})
	db.FirstOrCreate(&catIndustrial, models.Categoria{Slug: "industrial"})
	log.Println("✅ 2 categorías insertadas")

	// ─── Maquinaria Agrícola ──────────────────────────────────────────────────
	maqAgricola := []models.Maquinaria{
		{Slug: "tractor", Nombre: "Tractor agrícola", IconoNombre: "Construction", ImagenURL: "/maquinas/Minicargador.png", Descripcion: "Neumáticos para tractores agrícolas", CategoriaID: catAgricola.ID},
		{Slug: "implemento", Nombre: "Implemento (Empacadora)", IconoNombre: "Truck", ImagenURL: "/maquinas/cargador.png", Descripcion: "Neumáticos para empacadoras", CategoriaID: catAgricola.ID},
		{Slug: "trilladora", Nombre: "Trilladora", IconoNombre: "Construction", ImagenURL: "/images/trilladora.png", Descripcion: "Neumáticos para trilladoras", CategoriaID: catAgricola.ID},
		{Slug: "minicargador-agricola", Nombre: "Minicargador", IconoNombre: "Truck", ImagenURL: "/maquinas/Minicargador.png", Descripcion: "Neumáticos para minicargadores agrícolas", CategoriaID: catAgricola.ID},
	}

	// ─── Maquinaria Industrial ────────────────────────────────────────────────
	maqIndustrial := []models.Maquinaria{
		{Slug: "grua", Nombre: "Grúa", IconoNombre: "Construction", ImagenURL: "/maquinas/grua.png", Descripcion: "Neumáticos para grúas", CategoriaID: catIndustrial.ID},
		{Slug: "montacargas", Nombre: "Montacargas", IconoNombre: "Truck", ImagenURL: "/maquinas/montacargas.png", Descripcion: "Neumáticos para montacargas", CategoriaID: catIndustrial.ID},
		{Slug: "cargador", Nombre: "Cargador", IconoNombre: "Construction", ImagenURL: "/maquinas/cargador.png", Descripcion: "Neumáticos para cargadores", CategoriaID: catIndustrial.ID},
		{Slug: "retroexcavadora", Nombre: "Retroexcavadora", IconoNombre: "Truck", ImagenURL: "/images/retroexcavadora.png", Descripcion: "Neumáticos para retroexcavadoras", CategoriaID: catIndustrial.ID},
		{Slug: "vibrocompactador", Nombre: "Vibro compactador", IconoNombre: "Construction", ImagenURL: "/maquinas/VibroCompactador.png", Descripcion: "Neumáticos para vibro compactadores", CategoriaID: catIndustrial.ID},
		{Slug: "motoconformadora", Nombre: "Motoconformadora", IconoNombre: "Truck", ImagenURL: "/maquinas/Motoconformadora.png", Descripcion: "Neumáticos para motoconformadoras", CategoriaID: catIndustrial.ID},
		{Slug: "camion", Nombre: "Camión", IconoNombre: "Construction", ImagenURL: "/maquinas/camion.png", Descripcion: "Neumáticos para camiones", CategoriaID: catIndustrial.ID},
		{Slug: "camion-muevetierra", Nombre: "Camión muevetierra", IconoNombre: "Truck", ImagenURL: "/maquinas/CamionMuevetierra.png", Descripcion: "Neumáticos para camiones muevetierra", CategoriaID: catIndustrial.ID},
		{Slug: "minicargador-industrial", Nombre: "Minicargador", IconoNombre: "Construction", ImagenURL: "/maquinas/Minicargador.png", Descripcion: "Neumáticos para minicargadores industriales", CategoriaID: catIndustrial.ID},
	}

	allMaq := append(maqAgricola, maqIndustrial...)
	maqMap := map[string]models.Maquinaria{}
	for i := range allMaq {
		db.FirstOrCreate(&allMaq[i], models.Maquinaria{Slug: allMaq[i].Slug})
		maqMap[allMaq[i].Slug] = allMaq[i]
	}
	log.Printf("✅ %d maquinarias insertadas", len(allMaq))

	// ─── Neumáticos ───────────────────────────────────────────────────────────
	// Buscar marcas por slug para asociar FK
	marcaMap := map[string]uint{}
	for _, m := range marcas {
		var found models.Marca
		db.Where("slug = ?", m.Slug).First(&found)
		marcaMap[m.Slug] = found.ID
	}

	neumaticos := []models.Neumatico{
		// Tractor
		{Nombre: "Neumático agrícola 15.5-38 RTW", Medida: "15.5-38", Patron: "RTW", Precio: "$1,250", ImagenURL: "/images/hero-background.png", MarcaID: marcaMap["seba-ultra"], MaquinariaID: maqMap["tractor"].ID},
		{Nombre: "Neumático agrícola 18.4-34", Medida: "18.4-34", Patron: "R-1", Precio: "$1,450", ImagenURL: "/images/hero-background.png", MarcaID: marcaMap["firestone"], MaquinariaID: maqMap["tractor"].ID},
		{Nombre: "Neumático agrícola 16.9-30", Medida: "16.9-30", Patron: "AGRIBIB", Precio: "$1,350", ImagenURL: "/images/hero-background.png", MarcaID: marcaMap["michelin"], MaquinariaID: maqMap["tractor"].ID},
		// Retroexcavadora
		{Nombre: "Neumático industrial 12.5/80-18", Medida: "12.5/80-18", Patron: "SKS 532", Precio: "$650", ImagenURL: "/images/hero-background.png", MarcaID: marcaMap["camso"], MaquinariaID: maqMap["retroexcavadora"].ID},
		{Nombre: "Neumático industrial 17.5L-24", Medida: "17.5L-24", Patron: "LD 250", Precio: "$950", ImagenURL: "/images/hero-background.png", MarcaID: marcaMap["titan"], MaquinariaID: maqMap["retroexcavadora"].ID},
	}

	for i := range neumaticos {
		db.Create(&neumaticos[i])
	}
	log.Printf("✅ %d neumáticos insertados", len(neumaticos))

	// ─── Servicios ────────────────────────────────────────────────────────────
	servicios := []models.Servicio{
		{Titulo: "Cotización", Descripcion: "Solicita precio para neumáticos agrícolas e industriales según medida, maquinaria o aplicación.", IconoNombre: "FileText", TextoBoton: "Solicitar cotización"},
		{Titulo: "Asesoría personalizada", Descripcion: "Te ayudamos a elegir el neumático ideal para tu maquinaria.", IconoNombre: "MessageCircle", TextoBoton: ""},
		{Titulo: "Montajes", Descripcion: "Contamos con personal capacitado para montaje de llantas en sitio o taller.", IconoNombre: "Settings", TextoBoton: ""},
	}

	for i := range servicios {
		db.FirstOrCreate(&servicios[i], models.Servicio{Titulo: servicios[i].Titulo})
	}
	log.Printf("✅ %d servicios insertados", len(servicios))

	log.Println("🎉 Seed completado exitosamente.")
}
