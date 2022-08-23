package services

import (
	"account-service-app-project/entities"
	"account-service-app-project/repositories"
)

type Transfer struct {
	Repositories repositories.Transfer
}

func Store(userId int, userIdPenerima int, nominal float64) (int, error) {
	return 0, nil
}

func HistoryTransfer(userId int) ([]entities.Transfer, error) {
	return []entities.Transfer{}, nil
}
