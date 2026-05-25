package migrations

import (
	"github.com/mechneerd/gow/database/migration"
	"github.com/mechneerd/gow/database/schema"
)

type CreateRoleUserTable struct{}

func (CreateRoleUserTable) Up(m *schema.Builder) error {
	return m.Create("role_user", func(table *schema.Blueprint) {
		table.Integer("role_id")
		table.Integer("user_id")
	})
}

func (CreateRoleUserTable) Down(m *schema.Builder) error {
	return m.Drop("role_user")
}

func init() {
	migration.Register("2026_05_24_000005_create_role_user_table", CreateRoleUserTable{})
}