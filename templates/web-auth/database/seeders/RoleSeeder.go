package seeders

import (
	"fmt"
	"log"

	"{{ .ModulePath }}/app/Models"
	"golang.org/x/crypto/bcrypt"
)

func RoleSeeder() {
	fmt.Println("🌱 Running RoleSeeder...")

	// Create default roles
	roles := []string{"Super Admin", "Admin", "Editor", "User"}
	for _, roleName := range roles {
		role := Models.Role{Name: roleName, GuardName: "web"}
		if err := role.Save(); err != nil {
			log.Printf("Role %s already exists or failed: %v", roleName, err)
		}
	}

	// Create all basic permissions
	permissions := []string{
		"view-dashboard",
		"manage-users",
		"manage-roles",
		"manage-permissions",
		"edit-content",
		"publish-content",
	}
	for _, permName := range permissions {
		perm := Models.Permission{Name: permName, GuardName: "web"}
		if err := perm.Save(); err != nil {
			log.Printf("Permission %s already exists or failed: %v", permName, err)
		}
	}

	// Assign all permissions to Super Admin role
	superAdminRole := Models.Role{}.Where("name", "Super Admin").First()
	allPermissions := Models.Permission{}.Get()
	for _, perm := range allPermissions {
		superAdminRole.AttachPermission(perm)
	}

	// Create Super Admin user with hardcoded credentials
	password := "12345678"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	superAdmin := Models.User{
		Name:     "Super Administrator",
		Username: "superadmin",
		Email:    "superadmin@example.com",
		Password: string(hashedPassword),
	}
	if err := superAdmin.Save(); err != nil {
		log.Printf("Super Admin user may already exist: %v", err)
	} else {
		// Assign Super Admin role
		superAdminRole.AttachUser(&superAdmin)
		fmt.Println("✅ Super Admin created successfully!")
		fmt.Println("   Username: superadmin")
		fmt.Println("   Password: 12345678")
	}

	fmt.Println("✅ RoleSeeder completed.")
}
