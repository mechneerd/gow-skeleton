package migrations

import (
	"github.com/mechneerd/gow/database/migration"
	"github.com/mechneerd/gow/database/schema"
)

func init() {
	migration.Register("2026_05_24_000003_create_roles_table", CreateRolesTable)
}

func CreateRolesTable(m *schema.Builder) error {
	return m.Create("roles", func(table *schema.Blueprint) {
		table.ID()
		table.String("name", 255)
		table.String("guard_name", 255).Default("web")
		table.Timestamps()
	})
}