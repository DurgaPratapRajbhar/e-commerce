package utils

import (
	"fmt"
	"time"
)

const (
	// DateFormat is the standard date format
	DateFormat = "2006-01-02"
	// TimeFormat is the standard time format
	TimeFormat = "15:04:05"
	// ISO8601Format is the ISO 8601 datetime format
	ISO8601Format = "2006-01-02T15:04:05Z07:00"
	// RFC3339Format is the RFC3339 datetime format
	RFC3339Format = time.RFC3339
)

// GetCurrentTime returns the current time in UTC
func GetCurrentTime() time.Time {
	return time.Now().UTC()
}

// GetCurrentTimeFormatted returns the current time formatted as standard format
func GetCurrentTimeFormatted() string {
	return GetCurrentTime().Format("2006-01-02 15:04:05")
}

// ParseDateTime parses a datetime string using the standard format
func ParseDateTime(datetimeStr string) (time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", datetimeStr)
	if err != nil {
		// Try alternative formats
		t, err = time.Parse(ISO8601Format, datetimeStr)
		if err != nil {
			t, err = time.Parse(RFC3339Format, datetimeStr)
			if err != nil {
				return time.Time{}, fmt.Errorf("unable to parse datetime: %v", err)
			}
		}
	}
	return t.UTC(), nil
}

// FormatDateTime formats a time value using the standard format
func FormatDateTime(t time.Time) string {
	return t.UTC().Format("2006-01-02 15:04:05")
}

// ParseDate parses a date string
func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse(DateFormat, dateStr)
}

// FormatDate formats a time value as date only
func FormatDate(t time.Time) string {
	return t.UTC().Format(DateFormat)
}

// IsSameDay checks if two times are on the same day
func IsSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// AddDays adds days to a time
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// BeginningOfDay returns the beginning of the day for a given time
func BeginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location()).UTC()
}

// EndOfDay returns the end of the day for a given time
func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location()).UTC()
}