# NEI API — Neumáticos Agroindustriales

API REST en Go para la gestión del catálogo de neumáticos agrícolas e industriales. Construida con **Gin**, **GORM** y **PostgreSQL**, documentada con **Swagger UI**.

---

## Stack

| Capa        | Tecnología                       |
|-------------|----------------------------------|
| Lenguaje    | Go 1.25                          |
| Framework   | Gin v1.12                        |
| ORM         | GORM v1.31 + driver `pgx`        |
| Base de datos | PostgreSQL (Railway)            |
| Autenticación | Better Auth (bridge Next.js) + JWT admin |
| Documentación | Swagger / swaggo               |
| Despliegue  | Railway                          |

---

## Estructura del proyecto

```
nei-api/
├── cmd/
│   ├── api/
│   │   └── main.go          # Punto de entrada del servidor
│   └── seed/
│       └── main.go          # Seed inicial de datos
├── docs/
│   ├── docs.go              # Swagger generado (swag init)
│   ├── swagger.json
│   └── swagger.yaml
├── internal/
│   ├── database/
│   │   └── database.go      # Conexión y AutoMigrate
│   ├── handlers/
│   │   ├── catalog.go       # Handlers CRUD (Categorías, Maquinaria, Neumáticos, Marcas, Servicios)
│   │   └── dto.go           # Structs de respuesta para Swagger
│   ├── middleware/
│   │   ├── auth.go          # AuthRequired — valida sesión Better Auth
│   │   └── admin_jwt.go     # AdminJWTRequired — valida JWT firmado desde Next.js
│   ├── models/
│   │   └── models.go        # Modelos GORM (Categoria, Maquinaria, Neumatico, Marca, Servicio)
│   └── router/
│       └── router.go        # Registro de rutas públicas y admin
├── .env.example
├── go.mod
└── go.sum
```

---

## Modelos de datos

```
Categoria  1──* Maquinaria  1──* Neumatico *──1 Marca
```

| Modelo      | Tabla         | Campos clave                                              |
|-------------|---------------|-----------------------------------------------------------|
| `Categoria` | `categorias`  | `slug` (unique), `nombre`, `descripcion`, `imagen_url`     |
| `Maquinaria`| `maquinarias` | `slug` (unique), `nombre`, `icono_nombre`, `categoria_id`  |
| `Neumatico` | `neumaticos`  | `nombre`, `medida`, `patron`, `precio`, `maquinaria_id`, `marca_id` |
| `Marca`     | `marcas`      | `slug` (unique), `nombre`, `logo_url`                     |
| `Servicio`  | `servicios`   | `titulo`, `descripcion`, `icono_nombre`, `texto_boton`    |

Todos los modelos incluyen `gorm.Model` (campos `id`, `created_at`, `updated_at`, `deleted_at` — soft delete).

---

## Endpoints

Base URL: `/api/v1`  
Swagger UI: `http://localhost:8080/swagger/index.html`

### Públicos (lectura)

| Método | Ruta                              | Descripción                            |
|--------|-----------------------------------|----------------------------------------|
| GET    | `/categorias`                     | Lista todas las categorías             |
| GET    | `/categorias/:slug/maquinaria`    | Maquinaria por categoría (por slug)    |
| GET    | `/maquinaria/:slug/neumaticos`    | Neumáticos por maquinaria (por slug)   |
| GET    | `/marcas`                         | Lista todas las marcas                 |
| GET    | `/servicios`                      | Lista todos los servicios              |
| GET    | `/health`                         | Health check `{"status":"ok"}`         |

### Admin (requieren autenticación)

> En **desarrollo** (`APP_ENV != production/staging`), el middleware omite la validación.  
> En **producción/staging**, se requiere un JWT válido emitido por Next.js (`AdminJWTRequired`).

| Método | Ruta                          | Descripción                    |
|--------|-------------------------------|--------------------------------|
| POST   | `/admin/categorias`           | Crear categoría                |
| PUT    | `/admin/categorias/:id`       | Actualizar categoría           |
| DELETE | `/admin/categorias/:id`       | Eliminar categoría             |
| POST   | `/admin/maquinaria`           | Crear maquinaria               |
| PUT    | `/admin/maquinaria/:id`       | Actualizar maquinaria          |
| DELETE | `/admin/maquinaria/:id`       | Eliminar maquinaria            |
| POST   | `/admin/neumaticos`           | Crear neumático                |
| PUT    | `/admin/neumaticos/:id`       | Actualizar neumático           |
| DELETE | `/admin/neumaticos/:id`       | Eliminar neumático             |
| POST   | `/admin/marcas`               | Crear marca                    |
| PUT    | `/admin/marcas/:id`           | Actualizar marca               |
| DELETE | `/admin/marcas/:id`           | Eliminar marca                 |

---

## Autenticación

El proyecto usa **dos capas de seguridad** para rutas admin:

### 1. Better Auth Bridge (`middleware/auth.go`)
Valida que exista una sesión activa consultando a Next.js en `/api/auth/get-session`.  
- El cookie `better-auth.session_token` se reenvía al servidor Next.js.
- Si la sesión es inválida o no existe, responde `401`.

