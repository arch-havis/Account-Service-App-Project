package main

import (
	"account-service-app-project/config"
	"account-service-app-project/controllers"
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

		/*
			kasus 1: ketika user belum login
		*/
		if NoHp == "" {
			switch Pilihan {
			case 1:
				break
			case 2:
				NoHp = controllers.Login(db)
				break
			case 0:
				fmt.Println("Terimakasih Telah bertransaksi.")
				NoHp = ""
			default:
				fmt.Println("Masukkan sesuai di menu.")
			}
		} else {
			/*
				kasus 2: ketika user sudah login
			*/
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
			default:
				fmt.Println("Masukkan sesuai di menu.")
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
