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
	dateTimeFormat = "2006-01-02 15:04:00"
)

func jst() *time.Location {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	return loc
}

func Today() string {
	return FormatDate(time.Now().In(JST))
}

func ParseJSTDateInUTC(value string) (time.Time, error) {
	return parseTime(dateFormat, value, JST, time.UTC)
}

func ParseJSTDateTimeInUTC(value string) (time.Time, error) {
	return parseTime(dateTimeFormat, value, JST, time.UTC)
}

func ParseUTCDateInJST(value string) (time.Time, error) {
	return parseTime(dateFormat, value, time.UTC, JST)
}

func ParseUTCDateTimeInJST(value string) (time.Time, error) {
	return parseTime(dateTimeFormat, value, time.UTC, JST)
}

func ParseDate(value string) (time.Time, error) {
	return time.Parse(dateFormat, value)
}

func ParseDateTime(value string) (time.Time, error) {
	return time.Parse(dateTimeFormat, value)
}

func parseTime(layout, value string, from, to *time.Location) (time.Time, error) {
	t, err := time.ParseInLocation(layout, value, from)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time format: %w", err)
	}
	return t.In(to), nil
}

func FormatDateInJST(t time.Time) string {
	return formatTimeInJST(dateFormat, t)
}

func FormatDateTimeInJST(t time.Time) string {
	return formatTimeInJST(dateTimeFormat, t)
}

func FormatDate(t time.Time) string {
	return t.Format(dateFormat)
}

func FormatDateTime(t time.Time) string {
	return t.Format(dateTimeFormat)
}

func formatTimeInJST(layout string, t time.Time) string {
	return t.UTC().In(JST).Format(layout)
}
