package repositories

import (
	"account-service-app-project/entities"
	"database/sql"
)

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

func (r *Transfer) GetHistory(userId int) ([]entities.Transfer, error) {
	transferHistory := []entities.Transfer{}

	namaUserPengirim := ""
	err := r.DB.QueryRow("SELECT nama FROM users WHERE user_id = ?", userId).Scan(&namaUserPengirim)

	transferStmt, err := r.DB.Prepare(`
		SELECT tf.nominal_transfer, tf.created_at, tf.user_id_penerima, u.nama FROM transfer tf
		JOIN users u ON u.user_id = tf.user_id_penerima
		WHERE tf.user_id = ?
	`)
	if err != nil {
		return transferHistory, err
	}

	rows, err := transferStmt.Query(userId)
	if err != nil {
		return transferHistory, err
	}

	for rows.Next() {
		transfer := entities.Transfer{}
		err := rows.Scan(
			&transfer.NominalTransfer,
			&transfer.CreatedAt,
			&transfer.UserIdPenerima,
			&transfer.UserPenerima.Nama,
		)
		transfer.UserPengirim.Nama = namaUserPengirim
		if err != nil {
			return transferHistory, err
		}
		transferHistory = append(transferHistory, transfer)
	}

	return transferHistory, nil
}
