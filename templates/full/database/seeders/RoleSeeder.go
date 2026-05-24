package seeders

import (
	"fmt"
	"log"

	"{{ .ModulePath }}/app/Models"
	"golang.org/x/crypto/bcrypt"
)

func RoleSeeder() {
	fmt.Println("🌱 Running RoleSeeder (Full Kit)...")

	roles := []string{"Super Admin", "Admin", "Editor", "User"}
	for _, roleName := range roles {
		role := Models.Role{Name: roleName, GuardName: "web"}
		role.Save()
	}

	permissions := []string{"view-dashboard", "manage-users", "manage-roles", "edit-content"}
	for _, permName := range permissions {
		perm := Models.Permission{Name: permName, GuardName: "web"}
		perm.Save()
	}

	// Super Admin setup
	superAdminRole := Models.Role{}.Where("name", "Super Admin").First()
	allPerms := Models.Permission{}.Get()
	for _, p := range allPerms {
		superAdminRole.AttachPermission(p)
	}

	password := "12345678"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	adminUser := Models.User{
		Name:     "Super Administrator",
		Username: "superadmin",
		Email:    "superadmin@example.com",
		Password: string(hash),
	}
	if err := adminUser.Save(); err == nil {
		superAdminRole.AttachUser(&adminUser)
		fmt.Println("✅ Super Admin created: superadmin / 12345678")
	}

	fmt.Println("✅ Full kit RoleSeeder completed.")
}
