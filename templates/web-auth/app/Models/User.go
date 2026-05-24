package Models

import (
	"{{ .ModulePath }}/database/orm"
)

type User struct {
	orm.Model
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

// HasRoles trait methods will be added here later
func (u *User) HasRole(roleName string) bool {
	// Placeholder - will be implemented with RBAC
	return false
}
