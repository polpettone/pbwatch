package cmd

import "fmt"

type Stat struct {
	ID   int64
	Date             DateTime `csv:"Date" pg:",unique"`
	RunningDistance  float64  `csv:"RunningDistance"`
	PressUpCount     int      `csv:"PressUpCount"`
	Weight           float64  `csv:"Weight"`
	BodyFeeling      int      `csv:"BodyFeeling"`
	EmotionalFeeling int      `csv:"EmotionalFeeling"`
}


func (stat Stat) niceString() string {
	return fmt.Sprintf("%v \t %f \t %f",
		 stat.Date, stat.RunningDistance, stat.Weight)
}