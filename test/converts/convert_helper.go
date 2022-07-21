package converts

import "time"

func ConvIntToTime(in int64) time.Time {
	return time.Unix(in, 0)
}

func ConvTimeToInt(in time.Time) int64 {
	return in.Unix()
}
