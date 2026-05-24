# {{ .AppName }} - API Starter

REST API project built with GoW.

Includes:
- Sanctum token authentication
- Form Request validation
- API Resources
- Eloquent-style ORM

## Getting Started

```bash
go run main.go
```

Recommended commands:
- `gow make:model Post --migration`
- `gow make:request StorePostRequest`
- `gow migrate`

### New in this starter

- Modern API-focused landing page at `/` (still beautiful, but API-oriented)
- GoW Livewire included (great for admin dashboards or internal tools)
- Use `gow make:view` for any web-based admin UI

```bash
gow make:view admin/dashboard
```

Generated with GoW on {{ .Year }}
