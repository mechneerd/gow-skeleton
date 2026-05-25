package migrations

import (
	"github.com/mechneerd/gow/database/migration"
	"github.com/mechneerd/gow/database/schema"
)

type CreateUsersTable struct{}

func (CreateUsersTable) Up(m *schema.Builder) error {
	return m.Create("users", func(table *schema.Blueprint) {
		table.ID()
		table.String("name", 255)
		table.String("username", 255).Unique()
		table.String("email", 255).Unique()
		table.String("password", 255)
		table.Boolean("is_active").Default(true)
		table.Text("remember_token").Nullable()
		table.Timestamps()
	})
}

func (CreateUsersTable) Down(m *schema.Builder) error {
	return m.Drop("users")
}

func init() {
	migration.Register("2026_05_24_000001_create_users_table", CreateUsersTable{})
}