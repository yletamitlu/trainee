package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/yletamitlu/trainee/internal/models"
	"github.com/yletamitlu/trainee/internal/stat"
	"time"
)

type StatRepos struct {
	conn *sqlx.DB
}

func NewStatRepos(conn *sqlx.DB) stat.StatRepository {
	return &StatRepos{
		conn: conn,
	}
}

func (sr *StatRepos) InsertInto(data *models.Data) error {
	date, err := time.Parse("2006-01-02", data.Date)
	if err != nil{
		return err
	}

	if _, err := sr.conn.Exec(
		`INSERT INTO statistic(date, views, clicks, cost, cpc, cpm) VALUES ($1, $2, $3, $4, $5, $6)`,
		date,
		data.Views,
		data.Clicks,
		data.Cost,
		data.Cpc,
		data.Cpm);
		err != nil {
		return err
	}

	return nil
}

func (sr *StatRepos) GetStatistic(since string, until string) ([]*models.Data, error) {
	var dataArray []*models.Data

	sinceDate, err := time.Parse("2006-01-02", since)
	if err != nil{
		return nil, err
	}
	untilDate, err := time.Parse("2006-01-02", until)
	if err != nil{
		return nil, err
	}

	if err := sr.conn.Select(&dataArray,
		`SELECT * from statistic where date > $1 and date < $2 order by date desc`,
		sinceDate, untilDate); err != nil {
		return nil, err
	}

	return dataArray, nil
}

func (sr *StatRepos) RemoveStatistic() error {
	clearQuery := "TRUNCATE statistic RESTART IDENTITY"
	_, err := sr.conn.Exec(clearQuery)

	if err != nil {
		return err
	}

	return nil
}
