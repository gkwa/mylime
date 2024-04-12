package cmd

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/taylormonacelli/mylime/newerthan"
)

var newerThanCmd = &cobra.Command{
	Use:   "newerthan DURATION PROJECT",
	Short: "Check if a project is newer than a specified duration",
	Long: `Check if a project is newer than a specified duration.

Duration format:
 y - years
 M - months (30 days)
 w - weeks
 d - days
 h - hours
 m - minutes
 s - seconds
 ms - milliseconds

Fractional units are also supported, for example: 2.5y, 1.75M, 3.14d

Examples:
 mylime newerthan 1y myproject
 mylime newerthan 6M myproject
 mylime newerthan 30d myproject
 mylime newerthan 12h myproject
 mylime newerthan 30m myproject
 mylime newerthan 45s myproject
 mylime newerthan 2.5y myproject
 mylime newerthan 1.75M myproject
 mylime newerthan 3.14d myproject`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			slog.Error("Invalid number of arguments")
			os.Exit(1)
		}

		duration := args[0]
		project := args[1]

		if err := newerthan.Run(duration, project, sentinelPath); err != nil {
			slog.Error("Failed to run newerthan command", "error", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(newerThanCmd)
}
