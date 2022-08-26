package controllers

import (
	"account-service-app-project/services"
	"fmt"
)

func Topup(topupService services.Topup) {
	var nominal float64
	var noHp string
	fmt.Println("\nMasukkkan nomor telepon tujuan: ")
	fmt.Scanln(&noHp)

	fmt.Println("Masukkkan nominal anda:")
	fmt.Scanln(&nominal)

	topupResult, err := topupService.Store(noHp, nominal)

	if err != nil {
		fmt.Println(err)
		return
	}

	if topupResult < 1 {
		fmt.Println("Tidak dapat topup, terdapat kesalahan sistem.")
		return
	}

	result := fmt.Sprintf("Topup berhasil ditambahkan dengan nominal: Rp.%v", nominal)
	fmt.Println(result)
}

func HistoryTopup(noHP string, topupService services.Topup) {
	topupHistoryResult, err := topupService.History(noHP)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(topupHistoryResult.UserTopup) < 1 {
		fmt.Println("Tidak ada data history topup")
	}

	for _, value := range topupHistoryResult.UserTopup {
		fmt.Println("\n=======")
		fmt.Println("Nama user     : ", topupHistoryResult.Nama)
		fmt.Println("Nominal Topup : ", value.NominalTopup)
		fmt.Println("Tanggal Topup : ", value.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}
