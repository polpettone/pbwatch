package cmd

import (
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"time"
)

func (app *Application) NewRunCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "",
		Long:  ``,

		Run: func(command *cobra.Command, args []string) {
			app.handleRunCommand()
		},
	}
}

func (app *Application) handleRunCommand() {

	repo := NewRepo(app.Logging,
		app.DBPort,
		app.DBUser,
		app.DBPassword,
		app.DBName)

	stat, err := repo.findStatByDate(DateTime{simpleDate(2020, 10, 29)})

	if err != nil {
		app.Logging.ErrorLog.Printf("%v", err)
		return
	}

	yamlStat, err := yaml.Marshal(stat)
	result, err := CaptureInputFromEditor(string(yamlStat))

	if err != nil {
		app.Logging.ErrorLog.Printf("%v", err)
		return
	}

	editedStat := &Stat{}

	err = yaml.Unmarshal([]byte(result), &editedStat)
	if err != nil {
		app.Logging.ErrorLog.Printf("%v", err)
		return
	}

	err = repo.saveStat(editedStat)

	if err != nil {
		app.Logging.ErrorLog.Printf("%v", err)
		return
	}



	app.Logging.Stdout.Printf("%v", editedStat)
	app.Logging.Stdout.Printf("------------------------")
	app.Logging.Stdout.Printf(result)

}



func init() {
	logging := NewLogging()
	app := NewApplication(logging)
	runCmd := app.NewRunCmd()
	rootCmd.AddCommand(runCmd)
}

func simpleDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 00, 00, 00, 00, time.Local)
}
