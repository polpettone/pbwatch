package cmd

import (
	"github.com/polpettone/pbwatch/pkg"
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
		Date: StatDateTime{SimpleDate(now.Year(), int(now.Month()), now.Day())},
	}

	yamlStat, err := yaml.Marshal(stat)
	result, err := pkg.CaptureInputFromEditor(string(yamlStat))

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

func SimpleDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 00, 00, 00, 00, time.Local)
}

func init() {
	logging := NewLogging()
	app := NewApplication(logging)
	addCmd := app.NewAddCmd()
	rootCmd.AddCommand(addCmd)
}
