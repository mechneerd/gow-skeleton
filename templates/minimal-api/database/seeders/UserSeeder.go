package seeders

import (
	"fmt"
	"log"

	"{{ .ModulePath }}/app/Models"
	"golang.org/x/crypto/bcrypt"
)

func UserSeeder() {
	fmt.Println("🌱 Running UserSeeder for minimal-api...")

	password := "12345678"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	admin := Models.User{
		Name:     "API Super Admin",
		Username: "superadmin",
		Email:    "superadmin@example.com",
		Password: string(hash),
	}
	if err := admin.Save(); err != nil {
		log.Println("User may already exist:", err)
	} else {
		fmt.Println("✅ API Super Admin created: superadmin / 12345678")
	}
}
