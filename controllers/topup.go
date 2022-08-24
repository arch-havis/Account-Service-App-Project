package controllers

import (
	"account-service-app-project/services"
	"fmt"
)

func Topup(noHp string, topupService services.Topup) {
	fmt.Println("Masukkkan nominal anda:")

	var nominal float64
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
