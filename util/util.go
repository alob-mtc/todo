package util

import (
	"time"
)

func TimeToString(val time.Time) string {
	formatted := val.Format(time.RFC3339)
	return formatted
}

func CompletedToWords(completed bool) string {
	if completed {
		return "Task Completed"
	}
	return "Task Pending"
}
