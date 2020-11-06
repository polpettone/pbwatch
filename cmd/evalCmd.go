package cmd

import (
	"github.com/polpettone/pbwatch/internal"
	"github.com/spf13/cobra"
)

func (app *Application) NewEvalCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "eval",
		Short: "",
		Long:  ``,

		Run: func(command *cobra.Command, args []string) {
			app.handleEvalCommand()
		},
	}
}

type EvaluationView struct {
	Stats []internal.Stat
}

func (evaluationView EvaluationView) totalRunningDistance() float64 {
	total := 0.0
	for _, stat	:= range evaluationView.Stats {
		total += stat.RunningDistance
	}
	return total
}

func (app *Application) handleEvalCommand() {
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

	evaluationView := EvaluationView {
		Stats: stats,
	}

	app.Logging.Stdout.Printf("Total Running Distance: %f", evaluationView.totalRunningDistance())
}

func init() {
	logging := internal.NewLogging(false)
	app := NewApplication(logging)
	evalCmd := app.NewEvalCmd()
	rootCmd.AddCommand(evalCmd)
}
