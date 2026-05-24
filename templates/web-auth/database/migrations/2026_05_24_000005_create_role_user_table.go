package migrations

import (
	"github.com/mechneerd/gow/database/migration"
	"github.com/mechneerd/gow/database/schema"
)

// Design note on pivot naming (2026-05-24):
// We use role_user and permission_role (simple, user-focused).
// Laravel Spatie uses model_has_roles / role_has_permissions for polymorphic support
// (roles assignable to Teams, Organizations, etc.).
// If GoW ever needs to assign roles to non-User models, we can evolve to polymorphic tables later.
// Current design keeps things simple and sufficient for 95% of applications.

func init() {
	migration.Register("2026_05_24_000005_create_role_user_table", CreateRoleUserTable)
}

func CreateRoleUserTable(m *schema.Builder) error {
	return m.Create("role_user", func(table *schema.Blueprint) {
		table.UnsignedBigInteger("role_id")
		table.UnsignedBigInteger("user_id")
		table.Timestamps()

		table.Primary("role_id", "user_id")
		table.Foreign("role_id").References("id").On("roles").OnDelete("cascade")
		table.Foreign("user_id").References("id").On("users").OnDelete("cascade")
	})
}
