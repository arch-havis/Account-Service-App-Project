package main

import (
	"account-service-app-project/config"
	"account-service-app-project/controllers"
	"account-service-app-project/repositories"
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
	var userRepo repositories.User = repositories.User{DB: db}
	var userService services.User = services.User{Repositories: userRepo}

	/*
		topup repositories dan topup service
	*/
	var topupRepository repositories.Topup = repositories.Topup{DB: db}
	var topupService services.Topup = services.Topup{
		RepositoriesUser:  userRepo,
		RepositoriesTopUp: topupRepository,
	}
	/*
		transfer repository dan transfer service
	*/
	var transferRepository repositories.Transfer = repositories.Transfer{DB: db}
	var transferService services.Transfer = services.Transfer{
		Repositories:     transferRepository,
		RepositoriesUser: userRepo,
	}

	for isLanjutkan {

		if NoHp == "" {
			fmt.Println("1. Register")
			fmt.Println("2. Login")
		} else {
			fmt.Println("3. Profil Saya\n4. Perbarui Profile Saya\n5. Hapus Profile Saya\n6. Topup\n7. Transfer\n8. History Topup\n9. History Transfer\n10.Lihat Profile Orang")
			fmt.Println("0. Logout dari akun")
		}

		// inputan user
		fmt.Scanln(&Pilihan)

		/*
			kasus 1: ketika user belum login
		*/
		if NoHp == "" {
			switch Pilihan {
			case 1:
				controllers.Register(userService)
			case 2:
				NoHp = controllers.Login(db)
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
				controllers.ReadUser(NoHp, userService)
			case 4:
				controllers.UpdateUser(NoHp, userService)
			case 5:
				userResult, err := controllers.DeleteUser(NoHp, userService)
				if err != nil {
					fmt.Println(err)
					break
				}
				if userResult < 1 {
					fmt.Println("Akun tidak bisa di hapus, coba lagi")
					break
				}
				fmt.Println("Akun sudah berhasil di hapus")
				NoHp = ""
			case 6:
				controllers.Topup(topupService)
			case 7:
				controllers.Transfer(NoHp, transferService)
				break
			case 8:
				controllers.HistoryTopup(NoHp, topupService)
			case 9:
				controllers.HistoryTransfer(NoHp, transferService)
			case 10:
				controllers.SearchProfile(userService)
			case 0:
				fmt.Println("Terimakasih Telah bertransaksi.")
				NoHp = ""
			default:
				fmt.Println("Masukkan sesuai di menu.")
			}
		}

		terminator := ""
		fmt.Print("\n\n\nApakah mau lanjut? y/n\n")
		fmt.Scanln(&terminator)
		if terminator != "y" {
			isLanjutkan = false
		}
	}
}
