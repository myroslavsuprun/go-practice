package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, err := time.Parse("1/02/2006 15:04:05", date)

	if err != nil {
		panic(err)
	}

	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	timeNow := time.Now()
	timeGiven, err := time.Parse("January 2, 2006 15:04:05", date)

	if err != nil {
		panic(err)
	}

	return timeNow.After(timeGiven)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	timeGiven, err := time.Parse("Monday, January 2, 2006 15:04:05", date)

	if err != nil {
		panic(err)
	}

	appointmentHour := timeGiven.Hour()
	if appointmentHour >= 12 && appointmentHour <= 18 {
		return true
	}

	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	timeGiven, err := time.Parse("1/2/2006 15:04:05", date)

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("You have an appointment on %s", timeGiven.Format("Monday, January 2, 2006, at 15:04."))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	anniversary, err := time.Parse("2006-01-02", fmt.Sprintf("%d-09-15", time.Now().Year()))

	if err != nil {
		panic(err)
	}

	return anniversary
}
