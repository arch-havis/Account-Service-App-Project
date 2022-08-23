package services

import (
	"account-service-app-project/entities"
	"account-service-app-project/repositories"
)

type Topup struct {
	Repositories repositories.Topup
}

func (s *Topup) Store(userId int, nominal float64) (int, error) {
	return 0, nil
}

func (s *Topup) TopupHistory() ([]entities.Topup, error) {
	return []entities.Topup{}, nil
}
