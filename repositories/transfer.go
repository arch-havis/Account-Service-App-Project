package repositories

import "database/sql"

type Transfer struct {
	DB *sql.DB
}

func (r *Transfer) Store(userIdPenerima int, userIdPengirim int, nominal float64) (int, error) {
	return 0, nil
}
