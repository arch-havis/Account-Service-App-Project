package repositories

import (
	"account-service-app-project/entities"
	"database/sql"
)

type User struct {
	DB *sql.DB
}

func (r *User) Store(user entities.Users) (int, error) {
	return 0, nil
}

func (r *User) Update(user entities.Users, id int) (int, error) {
	userStmt, err := r.DB.Prepare("UPDATE users SET no_telepon = ?, nama = ?, password = ?, alamat = ?, gender = ?, saldo = ?")
	if err != nil {
		return -1, err
	}

	userResult, err := userStmt.Exec(user.NoTelepon, user.Nama, user.Password, user.Alamat, user.Gender, user.Saldo)
	if err != nil {
		return -1, err
	}

	rowAffected, err := userResult.RowsAffected()
	if err != nil {
		return -1, err
	}

	return int(rowAffected), nil
}

func (r *User) Destroy(id int) (int, error) {
	return 0, nil
}

func (r *User) FindById(id int) (entities.Users, error) {
	return entities.Users{}, nil
}

func (r *User) FindByNoHp(nohp string) (entities.Users, error) {
	userStmt, err := r.DB.Prepare(`
		SELECT u.*, tp.* FROM users u
		INNER JOIN topup tp ON tp.user_id = u.user_id
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
