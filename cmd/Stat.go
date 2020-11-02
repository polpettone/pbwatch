package cmd

type Stat struct {
	ID   int64
	Date             DateTime `csv:"Date" pg:",unique"`
	RunningDistance  float64  `csv:"RunningDistance"`
	PressUpCount     int      `csv:"PressUpCount"`
	Weight           float64  `csv:"Weight"`
	BodyFeeling      int      `csv:"BodyFeeling"`
	EmotionalFeeling int      `csv:"EmotionalFeeling"`
}



