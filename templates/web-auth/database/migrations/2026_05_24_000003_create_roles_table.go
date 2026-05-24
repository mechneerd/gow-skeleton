package migrations

import (
	"github.com/mechneerd/gow/database/migration"
	"github.com/mechneerd/gow/database/schema"
)

// Design note (2026-05-24):
// guard_name is kept for Spatie-style parity and future multi-guard support.
// In current GoW usage it will almost always be "web".
// If you only ever need one guard, you can treat this as a constant in application code
// instead of a column. We keep the column for now to stay compatible with common RBAC patterns.

func init() {
	migration.Register("2026_05_24_000003_create_roles_table", CreateRolesTable)
}

func CreateRolesTable(m *schema.Builder) error {
	return m.Create("roles", func(table *schema.Blueprint) {
		table.ID()
		table.String("name", 255)
		table.String("guard_name", 255).Default("web")
		table.Timestamps()

		table.Unique("name", "guard_name")
	})
}
