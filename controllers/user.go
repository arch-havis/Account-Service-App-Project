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
	PilihanGender := false
	fmt.Println("Masukan Nama")
	fmt.Scanln(&userData.Nama)
	fmt.Println("Masukan No Telepon")
	fmt.Scanln(&userData.NoTelepon)
	fmt.Println("Masukan Password")
	fmt.Scanln(&userData.Password)
	fmt.Println("Masukan Alamat")
	fmt.Scanln(&userData.Alamat)
	fmt.Println("Masukan Gender")
	fmt.Println("1 Laki-laki\n2 Perempuan")
	fmt.Scanln(&PilihanGender)
	fmt.Println("Masukan Saldo")
	fmt.Scanln(&userData.Saldo)

	switch PilihanGender {
	case true:
		userData.Gender = "Laki-laki"
	case false:
		userData.Gender = "Perempuan"
	}
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

func ReadUser(NoHp string, userService services.User) {
	UserData, err := userService.FindByNoHp(NoHp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("======")
	fmt.Println("Atas nama: ", UserData.Nama)
	fmt.Println("No Telpon: ", UserData.NoTelepon)
	fmt.Println("Alamat: ", UserData.Alamat)
	fmt.Println("Gender: ", UserData.Gender)
	fmt.Println("Saldo: ", UserData.Saldo)
	fmt.Println("======")
}

func UpdateUser(NoHp string, userService services.User) {
	userData := entities.Users{}
	PilihanGender := 1
	fmt.Println("Masukan Nama: ")
	fmt.Scanln(&userData.Nama)
	fmt.Println("Masukan No Telepon: ")
	fmt.Scanln(&userData.NoTelepon)
	fmt.Println("Masukan Password: ")
	fmt.Scanln(&userData.Password)
	fmt.Println("Masukan Alamat: ")
	fmt.Scanln(&userData.Alamat)
	fmt.Println("Masukan Gender")
	fmt.Println("1 Laki-laki\n2 Perempuan\n3 Skip Gender")
	fmt.Scanln(&PilihanGender)

	switch PilihanGender {
	case 1:
		userData.Gender = "Laki-laki"
	case 2:
		userData.Gender = "Perempuan"
	case 3:
		userData.Gender = ""
	}

	userResult, err := userService.Update(userData, NoHp)
	if err != nil {
		fmt.Println(err)
		return

	}
	if userResult < 1 {
		fmt.Println("Akun tidak berhasil di update, mohon coba lagi")
		return
	}
	fmt.Println("Akun berhasil di update")
}

func DeleteUser(NoHp string, userService services.User) (int, error) {
	userResult, err := userService.Destroy(NoHp)
	if err != nil {
		return -1, err

	}
	if userResult < 1 {
		return 0, nil
	}
	return 1, nil
}

func CheckSaldo(NoHp string, userService services.User) (float64, error) {
	userData, err := userService.FindByNoHp(NoHp)
	if err != nil {
		return 0, err
	}
	return userData.Saldo, nil
}
