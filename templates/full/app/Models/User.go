package Models

import "{{ .ModulePath }}/database/orm"

type User struct {
	orm.Model
	Name     string
	Username string
	Email    string
	Password string
}

func (User) TableName() string { return "users" }
