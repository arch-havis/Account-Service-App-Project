package repositories

import "database/sql"

type Topup struct {
	DB *sql.DB
}

func (r *Topup) Store(userId int, nominal float64) (int, error) {

	/*
		-1 => dia gagal dieksekusi  dan die juga error
		0 => dia gagal dieksekusi tidak error, karena tidak ada data yang dieksekusi
		1 => berhasil dieksekusi.
	*/

	db, err := r.DB.Prepare("INSERT INTO topup (user_id, nominal_topup) VALUES (?,?)")
	if err != nil {
		return -1, err
	}

	result, err := db.Exec(userId, nominal)
	if err != nil {
		return -1, err
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return -1, err
	}

	return int(rowAffected), nil
}
