package newerthan

import (
	"log/slog"
	"os"
	"time"

	"github.com/hako/durafmt"
	"github.com/taylormonacelli/mylime/cache"
)

func Run(durationStr string, project string, sentinelPath string) error {
	duration, err := durafmt.ParseString(durationStr)
	if err != nil {
		return err
	}

	if IsProjectNewerThan(project, duration.Duration(), sentinelPath) {
		slog.Debug("project is newer than the specified duration")
		os.Exit(0)
	} else {
		slog.Debug("project is older than the specified duration")
		os.Exit(1)
	}

	return nil
}

func IsProjectNewerThan(project string, duration time.Duration, sentinelPath string) bool {
	sentinelPath = cache.GetSentinelPath(project, sentinelPath)

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
