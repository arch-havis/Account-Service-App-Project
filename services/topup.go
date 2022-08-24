package services

import (
	"account-service-app-project/entities"
	"account-service-app-project/repositories"
)

type Topup struct {
	RepositoriesTopUp repositories.Topup
	RepositoriesUser  repositories.User
}

func (s *Topup) Store(noHp string, nominal float64) (int, error) {
	userResult, err := s.RepositoriesUser.FindByNoHp(noHp)
	if err != nil {
		return -1, err
	}

	topupResult, err := s.RepositoriesTopUp.Store(userResult.UserId, nominal)
	if err != nil {
		return -1, err
	}

	if topupResult > 0 {
		userResult.Saldo += nominal
		updateUserResult, err := s.RepositoriesUser.Update(userResult, userResult.UserId)
		if err != nil {
			return -1, err
		}

		if updateUserResult > 0 {
			return 1, nil
		}
	}

	return 0, nil
}

func (s *Topup) TopupHistory() ([]entities.Topup, error) {
	return []entities.Topup{}, nil
}
