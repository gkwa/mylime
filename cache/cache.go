package cache

import (
	"log/slog"
	"path/filepath"

	"github.com/kirsle/configdir"
)

func GetSentinelPath(project string, sentinelPath string) string {
	if sentinelPath != "" {
		slog.Debug("Using provided sentinel path", "path", sentinelPath)
		return sentinelPath
	}

	configPath := configdir.LocalCache("mylime")
	err := configdir.MakePath(configPath)
	if err != nil {
		panic(err)
	}

	sentinelPath = filepath.Join(configPath, project, "last_update")
	slog.Debug("Using default sentinel path", "path", sentinelPath)

	return sentinelPath
}
