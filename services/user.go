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
	return entities.Users{}, nil
}
