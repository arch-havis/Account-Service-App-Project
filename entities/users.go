package entities

type Users struct {
	UserId       int
	NoTelepon    string
	Nama         string
	Password     string
	Alamat       string
	Gender       string
	Saldo        float64
	UserTopup    []Topup
	UserTransfer []Transfer
}
