package migrations

import (
	"github.com/mechneerd/gow/database/migration"
	"github.com/mechneerd/gow/database/schema"
)

func init() {
	migration.Register("2026_05_24_000006_create_permission_role_table", CreatePermissionRoleTable)
}

func CreatePermissionRoleTable(m *schema.Builder) error {
	return m.Create("permission_role", func(table *schema.Blueprint) {
		table.Integer("permission_id")
		table.Integer("role_id")
		table.Unique()
	})
}