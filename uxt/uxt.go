package uxt

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ConvertUnixTimeToFormattedString converts a Unix timestamp to a formatted string.
func ConvertUnixTimeToFormattedString(unixTime int64, format string) (string, error) {
	t := time.Unix(unixTime, 0)
	var strDate string

	if strings.ToLower(format) == "rfc3339" || format == "3339" {
		strDate = t.Format(time.RFC3339)
	} else {
		strDate = t.Format(time.UnixDate)
	}

	return strDate, nil
}

// AdjustCurrentUnixTime adjusts the current Unix time by the specified number of seconds.
func AdjustCurrentUnixTime(adjustment int) (string, error) {
	now := time.Now().Unix()
	finalTime := now + int64(adjustment)
	return strconv.FormatInt(finalTime, 10), nil
}

// GetCurrentUnixTime returns the current Unix time.
func GetCurrentUnixTime() string {
	return fmt.Sprintf("%v", time.Now().Unix())
}

// GetTimeSince returns the time elapsed since the given Unix timestamp.
func GetTimeSince(unixTime int64) string {
	t := time.Unix(unixTime, 0)
	duration := time.Since(t)

	days := int(duration.Hours()) / 24
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	if days > 0 {
		return fmt.Sprintf("%d days, %d hours, %d minutes, %d seconds", days, hours, minutes, seconds)
	}
	return fmt.Sprintf("%d hours, %d minutes, %d seconds", hours, minutes, seconds)
}
