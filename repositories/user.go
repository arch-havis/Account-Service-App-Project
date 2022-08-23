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
	return 0, nil
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
