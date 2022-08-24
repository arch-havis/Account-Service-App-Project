package services

import (
	"account-service-app-project/entities"
	"account-service-app-project/repositories"
)

type User struct {
	Repositories repositories.User
}

func (s *User) Update(user entities.Users, id int) (int, error) {
	return 0, nil
}

func (s *User) Destroy(id int) (int, error) {
	return 0, nil
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
