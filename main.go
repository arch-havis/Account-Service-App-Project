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
			fmt.Println("=======[Selamat Datang Di Account Service App]===========")
			fmt.Println("")
			fmt.Println("1. Register")
			fmt.Println("2. Login")
		} else {
			fmt.Println("==========[Halaman Beranda]===========")
			fmt.Println("")
			fmt.Println("3. Profil Saya\n4. Perbarui Profile Saya\n5. Hapus Profile Saya\n6. Top-up\n7. Transfer\n8. History Top-up\n9. History Transfer\n10.Lihat Profile Orang")
			fmt.Println("0. Logout dari akun")
		}

		// inputan user
		fmt.Println("\nMasukkan pilihan anda: ")
		fmt.Scanln(&Pilihan)

		/*
			kasus 1: ketika user belum login
		*/
		if NoHp == "" {
			switch Pilihan {
			case 1:
				fmt.Println("\n======[Register Akun]=======")
				controllers.Register(userService)
			case 2:
				fmt.Println("\n======[Halaman Login]========")
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
				fmt.Println("\n=======[Profile Saya]=======")
				controllers.ReadUser(NoHp, userService)
			case 4:
				fmt.Println("\n=======[Perbarui Profile]=======")
				controllers.UpdateUser(NoHp, userService)
			case 5:
				fmt.Println("\n=======[Hapus Profile]=======")
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
				fmt.Println("\n=======[Top-up]=======")
				controllers.Topup(topupService)
			case 7:
				fmt.Println("\n=======[Transfer]=======")
				controllers.Transfer(NoHp, transferService)
			case 8:
				fmt.Println("\n=======[History Topup]=======")
				controllers.HistoryTopup(NoHp, topupService)
			case 9:
				fmt.Println("\n=======[History Transfer]=======")
				controllers.HistoryTransfer(NoHp, transferService)
			case 10:
				fmt.Println("\n=======[Cari Profile Orang]=======")
				controllers.SearchProfile(userService)
			case 0:
				fmt.Println("\n=======[Terimakasih telah bertransaksi.]=======")
				NoHp = ""
			default:
				fmt.Println("Masukkan sesuai di menu.")
			}
		}

		terminator := ""
		fmt.Print("\n\nApakah mau lanjut? y/n\n")
		fmt.Scanln(&terminator)
		if terminator != "y" {
			isLanjutkan = false
		}
	}
}
