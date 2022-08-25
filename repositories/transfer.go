package repositories

import "database/sql"

type Transfer struct {
	DB *sql.DB
}

func (r *Transfer) Store(userIdPenerima int, userIdPengirim int, nominal float64) (int, error) {
	transferStmt, err := r.DB.Prepare("INSERT INTO transfer (user_id, user_id_penerima, nominal_transfer) VALUES (?, ?, ?)")
	if err != nil {
		return -1, err
	}

	transferResult, err := transferStmt.Exec(userIdPenerima, userIdPengirim, nominal)
	if err != nil {
		return -1, err
	}

	transferAffected, err := transferResult.RowsAffected()
	if err != nil {
		return -1, err
	}

	return int(transferAffected), err
}
