package cmd

import (
	"github.com/polpettone/pbwatch/internal"
	"github.com/spf13/cobra"
)

func (app *Application) NewImportCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "import",
		Short: "",
		Long:  ``,

		Run: func(command *cobra.Command, args []string) {
			app.handleImportCommand(command)
		},
	}
}

func (app *Application) handleImportCommand(cobraCommand *cobra.Command) {
	initial, _ := cobraCommand.Flags().GetBool("initial")

	app.Logging.Stdout.Printf("import")

	repo := internal.NewRepo(app.Logging,
		app.DBPort,
		app.DBUser,
		app.DBPassword,
		app.DBName)

	if initial {
		app.Logging.Stdout.Printf("Create Schema")
		err := repo.CreateSchema()
		if err != nil {
			app.Logging.Stdout.Printf("%v", err)
		}
	}

	pathToCSV := "/home/icke/.pbwatch/stats.csv"
	stats, err := internal.ReadStatCSV(pathToCSV)
	if err != nil {
		app.Logging.Stdout.Printf("%v", err)
	}
	for _, s := range stats {
		err = repo.SaveStat(s)
		if err != nil {
			app.Logging.Stdout.Printf("%v", err)
		}
	}
}

func init() {
	logging := internal.NewLogging(false)
	app := NewApplication(logging)
	importCmd := app.NewImportCmd()

	importCmd.Flags().BoolP(
		"initial",
		"i",
		false,
		"Indicates an initial csv import and will try to create the db schema",
	)

	rootCmd.AddCommand(importCmd)
}
