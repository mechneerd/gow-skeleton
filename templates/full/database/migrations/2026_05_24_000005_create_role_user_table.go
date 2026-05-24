package migrations

import (
	"github.com/mechneerd/gow/database/migration"
	"github.com/mechneerd/gow/database/schema"
)

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
