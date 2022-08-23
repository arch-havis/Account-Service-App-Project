package repositories

import "database/sql"

type Topup struct {
	DB *sql.DB
}

func (r *Topup) Store(userId int, nominal float64) (int, error) {
	return 0, nil
}
