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
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/Mexico_City",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

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
