package touch

import (
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/taylormonacelli/mylime/cache"
)

func Run(project string, sentinelPath string) error {
	sentinelPath = cache.GetSentinelPath(project, sentinelPath)

	err := os.MkdirAll(filepath.Dir(sentinelPath), os.ModePerm)
	if err != nil {
		return err
	}

	_, err = os.Stat(sentinelPath)
	if os.IsNotExist(err) {
		_, err = os.Create(sentinelPath)
		if err != nil {
			return err
		}
	}

	currentTime := time.Now().Local()
	err = os.Chtimes(sentinelPath, currentTime, currentTime)
	if err != nil {
		return err
	}

	slog.Debug("Sentinel file touched successfully", "file", sentinelPath)
	return nil
}
