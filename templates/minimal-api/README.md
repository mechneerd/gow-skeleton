# Minimal API - {{ .AppName }}

Ultra-light API-only starter for GoW.

**Status**: ✅ Ready

## Getting Started

```bash
go run main.go
```

## Environment

Copy `.env.example` to `.env` and configure as needed.

```bash
go run main.go
```

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
