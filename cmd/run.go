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


	dummyStat := &Stat{
		Date:   DateTime{simpleDate(2020, 10, 29)},
		RunningDistance: 4.5,
		PressUpCount: 4,
		Weight: 75.5,
		BodyFeeling: 3,
		EmotionalFeeling: 4,
	}

	yamlDummyStat, err := yaml.Marshal(dummyStat)
	result, err := CaptureInputFromEditor(string(yamlDummyStat))

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
