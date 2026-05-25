package migrations

import (
	"github.com/mechneerd/gow/database/migration"
	"github.com/mechneerd/gow/database/schema"
)

type CreatePasswordResetTokensTable struct{}

func (CreatePasswordResetTokensTable) Up(m *schema.Builder) error {
	return m.Create("password_reset_tokens", func(table *schema.Blueprint) {
		table.String("email", 255)
		table.String("token", 255)
		table.Text("created_at").Nullable()
	})
}

func (CreatePasswordResetTokensTable) Down(m *schema.Builder) error {
	return m.Drop("password_reset_tokens")
}

func init() {
	migration.Register("2026_05_24_000009_create_password_reset_tokens_table", CreatePasswordResetTokensTable{})
}