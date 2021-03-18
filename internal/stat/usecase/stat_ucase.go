package usecase

import (
	"github.com/yletamitlu/trainee/internal/models"
	"github.com/yletamitlu/trainee/internal/stat"
	"math"
)

type StatUcase struct {
	statRepo stat.StatRepository
}

func NewStatUcase(repos stat.StatRepository) stat.StatUsecase {
	return &StatUcase{
		statRepo: repos,
	}
}

func (su *StatUcase) AddNewStatistic(data *models.Data) error {
	if data.Clicks != 0 {
		cpc := data.Cost / float64(data.Clicks)
		data.Cpc = math.Round(cpc * 100) / 100
	} else {
		data.Cpc = 0
	}

	if data.Views != 0 {
		cpm := data.Cost / float64(data.Views)
		data.Cpm = math.Round(cpm * 100) / 100
	} else {
		data.Cpm = 0
	}

	err := su.statRepo.InsertInto(data)
	if err != nil {
		return err
	}
	return nil
}

func (su *StatUcase) GetStatisticByPeriod(since string, until string) ([]*models.Data, error) {
	data, err := su.statRepo.GetStatistic(since, until)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (su *StatUcase) RemoveStatistic() error {
	err := su.statRepo.RemoveStatistic()
	if err != nil {
		return err
	}
	return nil
}
