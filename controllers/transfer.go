package controllers

import (
	"account-service-app-project/services"
	"fmt"
)

func Transfer(NoTelpPengirim string, serviceTransfer services.Transfer) {
	var NoTelpPenerima string
	var Nominal float64

	fmt.Println("Masukkan nomor telepon tujuan: ")
	fmt.Scanln(&NoTelpPenerima)
	fmt.Println("Masukkan nominal: ")
	fmt.Scanln(&Nominal)

	rowAffected, err := serviceTransfer.Store(NoTelpPengirim, NoTelpPenerima, Nominal)
	if err != nil {
		fmt.Println(err)
		return
	}
	if rowAffected < 1 {
		fmt.Println("Gagal melakukan transfer")
		return
	}
	fmt.Println("Berhasil melakukan transfer")
}
