package cmd

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/taylormonacelli/mylime/cleancache"
)

var cleanCacheCmd = &cobra.Command{
	Use:   "cleancache",
	Short: "Delete the cache file",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cleancache.Run(sentinelPath); err != nil {
			slog.Error("Failed to clean cache", "error", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCacheCmd)
}
