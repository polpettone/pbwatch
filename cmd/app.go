package cmd



type Application struct {
	Logging *Logging
	DBPort string
	DBUser string
	DBPassword string
	DBName string
}

func NewApplication(logging *Logging) *Application {
	return &Application{
		Logging: logging,
		DBPort: ":5432",
		DBUser: "polpettone",
		DBPassword :"pbwatch",
		DBName : "pbwatch",
	}
}

