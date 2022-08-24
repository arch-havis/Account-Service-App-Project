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
	userResult := entities.Users{}
	err := r.DB.QueryRow("SELECT user_id, no_telepon, password, nama, alamat, gender, saldo FROM users WHERE no_telepon = ?", nohp).Scan(&userResult.UserId, &userResult.NoTelepon, &userResult.Password, &userResult.Nama, &userResult.Alamat, &userResult.Gender, &userResult.Saldo)

	if err != nil {
		return entities.Users{}, err
	}

	return userResult, nil
}
