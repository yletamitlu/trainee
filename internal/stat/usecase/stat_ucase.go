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
		data.Cpm = (math.Round(cpm * 100) / 100 ) * 1000
	} else {
		data.Cpm = 0
	}

	err := su.statRepo.InsertInto(data)
	if err != nil {
		return err
	}
	return nil
}

func (su *StatUcase) GetStatisticByPeriod(since string, until string, param string) ([]*models.Data, error) {
	isValid := su.checkParam(param)
	if !isValid {
		param = ""
	}

	data, err := su.statRepo.GetStatistic(since, until, param)
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

func (su *StatUcase) checkParam(value string) bool {
	switch value {
	case "date":
		return true
	case "views":
		return true
	case "clicks":
		return true
	case "cost":
		return true
	case "cpc":
		return true
	case "cpm":
		return true
	default:
		return false
	}
}
