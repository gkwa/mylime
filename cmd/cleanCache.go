package cmd

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/taylormonacelli/mylime/cleancache"
)

var cleanCacheCmd = &cobra.Command{
	Use:   "cleancache PROJECT",
	Short: "Delete the cache file for a project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			slog.Error("Invalid number of arguments")
			os.Exit(1)
		}

		project := args[0]

		if err := cleancache.Run(project, sentinelPath); err != nil {
			slog.Error("Failed to clean cache", "error", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCacheCmd)
}
