package newerthan

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

var durationRegex = regexp.MustCompile(`^(\d+(?:\.\d+)?)([yMdhms])$`)

func ParseCustomDuration(s string) (time.Duration, error) {
	matches := durationRegex.FindStringSubmatch(s)
	if len(matches) != 3 {
		return 0, errors.New("invalid duration format")
	}

	value, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0, errors.New("invalid duration value")
	}

	unit := matches[2]

	switch unit {
	case "y":
		return time.Duration(value * 365 * 24 * float64(time.Hour)), nil
	case "M":
		return time.Duration(value * 30 * 24 * float64(time.Hour)), nil
	case "d":
		return time.Duration(value * 24 * float64(time.Hour)), nil
	case "h":
		return time.Duration(value * float64(time.Hour)), nil
	case "m":
		return time.Duration(value * float64(time.Minute)), nil
	case "s":
		return time.Duration(value * float64(time.Second)), nil
	default:
		return 0, errors.New("invalid duration unit")
	}
}
