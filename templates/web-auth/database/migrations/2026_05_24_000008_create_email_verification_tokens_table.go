package migrations

import (
	"github.com/mechneerd/gow/database/migration"
	"github.com/mechneerd/gow/database/schema"
)

type CreateEmailVerificationTokensTable struct{}

func (CreateEmailVerificationTokensTable) Up(m *schema.Builder) error {
	return m.Create("email_verification_tokens", func(table *schema.Blueprint) {
		table.String("email", 255)
		table.String("token", 255)
		table.Text("created_at").Nullable()
	})
}

func (CreateEmailVerificationTokensTable) Down(m *schema.Builder) error {
	return m.Drop("email_verification_tokens")
}

func init() {
	migration.Register("2026_05_24_000008_create_email_verification_tokens_table", CreateEmailVerificationTokensTable{})
}