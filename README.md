# Neumáticos Agroindustriales API

## Variables de Entorno

La API usa `os.Getenv()` directamente (sin biblioteca externa).

### Base de Datos

| Variable | Descripción | Ejemplo |
|----------|-------------|---------|
| `DATABASE_URL` | PostgreSQL de Railway (producción) | `postgresql://postgres:pass@host:5432/railway` |
| `DATABASE_PUBLIC_URL` | PostgreSQL público (desarrollo) | `postgresql://postgres:pass@proxy:19560/railway` |
| `DB_SSLMODE` | Modo SSL (`disable` en desarrollo) | `disable` |
| `PGHOST`, `PGUSER`, `PGPASSWORD`, `PGDATABASE`, `PGPORT` | Variables individuales (fallback) | - |

**Lógica de conexión** (`internal/database/database.go`):
1. Usa `DATABASE_URL` si existe (producción)
2. Falls back a `DATABASE_PUBLIC_URL` (desarrollo)
3. Si ninguna, usa variables individuales PG*

### Aplicación

| Variable | Descripción | Default |
|----------|-------------|---------|
| `APP_ENV` | Entorno | `development` |
| `PORT` | Puerto del servidor | `8080` |
| `FRONTEND_URL` | URL del frontend para CORS | `http://localhost:3000` |
| `NEXTJS_URL` | URL de Next.js para auth bridge | `http://localhost:3000` |
| `ADMIN_API_JWT_SECRET` | Secret para JWT admin | - |

### Desarrollo Local

```bash
# Copiar ejemplo
cp .env.example .env

# Editar con tus valores
nano .env
```

### Producción (Railway)

Railway inyecta automáticamente:
- `DATABASE_URL`
- `DATABASE_PUBLIC_URL`
- Variables PG*