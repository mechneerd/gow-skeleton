package migrations

import (
	"github.com/mechneerd/gow/database/migration"
	"github.com/mechneerd/gow/database/schema"
)

type CreatePersonalAccessTokensTable struct{}

func (CreatePersonalAccessTokensTable) Up(m *schema.Builder) error {
	return m.Create("personal_access_tokens", func(table *schema.Blueprint) {
		table.ID()
		table.String("tokenable_type", 255)
		table.Integer("tokenable_id")
		table.String("name", 255)
		table.String("token", 64).Unique()
		table.Text("abilities").Nullable()
		table.Text("last_used_at").Nullable()
		table.Text("expires_at").Nullable()
		table.Timestamps()
	})
}

func (CreatePersonalAccessTokensTable) Down(m *schema.Builder) error {
	return m.Drop("personal_access_tokens")
}

func init() {
	migration.Register("2026_05_24_000002_create_personal_access_tokens_table", CreatePersonalAccessTokensTable{})
}