package cmd

import (
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"time"
)

func (app *Application) NewAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "",
		Long:  ``,

		Run: func(command *cobra.Command, args []string) {
			app.handleAddCommand()
		},
	}
}

func (app *Application) handleAddCommand() {
	repo := NewRepo(app.Logging,
		app.DBPort,
		app.DBUser,
		app.DBPassword,
		app.DBName)

	now := time.Now()
	stat := &Stat{
		Date: DateTime{simpleDate(now.Year(), int(now.Month()), now.Day())},
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

}

func init() {
	logging := NewLogging()
	app := NewApplication(logging)
	addCmd := app.NewAddCmd()
	rootCmd.AddCommand(addCmd)
}
