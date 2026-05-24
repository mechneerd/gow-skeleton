package Models

import "{{ .ModulePath }}/database/orm"

type Role struct {
	orm.Model
	Name      string
	GuardName string
}

func (Role) TableName() string { return "roles" }

func (r *Role) AttachPermission(p Permission) {}
func (r *Role) AttachUser(u *User)            {}
