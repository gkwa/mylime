package cache

import "path/filepath"

func GetSentinelPath(project string, sentinelPath string) string {
	if sentinelPath != "" {
		return sentinelPath
	}
	return filepath.Join("/var/cache", project, "last_update")
}
