package cmd

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/taylormonacelli/mylime/touch"
)

var touchCmd = &cobra.Command{
	Use:   "touch PROJECT",
	Short: "Update the timestamp of the sentinel file for a project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			slog.Error("Invalid number of arguments")
			os.Exit(1)
		}

		project := args[0]

		if err := touch.Run(project, sentinelPath); err != nil {
			slog.Error("Failed to touch sentinel file", "error", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(touchCmd)
}
