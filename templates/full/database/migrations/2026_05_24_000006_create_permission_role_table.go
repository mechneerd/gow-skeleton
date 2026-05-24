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
		table.UnsignedBigInteger("permission_id")
		table.UnsignedBigInteger("role_id")
		table.Timestamps()
		table.Primary("permission_id", "role_id")
		table.Foreign("permission_id").References("id").On("permissions").OnDelete("cascade")
		table.Foreign("role_id").References("id").On("roles").OnDelete("cascade")
	})
}
