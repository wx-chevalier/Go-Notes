package timer

import "time"

func GetNowTime() time.Time {
	return time.Now()
}

func GetCalculateTime(currentTimer time.Time, d string) (time.Timer, error) {
	duration, err := time.ParseDuration(d)

	if err != nil {
		return time.Timer{}, err
	}
}
