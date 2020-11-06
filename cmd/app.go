package cmd

import "github.com/polpettone/pbwatch/internal"

type Application struct {
	Logging *internal.Logging
	DBPort string
	DBUser string
	DBPassword string
	DBName string
}

func NewApplication(logging *internal.Logging) *Application {
	return &Application{
		Logging: logging,
		DBPort: ":5432",
		DBUser: "polpettone",
		DBPassword :"polpettone",
		DBName : "pbwatch",
	}
}

