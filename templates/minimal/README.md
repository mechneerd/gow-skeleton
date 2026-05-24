# {{ .AppName }} - Minimal Starter

**Status**: ✅ Ready

Ultra-light web starter for GoW.

## Getting Started

```bash
go run main.go
```

## Structure

- `main.go` - entrypoint
- `bootstrap/app.go`
- `config/app.go`
- `routes/` (basic web routes)
- `.env.example`

## Next Steps

Use `gow make:*` to generate models, controllers, etc.

### New in this starter

- **Beautiful landing page** at `/` built with GoBlade (`.goblade`)
- **GoW Livewire** included out of the box — build reactive UI without page reloads
- Example: `app/Livewire/Counter.go` + `resources/views/welcome.goblade`

### Create new views

```bash
gow make:view dashboard          # Creates resources/views/dashboard.goblade
gow make:view admin/users        # Creates resources/views/admin/users.goblade
```

Generated with GoW on {{ .Year }}
