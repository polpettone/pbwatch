package cmd

import (
	"github.com/spf13/cobra"
)

func (app *Application) NewViewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "view",
		Short: "",
		Long:  ``,

		Run: func(command *cobra.Command, args []string) {
			app.handleViewCommand()
		},
	}
}

func (app *Application) handleViewCommand() {
	repo := NewRepo(app.Logging,
		app.DBPort,
		app.DBUser,
		app.DBPassword,
		app.DBName)

	stats, err := repo.findAllStats()

	if err != nil {
		app.Logging.ErrorLog.Printf("%v", err)
		return
	}

	for _,s := range stats {
		app.Logging.Stdout.Printf("%s", s.niceString())
	}
}

func init() {
	logging := NewLogging()
	app := NewApplication(logging)
	viewCmd := app.NewViewCmd()
	rootCmd.AddCommand(viewCmd)
}
