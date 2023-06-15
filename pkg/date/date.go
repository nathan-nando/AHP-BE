package date

import "time"

func DateTodayLocal() *time.Time {
	now := time.Now().UTC().Add(time.Hour * 7)
	return &now
}
