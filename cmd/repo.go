package cmd

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Repo struct {
	Logging   *Logging
	DBOptions *pg.Options
}

func NewRepo(logging *Logging, addr string, user string, password string, database string) *Repo {
	return &Repo{
		Logging: logging,
		DBOptions: &pg.Options{
			Addr:     addr,
			User:     user,
			Password: password,
			Database: database,
		},
	}
}

func (repo *Repo) createSchema() error {
	db := pg.Connect(repo.DBOptions)
	defer db.Close()
	models := []interface{}{
		(*Stat)(nil),
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *Repo) saveStat(stat *Stat) error {
	db := pg.Connect(repo.DBOptions)
	defer db.Close()
	result, err := db.Model(stat).
		OnConflict("(date) DO UPDATE").
		Insert()
	if err != nil {
		return err
	}
	repo.Logging.DebugLog.Printf("%v", result)
	return nil
}

func (repo *Repo) findStatByDate(date StatDateTime) (*Stat, error) {
	db := pg.Connect(repo.DBOptions)
	defer db.Close()
	stat := &Stat{Date: date}
	err := db.Model(stat).Where("date = ?", date).Select()
	if err != nil {
		return nil, err
	}
	return stat, nil
}


func (repo *Repo) findAllStats() ([]Stat, error) {
	db := pg.Connect(repo.DBOptions)
	defer db.Close()
	var stats []Stat
	err := db.Model(&stats).Select()
	if err != nil {
		return nil, err
	}
	return stats, nil
}

