package util

import "time"

func ConvertDateToTimestamp(dateStr string) int64 {
	// Define the layout (format) for parsing
	layout := "01/02/2006" // MM/DD/YYYY

	// Parse the date string into a time.Time object
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return 0
	}

	// Return the Unix timestamp
	return parsedTime.Unix()
}
