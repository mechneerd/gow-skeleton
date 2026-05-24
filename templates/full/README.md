# Full Starter Kit - {{ .AppName }}

**Status**: ✅ Full-featured with Auth + RBAC

This is the most complete starter kit for GoW. It includes everything from web-auth plus advanced structure.

## What's Included

- Full web + API routing
- Complete authentication (session + Sanctum)
- Full RBAC (Roles + Permissions + pivots)
- **Super Admin seeder**:
  - Username: `superadmin`
  - Password: `12345678`
- 6 database migrations (users, tokens, roles, permissions, pivots)
- Configuration system
- Ready for Queue, Mail, and more

## Getting Started

```bash
go run main.go

# Setup database
gow migrate
gow db:seed
```

After seeding you can login as:
- **superadmin / 12345678**

### New in this starter

- Rich landing page at `/` with GoBlade + Livewire demo
- GoW Livewire fully ready for interactive features

Quickly create new views:

```bash
gow make:view dashboard
gow make:view admin/settings
```

## Database Structure
Same as web-auth (full RBAC + sessions + verification/reset tokens).

See `database/migrations/` for design notes on guard_name and pivot naming.

## Recommended Use
Use this kit for medium to large production applications that need both web interface and API with proper authorization.

Generated with GoW on {{ .Year }}
