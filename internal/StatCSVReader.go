package internal

import (
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"io"
	"os"
	"time"
)


func (date *StatDateTime) UnmarshalCSV(csv string) (err error) {
	if csv != "" {
		date.Time, err = time.ParseInLocation("02.01.2006", csv, time.Local)
		return err
	}
	return nil
}

func ReadStatCSV(file string) ([]*Stat, error) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var stats []*Stat

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		return r
	})

	if err := gocsv.UnmarshalFile(f, &stats); err != nil {
		return nil, err
	}
	return stats, nil
}