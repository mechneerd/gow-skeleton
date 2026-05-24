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
