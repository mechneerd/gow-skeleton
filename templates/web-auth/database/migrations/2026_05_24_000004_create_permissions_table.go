package migrations

import (
	"github.com/mechneerd/gow/database/migration"
	"github.com/mechneerd/gow/database/schema"
)

// Design note: see create_roles_table.go for guard_name rationale.
func init() {
	migration.Register("2026_05_24_000004_create_permissions_table", CreatePermissionsTable)
}

func CreatePermissionsTable(m *schema.Builder) error {
	return m.Create("permissions", func(table *schema.Blueprint) {
		table.ID()
		table.String("name", 255)
		table.String("guard_name", 255).Default("web")
		table.Timestamps()

		table.Unique("name", "guard_name")
	})
}
