package booking

import (
    "fmt"
    "time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layouts := []string{
        "Monday, January 2, 2006 15:04:05",
        "January 2, 2006 15:04:05", 
        "1/2/2006 15:04:05",       
    }
    for _, layout := range layouts {
        if t, err := time.Parse(layout, date); err == nil {
            return t
        }
    }
    panic("unable to parse date: " + date)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	return Schedule(date).Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    hour := Schedule(date).Hour()
    return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    t := Schedule(date)
    return fmt.Sprintf("You have an appointment on %s, at %s.",
        t.Format("Monday, January 2, 2006"),
        t.Format("15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}