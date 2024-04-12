package cleancache

import (
	"log/slog"
	"os"
	"path/filepath"
)

func Run(sentinelPath string) error {
	if sentinelPath == "" {
		return nil
	}

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
