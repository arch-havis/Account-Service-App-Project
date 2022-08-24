package controllers

import (
	"account-service-app-project/entities"
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

func Register(serviceUser services.User) {
	var userData entities.Users
	fmt.Println("Masukan Nama")
	fmt.Scanln(&userData.Nama)
	fmt.Println("Masukan No Telepon")
	fmt.Scanln(&userData.NoTelepon)
	fmt.Println("Masukan Password")
	fmt.Scanln(&userData.Password)
	fmt.Println("Masukan Alamat")
	fmt.Scanln(&userData.Alamat)
	fmt.Println("Masukan Gender")
	fmt.Scanln(&userData.Gender)
	fmt.Println("Masukan Saldo")
	fmt.Scanln(&userData.Saldo)

	fmt.Println(userData)

	userResult, err := serviceUser.Store(userData)

	if err != nil {
		fmt.Println(err)
		return
	}

	if userResult < 1 {
		fmt.Println("Gagal Register Akun")
		return
	}

	fmt.Println("Berhasil Membuat Akun")

}
