package entities

import "time"

type Users struct {
	UserId       int
	NoTelepon    string
	Nama         string
	Password     string
	Alamat       string
	Gender       string
	Saldo        float64
	UpdatedAt    time.Time
	UserTopup    []Topup
	UserTransfer []Transfer
}
