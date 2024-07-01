package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
// Schedule("7/25/2019 13:45:00")
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"

	parsed, _ := time.Parse(layout, date)

	return parsed
}

// HasPassed returns whether a date has passed.
// HasPassed("July 25, 2019 13:45:00")
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"

	parsed, _ := time.Parse(layout, date)

	return parsed.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
// IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00")
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"

	parsed, _ := time.Parse(layout, date)

	return parsed.Hour() >= 12 && parsed.Hour() < 18
}

// Description returns a formatted string of the appointment time.
// Description("7/25/2019 13:45:00")
func Description(date string) string {
	layout := "1/2/2006 15:04:05"

	parsed, _ := time.Parse(layout, date)

	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", parsed.Weekday(), parsed.Month(), parsed.Day(), parsed.Year(), parsed.Hour(), parsed.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(2024, time.September, 15, 0, 0, 0, 0, time.UTC)
}
