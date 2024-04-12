package newerthan

import (
	"errors"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

func Run(args []string, sentinelPath string) error {
	if len(args) != 2 {
		return errors.New("invalid number of arguments")
	}

	duration, err := ParseCustomDuration(args[0])
	if err != nil {
		return err
	}

	project := args[1]

	if IsProjectNewerThan(project, duration, sentinelPath) {
		slog.Debug("project is newer than the specified duration")
		os.Exit(0)
	} else {
		slog.Debug("project is older than the specified duration")
		os.Exit(1)
	}

	return nil
}

func IsProjectNewerThan(project string, duration time.Duration, sentinelPath string) bool {
	if sentinelPath == "" {
		sentinelPath = filepath.Join("/var/cache", project, "last_update")
	}

	if _, err := os.Stat(sentinelPath); os.IsNotExist(err) {
		slog.Debug("Sentinel file does not exist", "file", sentinelPath)
		return false
	}

	fileInfo, err := os.Stat(sentinelPath)
	if err != nil {
		slog.Error("Failed to get file info", "error", err)
		return false
	}

	lastUpdateTime := fileInfo.ModTime()
	currentTime := time.Now()
	calculatedTime := currentTime.Add(-duration)

	slog.Debug("Sentinel file timestamp", "timestamp", lastUpdateTime)
	slog.Debug("Current time", "timestamp", currentTime)
	slog.Debug("Specified duration", "duration", duration)
	slog.Debug("Calculated timestamp", "timestamp", calculatedTime)

	elapsedTime := currentTime.Sub(lastUpdateTime)

	slog.Debug("Project last updated", "project", project, "lastUpdate", lastUpdateTime)
	slog.Debug("Elapsed time", "elapsedTime", elapsedTime)

	return elapsedTime < duration
}
