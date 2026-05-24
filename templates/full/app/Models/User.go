package Models

import (
	"{{ .ModulePath }}/auth/rbac"
	"{{ .ModulePath }}/database/orm"
)

type User struct {
	orm.Model
	rbac.HasRoles
	Name     string
	Username string
	Email    string
	Password string
}

func (User) TableName() string { return "users" }
