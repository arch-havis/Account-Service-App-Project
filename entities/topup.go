package entities

import "time"

type Topup struct {
	TopupId      int
	UserId       int
	NominalTopup float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserData     Users
}
