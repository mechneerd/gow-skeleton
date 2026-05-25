package migrations

import (
	"github.com/mechneerd/gow/database/migration"
	"github.com/mechneerd/gow/database/schema"
)

type CreateSessionsTable struct{}

func (CreateSessionsTable) Up(m *schema.Builder) error {
	return m.Create("sessions", func(table *schema.Blueprint) {
		table.String("id", 255)
		table.Integer("user_id").Nullable()
		table.String("ip_address", 45).Nullable()
		table.Text("user_agent").Nullable()
		table.Text("payload")
		table.Integer("last_activity")
	})
}

func (CreateSessionsTable) Down(m *schema.Builder) error {
	return m.Drop("sessions")
}

func init() {
	migration.Register("2026_05_24_000007_create_sessions_table", CreateSessionsTable{})
}