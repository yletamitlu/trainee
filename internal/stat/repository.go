package stat

import (
	"github.com/yletamitlu/trainee/internal/models"
)

type StatRepository interface {
	InsertInto(data *models.Data) error
	GetStatistic(since string, until string, param string) ([]*models.Data, error)
	RemoveStatistic() error
}
