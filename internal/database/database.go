package database

import (
	"fmt"
	"log"
	"os"

	"github.com/WhosAnder/nei-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	// Railway inyecta DATABASE_URL (red privada).
	// En local, usa DATABASE_PUBLIC_URL o construye desde variables individuales.

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Println("📦 DATABASE_URL no definido, intentando DATABASE_PUBLIC_URL")
		dsn = os.Getenv("DATABASE_PUBLIC_URL")
	}
	if dsn == "" {
		// Fallback para desarrollo local con variables separadas
		log.Println("📦 Usando variables PG individuales (desarrollo local)")
		sslmode := os.Getenv("DB_SSLMODE")
		if sslmode == "" {
			sslmode = "disable"
		}
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/Mexico_City",
			os.Getenv("PGHOST"),
			os.Getenv("PGUSER"),
			os.Getenv("PGPASSWORD"),
			os.Getenv("PGDATABASE"),
			os.Getenv("PGPORT"),
			sslmode,
		)
	} else {
		log.Println("📦 Usando DATABASE_URL de Railway")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("❌ Error conectando a la base de datos: %v", err)
	}

	log.Println("✅ Conexión a PostgreSQL establecida.")
	DB = db
}

func Migrate() {
	err := DB.AutoMigrate(
		&models.Categoria{},
		&models.Maquinaria{},
		&models.Neumatico{},
		&models.Marca{},
		&models.Servicio{},
	)
	if err != nil {
		log.Fatalf("❌ Error en migración: %v", err)
	}
	log.Println("✅ Tablas migradas correctamente.")
}
