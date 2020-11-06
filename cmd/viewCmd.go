package cmd

import (
	"github.com/polpettone/pbwatch/internal"
	"github.com/spf13/cobra"
	"sort"
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
	repo := internal.NewRepo(app.Logging,
		app.DBPort,
		app.DBUser,
		app.DBPassword,
		app.DBName)

	stats, err := repo.FindAllStats()

	if err != nil {
		app.Logging.ErrorLog.Printf("%v", err)
		return
	}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Date.After(stats[j].Date.Time)
	})

	for _,s := range stats {
		app.Logging.Stdout.Printf("%s", s.NiceString())
	}
}

func init() {
	logging := internal.NewLogging(false)
	app := NewApplication(logging)
	viewCmd := app.NewViewCmd()
	rootCmd.AddCommand(viewCmd)
}