### 2. JWT Admin (`middleware/admin_jwt.go`)
Para rutas que requieren rol explícito de admin:
- Espera header `Authorization: Bearer <token>`.
- Token firmado con `ADMIN_API_JWT_SECRET` (HS256).
- Claims requeridos: `role: "admin"`, `iss: "next-admin"`, `aud: "nei-api-admin"`.

---

## Variables de entorno

Copia `.env.example` a `.env` para desarrollo local:

```bash
cp .env.example .env
```

### Base de datos (PostgreSQL / Railway)

| Variable              | Descripción                          | Ejemplo                                           |
|-----------------------|--------------------------------------|---------------------------------------------------|
| `DATABASE_URL`        | URL privada de Railway (producción)  | `postgresql://postgres:pass@host:5432/railway`    |
| `DATABASE_PUBLIC_URL` | URL pública (desarrollo local)       | `postgresql://postgres:pass@proxy:PORT/railway`   |
| `DB_SSLMODE`          | Modo SSL (solo fallback individual)  | `disable`                                         |
| `PGHOST`              | Host (fallback)                      | —                                                 |
| `PGUSER`              | Usuario (fallback)                   | `postgres`                                        |
| `PGPASSWORD`          | Contraseña (fallback)                | —                                                 |
| `PGDATABASE`          | Nombre de la base (fallback)         | `railway`                                         |
| `PGPORT`              | Puerto (fallback)                    | `5432`                                            |

**Lógica de conexión** (`internal/database/database.go`):
1. Usa `DATABASE_URL` si existe → producción/Railway red privada.
2. Si no, usa `DATABASE_PUBLIC_URL` → desarrollo local vía proxy público.
3. Si ninguna, construye DSN desde variables `PG*` individuales.

### Aplicación

| Variable              | Descripción                              | Default                |
|-----------------------|------------------------------------------|------------------------|
| `APP_ENV`             | Entorno (`development`, `production`)    | `development`          |
| `PORT`                | Puerto del servidor                      | `8080`                 |
| `FRONTEND_URL`        | URL del frontend para CORS               | `http://localhost:3000`|
| `NEXTJS_URL`          | URL de Next.js para el auth bridge       | `http://localhost:3000`|
| `ADMIN_API_JWT_SECRET`| Secret para firmar/validar JWTs admin    | —                      |

---

## Desarrollo local

### Requisitos

- Go 1.25+
- PostgreSQL accesible (local o Railway public proxy)

### Iniciar el servidor

```bash
# 1. Configurar variables
cp .env.example .env
# Editar .env con tus credenciales de PostgreSQL

# 2. Ejecutar
go run ./cmd/api/main.go
```

El servidor estará disponible en `http://localhost:8080`.

### Seed de datos iniciales

Pobla la base de datos con categorías, maquinaria, marcas, neumáticos y servicios de ejemplo:

```bash
go run ./cmd/seed/main.go
```

El seed usa `FirstOrCreate` para ser idempotente (seguro de ejecutar varias veces).

**Datos incluidos:**
- **2 categorías**: Agrícola, Industrial
- **13 tipos de maquinaria**: Tractor, Implemento, Trilladora, Minicargador (agrícola), Grúa, Montacargas, Cargador, Retroexcavadora, Vibro compactador, Motoconformadora, Camión, Camión muevetierra, Minicargador (industrial)
- **13 marcas**: Pirelli, SEBA, Goodyear, Eurogrip, Samson, Galaxy, Numa, y otras
- **5 neumáticos** de ejemplo con medidas y patrones reales
- **3 servicios**: Cotización, Asesoría personalizada, Montajes

### Regenerar documentación Swagger

```bash
# Instalar swag si no está instalado
go install github.com/swaggo/swag/cmd/swag@latest

# Generar desde la raíz del proyecto
swag init -g cmd/api/main.go -o docs
```

---

## Despliegue en Railway

Railway inyecta automáticamente las siguientes variables:
- `DATABASE_URL` (red privada)
- `DATABASE_PUBLIC_URL` (proxy público)
- `RAILWAY_SERVICE_NAME` ← detectado por `main.go` para confirmar entorno Railway
- Variables `PG*` individuales

Configura manualmente en Railway:
- `APP_ENV=production`
- `FRONTEND_URL=https://tu-dominio.com`
- `NEXTJS_URL=https://tu-app-nextjs.com`
- `ADMIN_API_JWT_SECRET=<secreto-seguro>`

---

## CORS

Configurado en `cmd/api/main.go`. Permite peticiones desde `FRONTEND_URL`:

```
Access-Control-Allow-Origin: <FRONTEND_URL>
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
Access-Control-Allow-Headers: Content-Type, Authorization, Cookie
Access-Control-Allow-Credentials: true
```

---

## Swagger UI

Accesible en: `http://localhost:8080/swagger/index.html`

La documentación incluye:
- Todos los endpoints públicos y admin
- Esquemas de request/response via DTOs (`handlers/dto.go`)
- Seguridad `BetterAuthSession` (cookie `better-auth.session_token`)