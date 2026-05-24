package migrations

import (
	"github.com/mechneerd/gow/database/migration"
	"github.com/mechneerd/gow/database/schema"
)

func init() {
	migration.Register("2026_05_24_000001_create_users_table", CreateUsersTable)
}

func CreateUsersTable(m *schema.Builder) error {
	return m.Create("users", func(table *schema.Blueprint) {
		table.ID()
		table.String("name", 255)
		table.String("username", 255).Unique()
		table.String("email", 255).Unique()
		table.String("password", 255)
		table.Timestamp("email_verified_at").Nullable()
		table.String("remember_token", 100).Nullable()
		table.Timestamps()
	})
}
