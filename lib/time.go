package lib

import "time"

var TimeNow = func() time.Time {
	return time.Now()
}
