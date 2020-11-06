package internal

import (
	"fmt"
	"time"
)

type StatDateTime struct {
	time.Time
}

type Stat struct {
	ID               int64
	Date             StatDateTime `csv:"Date" pg:",unique"`
	RunningDistance  float64      `csv:"RunningDistance"`
	PressUpCount     int          `csv:"PressUpCount"`
	Weight           float64      `csv:"Weight"`
	BodyFeeling      int          `csv:"BodyFeeling"`
	EmotionalFeeling int          `csv:"EmotionalFeeling"`
}

func (stat Stat) NiceString() string {
	return fmt.Sprintf("%s \t %f \t %d \t %f",
		stat.Date.Format("02.01.2006"), stat.RunningDistance, stat.PressUpCount, stat.Weight)
}

func SimpleDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 00, 00, 00, 00, time.Local)
}

