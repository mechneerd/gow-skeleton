package migrations

import (
	"github.com/mechneerd/gow/database/migration"
	"github.com/mechneerd/gow/database/schema"
)

func init() {
	migration.Register("2026_05_24_000007_create_sessions_table", CreateSessionsTable)
}

func CreateSessionsTable(m *schema.Builder) error {
	return m.Create("sessions", func(table *schema.Blueprint) {
		table.String("id", 255).Primary()
		table.Integer("user_id").Nullable()
		table.String("ip_address", 45).Nullable()
		table.Text("user_agent").Nullable()
		table.Text("payload")
		table.Integer("last_activity")
	})
}