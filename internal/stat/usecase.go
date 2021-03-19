package stat

import "github.com/yletamitlu/trainee/internal/models"

type StatUsecase interface {
	AddNewStatistic(data *models.Data) error
	GetStatisticByPeriod(since string, until string, param string) ([]*models.Data, error)
	RemoveStatistic() error
}
