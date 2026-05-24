package Models

import (
	"{{ .ModulePath }}/database/orm"
)

type Role struct {
	orm.Model
	Name      string
	GuardName string
}

func (Role) TableName() string {
	return "roles"
}

// AttachPermission links a permission to this role
func (r *Role) AttachPermission(perm Permission) {
	// Placeholder for many-to-many logic
}

// AttachUser assigns this role to a user
func (r *Role) AttachUser(user *User) {
	// Placeholder for many-to-many logic
}
