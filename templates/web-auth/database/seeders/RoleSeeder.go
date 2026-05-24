package seeders

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"{{ .ModulePath }}/app/Models"
	"golang.org/x/crypto/bcrypt"
)

// RoleSeeder seeds default roles, permissions, and the Super Admin user.
// It accepts an optional *sql.DB. If nil, it will try to open one from environment variables.
func RoleSeeder(db *sql.DB) {
	fmt.Println("🌱 Running RoleSeeder...")

	var err error
	if db == nil {
		db, err = openDefaultDB()
		if err != nil {
			log.Printf("Could not open database for seeding: %v", err)
			return
		}
		defer db.Close()
	}

	// Create default roles
	roles := []string{"Super Admin", "Admin", "Editor", "User"}
	for _, roleName := range roles {
		_, err := db.Exec(`
			INSERT OR IGNORE INTO roles (name, guard_name, created_at, updated_at) 
			VALUES (?, 'web', datetime('now'), datetime('now'))
		`, roleName)
		if err != nil {
			log.Printf("Role %s: %v", roleName, err)
		}
	}

	// Create permissions
	permissions := []string{
		"view-dashboard", "manage-users", "manage-roles",
		"manage-permissions", "edit-content", "publish-content",
	}
	for _, permName := range permissions {
		_, err := db.Exec(`
			INSERT OR IGNORE INTO permissions (name, guard_name, created_at, updated_at)
			VALUES (?, 'web', datetime('now'), datetime('now'))
		`, permName)
		if err != nil {
			log.Printf("Permission %s: %v", permName, err)
		}
	}

	// Give all permissions to Super Admin
	var superAdminRoleID int
	db.QueryRow("SELECT id FROM roles WHERE name = 'Super Admin'").Scan(&superAdminRoleID)

	rows, _ := db.Query("SELECT id FROM permissions")
	defer rows.Close()
	for rows.Next() {
		var permID int
		rows.Scan(&permID)
		db.Exec(`
			INSERT OR IGNORE INTO permission_role (permission_id, role_id, created_at, updated_at)
			VALUES (?, ?, datetime('now'), datetime('now'))
		`, permID, superAdminRoleID)
	}

	// Create Super Admin user
	password := "12345678"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	_, err = db.Exec(`
		INSERT OR IGNORE INTO users (name, username, email, password, created_at, updated_at)
		VALUES (?, ?, ?, ?, datetime('now'), datetime('now'))
	`, "Super Administrator", "superadmin", "superadmin@example.com", string(hashedPassword))

	if err != nil {
		log.Printf("Super Admin user error: %v", err)
	} else {
		var userID int
		db.QueryRow("SELECT id FROM users WHERE username = 'superadmin'").Scan(&userID)
		db.Exec(`
			INSERT OR IGNORE INTO role_user (role_id, user_id, created_at, updated_at)
			VALUES (?, ?, datetime('now'), datetime('now'))
		`, superAdminRoleID, userID)

		fmt.Println("✅ Super Admin created successfully!")
		fmt.Println("   Username: superadmin")
		fmt.Println("   Password: 12345678")
	}

	fmt.Println("✅ RoleSeeder completed.")
}

// openDefaultDB tries to open a database connection using environment variables.
// This is used when no DB is passed to RoleSeeder.
func openDefaultDB() (*sql.DB, error) {
	driver := "sqlite"
	dsn := "database.sqlite"

	if d := os.Getenv("DB_CONNECTION"); d != "" {
		driver = d
	}

	switch driver {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	case "postgres":
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
			os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
	}

	return sql.Open(driver, dsn)
}
