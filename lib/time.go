package lib

import (
	"fmt"
	"time"
)

var TimeNow = func() time.Time {
	return time.Now().UTC()
}

var JST *time.Location = jst()

var (
	dateFormat     = "2006-01-02"
	dateTimeFormat = "2006-01-02 15:04"
)

func jst() *time.Location {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	return loc
}

func ParseJSTDateInUTC(value string) (time.Time, error) {
	return parseJSTTimeInUTC(dateFormat, value)
}

func ParseJSTDateTimeInUTC(value string) (time.Time, error) {
	return parseJSTTimeInUTC(dateTimeFormat, value)
}

func parseJSTTimeInUTC(layout, value string) (time.Time, error) {
	t, err := time.ParseInLocation(layout, value, JST)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time format: %w", err)
	}
	return t.UTC(), err
}

func FormatDateInJST(t time.Time) string {
	return formatTimeInJST(dateFormat, t)
}

func FormatDateTimeInJST(t time.Time) string {
	return formatTimeInJST(dateTimeFormat, t)
}

func formatTimeInJST(layout string, t time.Time) string {
	return t.UTC().In(JST).Format(layout)
}
