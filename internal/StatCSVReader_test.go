package internal

import (
	"github.com/go-test/deep"
	"testing"
)

func Test_should_read_csv(t *testing.T) {
	fileToRead := "testStats.csv"
	stats, err := readStatCSV(fileToRead)

	expectedStat0 := &Stat{
		Date:             StatDateTime{SimpleDate(2020, 10, 29)},
		RunningDistance:  4.5,
		PressUpCount:     4,
		Weight:           75.5,
		BodyFeeling:      3,
		EmotionalFeeling: 4,
	}

	expectedStat1 := &Stat{
		Date:             StatDateTime{SimpleDate(2020, 10, 30)},
		RunningDistance:  5.5,
		PressUpCount:     8,
		Weight:           76.5,
		BodyFeeling:      4,
		EmotionalFeeling: 3,
	}

	if err != nil {
		t.Errorf("No error expected but got %v", err)
	}

	if len(stats) != 2 {
		t.Errorf("wanted %d but got %d", 1, len(stats))
	}

	if diff := deep.Equal(stats[0], expectedStat0); diff != nil {
		t.Error(diff)
	}

	if diff := deep.Equal(stats[1], expectedStat1); diff != nil {
		t.Error(diff)
	}
}

