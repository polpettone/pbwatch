package cmd

type Stat struct {
	Date             DateTime `csv:"Date"`
	RunningDistance  float64  `csv:"RunningDistance"`
	PressUpCount     int      `csv:"PressUpCount"`
	Weight           float64  `csv:"Weight"`
	BodyFeeling      int      `csv:"BodyFeeling"`
	EmotionalFeeling int      `csv:"EmotionalFeeling"`
}



