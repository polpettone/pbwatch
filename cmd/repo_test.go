package cmd

import (
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/go-test/deep"
	"testing"
)

const addr = ":16432"
const user = "pbwatch"
const password = "pbwatch"
const db = "pbwatch"


func Test_save_and_load_stats(t *testing.T) {
	logging := NewLogging()
	repo := NewRepo(logging, addr, user, password, db)
	postgres := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().
		Username(repo.DBOptions.User).
		Password(repo.DBOptions.Password).
		Database(repo.DBOptions.Database).
		Port(16432))

	err := postgres.Start()
	if err != nil {
		t.Errorf("Error %v start embedded db", err)
	}

	err = repo.createSchema()
	if err != nil {
		t.Errorf("Create schema error %v", err)
	}

	stat0:= &Stat{
		Date:   DateTime{SimpleDate(2020, 10, 29)},
		RunningDistance: 4.5,
		PressUpCount: 4,
		Weight: 75.5,
		BodyFeeling: 3,
		EmotionalFeeling: 4,
	}

	err = repo.saveStat(stat0)
	if err != nil {
		t.Errorf("save stat %v", err)
	}

	expectedStatCount := 1
	foundStats, err := repo.findAllStats()
	if err != nil {
		t.Errorf("Find all stats %v", err)
	}

	if expectedStatCount != len(foundStats) {
		t.Errorf("wanted %d stats but got %d", expectedStatCount, len(foundStats))
	}

	if diff := deep.Equal(foundStats[0], *stat0); diff != nil {
		t.Error(diff)
	}

	stat1:= &Stat{
		Date:   DateTime{SimpleDate(2020, 10, 29)},
		RunningDistance: 4.5,
		PressUpCount: 4,
		Weight: 75.5,
		BodyFeeling: 0,
		EmotionalFeeling: 1,
	}

	err = repo.saveStat(stat1)
	if err != nil {
		t.Errorf("save stat %v", err)
	}

	expectedStatCount = 1
	foundStats, err = repo.findAllStats()
	if err != nil {
		t.Errorf("Find all stats%v", err)
	}

	if expectedStatCount != len(foundStats) {
		t.Errorf("wanted %d stats but got %d", expectedStatCount, len(foundStats))
	}

	if diff := deep.Equal(foundStats[0], *stat1); diff != nil {
		t.Error(diff)
	}

	foundStat, err := repo.findStatByDate(DateTime{SimpleDate(2020, 10, 29)})

	if err != nil {
		t.Errorf("%v", err)
	}

	if diff := deep.Equal(foundStat, stat1); diff != nil {
		t.Error(diff)
	}

	shouldNotFound, err := repo.findStatByDate(DateTime{SimpleDate(2010, 1, 1)})

	if err == nil {
		t.Errorf("%s", "error should not be nil")
	}

	if shouldNotFound != nil {
		t.Errorf("%v", shouldNotFound)
	}


	stat2:= &Stat{
		Date:   DateTime{SimpleDate(2020, 12, 1)},
		RunningDistance: 22,
		PressUpCount: 100,
		Weight: 80,
		BodyFeeling: 5,
		EmotionalFeeling: 5,
	}


	err = repo.saveStat(stat2)
	if err != nil {
		t.Errorf("save stat %v", err)
	}

	foundStat, err = repo.findStatByDate(DateTime{SimpleDate(2020, 12, 1)})

	if err != nil {
		t.Errorf("%v", err)
	}

	if diff := deep.Equal(foundStat, stat2); diff != nil {
		t.Error(diff)
	}

	err = postgres.Stop()
	if err != nil {
		t.Errorf("Error %v stop embedded db", err)
	}

}


