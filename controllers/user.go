package controllers

import (
	"account-service-app-project/services"
	"database/sql"
	"fmt"
)

func Login(db *sql.DB) string {
	noHp := ""
	password := ""
	fmt.Println("Masukkan no hp")
	fmt.Scanln(&noHp)
	fmt.Println("Masukkan password")
	fmt.Scanln(&password)

	userResult, err := services.Login(db, noHp, password)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	if userResult.UserId == 0 {
		fmt.Println("No Hp atau password tidak ditemukan.")
		return ""
	}

	fmt.Println("Selamat datang berhasil login.")
	return userResult.Nama
}
