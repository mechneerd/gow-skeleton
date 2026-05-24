# Minimal API - {{ .AppName }}

Ultra-light API-only starter for GoW.

**Status**: ✅ Ready (with basic Auth support)

## Getting Started

```bash
go run main.go
```

## Database Setup

```bash
gow migrate
gow db:seed
```

After seeding, a default API user is available:
- Username: `superadmin`
- Password: `12345678`

### New in this starter

- Clean API landing page at `/`
- GoW Livewire pre-included (useful for internal admin UIs)
- Create views easily with the new generator:

```bash
gow make:view admin/users
```

## Environment

Copy `.env.example` to `.env` and configure as needed.

## Structure

- `main.go` - Application entrypoint (uses bootstrap)
- `bootstrap/app.go` - Bootstraps config + routes
- `config/app.go` + `config/database.go` - Configuration
- `routes/api.go` - API route definitions
- `.env.example` - Environment variables template

## Next Steps

- Add your API routes
- Set up database connection
- Add authentication (Sanctum recommended)
- Create models and controllers

Generated with GoW on {{ .Year }}
