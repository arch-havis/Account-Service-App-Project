package main

import (
	"account-service-app-project/config"
	"account-service-app-project/services"
	"fmt"
)

func main() {
	db := config.DB()
	defer db.Close()

	/***
		global variabel, digunakan untuk menampung beberapa value,
		yang butuh variable global.
	***/
	var isLanjutkan bool = true
	var NoHp string
	var Pilihan int

	/*
		user repositories dan user service
	*/
	// var userRepo repositories.User = repositories.User{DB: db}
	// var userService services.User = services.User{Repositories: userRepo}

	for isLanjutkan {
		if NoHp == "" {
			fmt.Println("1. Register")
			fmt.Println("2. Login")
		} else {
			fmt.Println("3. Profil Saya\n4. Perbarui Profile Saya\n5. Hapus Profile Saya\n6. Topup\n7. Transfer\n8. History Topup\n9. History Transfer\n10.Lihat Profile Orang")
		}
		fmt.Println("0. Logout dari akun")
		fmt.Scanln(&Pilihan)

		if NoHp == "" {
			switch Pilihan {
			case 1:
				break
			case 2:
				noHp := ""
				password := ""
				fmt.Println("Masukkan no hp")
				fmt.Scanln(&noHp)
				fmt.Println("Masukkan password")
				fmt.Scanln(&password)

				userResult, err := services.Login(db, noHp, password)
				fmt.Println(userResult)
				if err != nil {
					fmt.Println(err)
					break
				}

				if userResult.UserId == 0 {
					fmt.Println("No Hp atau password tidak ditemukan.")
					break
				}
				NoHp = userResult.NoTelepon
				fmt.Println("Selamat datang berhasil login.")
				break
			case 0:
				fmt.Println("Terimakasih Telah bertransaksi.")
				NoHp = ""
				break
			default:
				fmt.Println("Masukkan sesuai di menu.")
				break
			}
		} else {
			switch Pilihan {
			case 3:
				break
			case 4:
				break
			case 5:
				break
			case 6:
				break
			case 7:
				break
			case 8:
				break
			case 9:
				break
			case 10:
				break
			case 0:
				fmt.Println("Terimakasih Telah bertransaksi.")
				NoHp = ""
				break
			default:
				fmt.Println("Masukkan sesuai di menu.")
				break
			}
		}
		terminator := ""
		fmt.Println("Apakah mau lanjut? y/n")
		fmt.Scanln(&terminator)
		if terminator != "y" {
			isLanjutkan = false
		}
	}
}
