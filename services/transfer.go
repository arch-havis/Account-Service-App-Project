package services

import (
	"account-service-app-project/entities"
	"account-service-app-project/repositories"
	"fmt"
	"time"
)

type Transfer struct {
	Repositories     repositories.Transfer
	RepositoriesUser repositories.User
}

func (s *Transfer) Store(NoTelpPengirim string, NoHpPenerima string, nominal float64) (int, error) {
	userDataPengirim, err := s.RepositoriesUser.FindByNoHp(NoTelpPengirim)
	if err != nil {
		return -1, err
	}

	userDataPenerima, err := s.RepositoriesUser.FindByNoHp(NoHpPenerima)
	if err != nil {
		return -1, err
	}

	if userDataPengirim.Saldo < nominal {
		return -1, fmt.Errorf("Maaf saldo tidak cukup.")
	}

	userDataPengirim.Saldo -= nominal
	userDataPenerima.Saldo += nominal

	userDataPengirim.UpdatedAt = time.Now()
	userDataPenerima.UpdatedAt = time.Now()

	pengirimAffectedRow, err := s.RepositoriesUser.Update(userDataPengirim, userDataPengirim.UserId)
	if err != nil {
		return -1, err
	}

	if pengirimAffectedRow < 1 {
		return pengirimAffectedRow, nil
	}

	penerimaAffectedRow, err := s.RepositoriesUser.Update(userDataPenerima, userDataPenerima.UserId)
	if err != nil {
		return -1, err
	}

	if pengirimAffectedRow < 1 {
		return penerimaAffectedRow, nil
	}

	affectedRow, err := s.Repositories.Store(userDataPengirim.UserId, userDataPenerima.UserId, nominal)
	if err != nil {
		return -1, err
	}

	return affectedRow, nil
}

func (s *Transfer) HistoryTransfer(NoHp string) ([]entities.Transfer, error) {
	userResult, err := s.RepositoriesUser.FindByNoHp(NoHp)
	if err != nil {
		return []entities.Transfer{}, err
	}

	transferResult, err := s.Repositories.GetHistory(userResult.UserId)
	if err != nil {
		return []entities.Transfer{}, err
	}

	return transferResult, nil
}
