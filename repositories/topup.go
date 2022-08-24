package repositories

import (
	"account-service-app-project/entities"
	"database/sql"
)

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

func (r *Topup) History(nohp string) (entities.Users, error) {
	userStmt, err := r.DB.Prepare(`
		SELECT u.*, tp.* FROM users u
		RIGHT JOIN topup tp ON tp.user_id = u.user_id
		WHERE u.no_telepon = ?
	`)

	if err != nil {
		return entities.Users{}, err
	}

	rows, err := userStmt.Query(nohp)
	if err != nil {
		return entities.Users{}, err
	}

	userData := entities.Users{}

	for rows.Next() {
		topupData := entities.Topup{}
		err := rows.Scan(
			&userData.UserId,
			&userData.NoTelepon,
			&userData.Nama, &userData.Password,
			&userData.Alamat,
			&userData.Gender,
			&userData.Saldo,
			&topupData.TopupId,
			&topupData.UserId,
			&topupData.NominalTopup,
			&topupData.CreatedAt,
			&topupData.UpdatedAt,
		)

		if err != nil {
			return entities.Users{}, err
		}

		userData.UserTopup = append(userData.UserTopup, topupData)
	}

	return userData, nil
}
