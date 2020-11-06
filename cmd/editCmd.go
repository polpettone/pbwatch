package cmd

import (
	"github.com/polpettone/pbwatch/internal"
	"github.com/polpettone/pbwatch/pkg"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"time"
)

func (app *Application) NewEditCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "edit",
		Short: "",
		Long:  ``,

		Run: func(command *cobra.Command, args []string) {
			app.handleEditCommand(args)
		},
	}
}

func (app *Application) handleEditCommand(args []string) {
	repo := internal.NewRepo(app.Logging,
		app.DBPort,
		app.DBUser,
		app.DBPassword,
		app.DBName)

	date, err := time.ParseInLocation("02.01.2006", args[0], time.Local)

	if err != nil {
		app.Logging.Stdout.Printf("Could not parse date %v", err)
	}

	stat, err := repo.FindStatByDate(internal.StatDateTime{date})

	if err != nil {
		app.Logging.Stdout.Printf("Could not find stat %v", err)
	}

	yamlStat, err := yaml.Marshal(stat)
	result, err := pkg.CaptureInputFromEditor(string(yamlStat))

	if err != nil {
		app.Logging.ErrorLog.Printf("%v", err)
		return
	}

	editedStat := &internal.Stat{}

	err = yaml.Unmarshal([]byte(result), &editedStat)
	if err != nil {
		app.Logging.ErrorLog.Printf("%v", err)
		return
	}

	err = repo.SaveStat(editedStat)

	if err != nil {
		app.Logging.ErrorLog.Printf("%v", err)
		return
	}

}

func init() {
	logging := internal.NewLogging(false)
	app := NewApplication(logging)
	editCmd := app.NewEditCmd()
	rootCmd.AddCommand(editCmd)
}
