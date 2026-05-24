package Models

import (
	"{{ .ModulePath }}/auth/rbac"
	"{{ .ModulePath }}/database/orm"
)

type User struct {
	orm.Model
	rbac.HasRoles
	Name            string
	Username        string
	Email           string
	Password        string
	EmailVerifiedAt *string
	RememberToken   *string
}

func (User) TableName() string {
	return "users"
}

// SetDB wires a database connection for this user instance (RBAC queries).
// Recommended usage in generated projects:
//   db, _ := sql.Open(...)
//   rbac.SetDefaultDB(db)   // once at startup
// Or per user: user.SetDB(db)
func (u *User) SetDB(db *sql.DB) {
	u.HasRoles.db = db
}
