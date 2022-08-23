package services

import (
	"account-service-app-project/entities"
	"account-service-app-project/repositories"
	"database/sql"
)

func Login(db *sql.DB, nomorHp string, password string) (entities.Users, error) {
	UserRepo := repositories.User{DB: db}
	userResult, err := UserRepo.FindByNoHp(nomorHp)

	if err != nil {
		return entities.Users{}, err
	}

	if userResult.Password != password {
		return entities.Users{}, nil
	}

	return userResult, nil
}
