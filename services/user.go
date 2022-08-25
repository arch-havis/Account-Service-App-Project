package services

import (
	"account-service-app-project/entities"
	"account-service-app-project/repositories"
)

type User struct {
	Repositories repositories.User
}

func (s *User) Store(user entities.Users) (int, error) {
	resultStore, err := s.Repositories.Store(user)
	if err != nil {
		return -1, err
	}
	return resultStore, nil
}

func (s *User) Update(userInput entities.Users, NoHp string) (int, error) {
	userDataFromDb, err := s.Repositories.FindByNoHp(NoHp)
	if err != nil {
		return -1, err
	}

	if userInput.Nama != "" && userInput.Nama != userDataFromDb.Nama {
		userDataFromDb.Nama = userInput.Nama
	}

	if userInput.NoTelepon != "" && userInput.NoTelepon != userDataFromDb.NoTelepon {
		userDataFromDb.NoTelepon = userInput.NoTelepon
	}

	if userInput.Alamat != "" && userInput.Alamat != userDataFromDb.Alamat {
		userDataFromDb.Alamat = userInput.Alamat
	}

	if userInput.Gender != "" && userInput.Gender != userDataFromDb.Gender {
		userDataFromDb.Gender = userInput.Gender
	}

	if userInput.Password != "" && userInput.Password != userDataFromDb.Password {
		userDataFromDb.Password = userInput.Password
	}
	userResult, err := s.Repositories.Update(userDataFromDb, userDataFromDb.UserId)
	if err != nil {
		return -1, err
	}
	return userResult, nil
}

func (s *User) Destroy(NoHp string) (int, error) {
	userResult, err := s.Repositories.FindByNoHp(NoHp)
	if err != nil {
		return -1, err
	}
	rowAffected, err := s.Repositories.Destroy(userResult.UserId)
	if err != nil {
		return -1, err
	}
	return rowAffected, nil
}

func (s *User) FindById(id int) (entities.Users, error) {
	return entities.Users{}, nil
}

func (s *User) FindByNoHp(noHp string) (entities.Users, error) {
	userResult, err := s.Repositories.FindByNoHp(noHp)
	if err != nil {
		return entities.Users{}, err
	}
	return userResult, nil
}
