# {{ .AppName }} - Web + Auth + RBAC Starter

**Status**: ✅ Production-ready Auth + RBAC included

This is the recommended starter kit when you need full authentication with Role-Based Access Control.

## What's Included

- Complete user authentication (login, register, logout, remember me)
- Sanctum API token authentication
- Full RBAC system (Roles + Permissions)
- Pre-built migrations for users, tokens, roles, permissions, and pivot tables
- **Super Admin seeder** with credentials:
  - Username: `superadmin`
  - Password: `12345678`
- Eloquent-style ORM with User, Role, Permission models
- Form Request validation + CSRF protection
- Blade-like templating

## Getting Started

```bash
# 1. Run the application
go run main.go
# or
gow serve

# 2. Run migrations (creates all auth + RBAC tables)
gow migrate

# 3. Seed the database (creates Super Admin + default roles/permissions)
gow db:seed
```

## Default Super Admin Account

After running `gow db:seed`:

- **Username**: `superadmin`
- **Password**: `12345678`
- **Role**: Super Admin (has every permission)

**Important**: Change this password immediately in production!

## Database Structure

- `users`
- `personal_access_tokens` (Sanctum - polymorphic)
- `roles` + `permissions` (with `guard_name` column for future parity)
- `role_user` + `permission_role` (pivots — user-focused for simplicity)
- `sessions`
- `email_verification_tokens`
- `password_reset_tokens`

### Schema Decisions
- Pivot tables use simple `role_user` / `permission_role` names (not polymorphic like Spatie).
- `guard_name` kept for Spatie compatibility and future multi-guard needs (currently always "web").
- `remember_token` included for session-based "remember me" support.
- Additional tables (`sessions`, verification/reset tokens) included for complete basic auth flows.

See migration files for detailed design notes.

## RBAC Usage Examples

```go
user := auth.User()
if user.HasRole("Super Admin") { ... }
if user.Can("manage-users") { ... }
```

In Blade templates:
```blade
@role('Super Admin')
    <p>Only super admins see this</p>
@endrole
```

## Next Steps

1. Customize routes in `routes/web.go`
2. Create controllers with `gow make:controller`
3. Generate more policies: `gow make:policy PostPolicy`
4. Add your own permissions in `database/seeders/RoleSeeder.go`

Generated with GoW on {{ .Year }}. Ready for production use.
