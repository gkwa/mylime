package cleancache

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/taylormonacelli/mylime/cache"
)

func Run(project string, sentinelPath string) error {
	sentinelPath = cache.GetSentinelPath(project, sentinelPath)

	if _, err := os.Stat(sentinelPath); os.IsNotExist(err) {
		slog.Debug("Sentinel file does not exist", "file", sentinelPath)
		return nil
	}

	err := os.Remove(sentinelPath)
	if err != nil {
		return err
	}

	if _, err := os.Stat(sentinelPath); os.IsNotExist(err) {
		slog.Debug("Sentinel file deleted successfully", "file", sentinelPath)
		return nil
	} else {
		return filepath.ErrBadPattern
	}
}
