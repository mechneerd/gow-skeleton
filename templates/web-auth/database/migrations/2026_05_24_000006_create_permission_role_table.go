package migrations

import (
	"github.com/mechneerd/gow/database/migration"
	"github.com/mechneerd/gow/database/schema"
)

type CreatePermissionRoleTable struct{}

func (CreatePermissionRoleTable) Up(m *schema.Builder) error {
	return m.Create("permission_role", func(table *schema.Blueprint) {
		table.Integer("permission_id")
		table.Integer("role_id")
	})
}

func (CreatePermissionRoleTable) Down(m *schema.Builder) error {
	return m.Drop("permission_role")
}

func init() {
	migration.Register("2026_05_24_000006_create_permission_role_table", CreatePermissionRoleTable{})
}