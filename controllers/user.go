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
	return userResult.NoTelepon
}

func SearchProfile(userService services.User) {
	var noHp string
	fmt.Println("Masukkan No Hp")
	fmt.Scanln(&noHp)
	userResult, err := userService.FindByNoHp(noHp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("=========")
	fmt.Println("Atas nama: ", userResult.Nama)
	fmt.Println("No Telepon: ", userResult.NoTelepon)
	fmt.Println("Gender: ", userResult.Gender)
	fmt.Println("=========")
}
