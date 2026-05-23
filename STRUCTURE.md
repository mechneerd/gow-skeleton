# GoW Skeleton Structure

This document explains how the `gow-skeleton` repository is organized so contributors can easily add or modify starter kits.

## Repository Layout

```
gow-skeleton/
├── templates/
│   ├── minimal/          # Basic project with routing + views
│   ├── api/              # API-only starter (Sanctum, resources, JSON)
│   ├── web/              # Full web app with Blade
│   └── web-auth/         # Web + complete authentication (login, register, middleware, etc.)
├── scripts/
│   ├── post-install.sh
│   └── post-install.ps1
├── README.md
├── STRUCTURE.md          # This file
└── .gitignore
```

## Template Rules

Each folder under `templates/` is a complete, independent starter kit.

### Required Files in Every Template

- `go.mod.template`
- `main.go`
- `bootstrap/app.go`
- `config/` folder with at least `app.go` and `database.go`
- `.env.example`
- `routes/` folder (`web.go` or `api.go`)
- `README.md` inside the template (explains what the kit includes)

### Optional but Recommended

- `app/Models/`
- `database/migrations/`
- `resources/views/` (for web kits)
- `app/Http/Controllers/`

## Placeholders

Use these consistent placeholders inside template files. The `gow` CLI will replace them during `gow new`.

| Placeholder           | Example Value                  | Common Locations                     |
|-----------------------|--------------------------------|--------------------------------------|
| `{{ .AppName }}`      | myblog                         | .env, README, titles                 |
| `{{ .ModulePath }}`   | github.com/user/myblog         | go.mod.template                      |
| `{{ .DatabaseDriver }}` | sqlite / mysql / postgres    | config/database.go                   |
| `{{ .IncludeAuth }}`  | true / false                   | Conditional logic in templates       |
| `{{ .Year }}`         | 2026                           | License headers, README              |

## Adding a New Starter Kit

1. Create a new folder under `templates/`, e.g. `templates/full`.
2. Copy the structure from `web-auth` as a base.
3. Update all placeholders.
4. Add a short `README.md` inside the new template.
5. Update the main `README.md` with the new kit.
6. Test by manually copying the folder and replacing placeholders.

## Scripts Folder

The `scripts/` directory contains post-install scripts that the CLI can execute after copying files.

- `post-install.sh` — for Unix/macOS/Linux
- `post-install.ps1` — for Windows

These scripts usually run:
- `go mod tidy`
- `.env` file generation
- Optional database migration

## Versioning

- The skeleton repo uses semantic versioning tags (v1.0.0, v1.1.0, etc.).
- The GoW CLI should check the skeleton version and warn if incompatible.

## Best Practices

- Keep templates clean and well-commented.
- Never commit compiled binaries or vendor folders.
- Every file the user receives must be editable and belong to them (no hidden magic).
- Prefer composition over duplication when possible (share common base files via symlinks or copy if needed).

---

This structure follows the same philosophy as Laravel’s `laravel/laravel` skeleton + Breeze/Jetstream.
