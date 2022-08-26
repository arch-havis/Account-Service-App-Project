package controllers

import (
	"account-service-app-project/services"
	"fmt"
)

func Transfer(NoTelpPengirim string, serviceTransfer services.Transfer) {
	var NoTelpPenerima string
	var Nominal float64

	fmt.Println("\nMasukkan nomor telepon tujuan: ")
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

func HistoryTransfer(NoHp string, serviceTransfer services.Transfer) {
	userResults, err := serviceTransfer.HistoryTransfer(NoHp)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, value := range userResults {
		fmt.Println("\n===========")
		fmt.Println("Nama Pengirim    : ", value.UserPengirim.Nama)
		fmt.Println("Nama Penerima    : ", value.UserPenerima.Nama)
		fmt.Println("Nominal Transfer : ", value.NominalTransfer)
		fmt.Println("Waktu transfer   : ", value.CreatedAt.Format("2006-01-02 15:04:05"))
	}

}
