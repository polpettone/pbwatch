package cmd

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

func (stat Stat) niceString() string {
	return fmt.Sprintf("%s \t %f \t %d \t %f",
		stat.Date.Format("02.01.2006"), stat.RunningDistance, stat.PressUpCount, stat.Weight)
}
