package entities

import "time"

type Transfer struct {
	TransferId      int
	UserId          int
	UserIdPenerima  int
	NominalTransfer float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
	UserPengirim    Users
	UserPenerima    Users
}
