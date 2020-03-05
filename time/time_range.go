package time

import (
	"time"

	"github.com/e-Sirius/tools/constants"
)

// 计算时间所在类型的时间维度
func CalTimeRange(t time.Time, timeType constants.TimeType) []time.Time {
	timeRange := make([]time.Time, 0)
	switch timeType {
	case constants.TIME_TYPE__DAY:
		y, m, d := t.Date()
		start := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
		end := time.Date(y, m, d+1, 0, 0, 0, 0, time.Local).Add(-1 * time.Second)
		timeRange = append(timeRange, start, end)
	case constants.TIME_TYPE__WEEK:
		wday := int(t.Weekday()+6) % 7
		sy, sm, sd := t.AddDate(0, 0, -wday).Date()
		start := time.Date(sy, sm, sd, 0, 0, 0, 0, time.Local)
		ey, em, ed := t.AddDate(0, 0, 7-wday).Date()
		end := time.Date(ey, em, ed, 0, 0, 0, 0, time.Local).Add(-1 * time.Second)
		timeRange = append(timeRange, start, end)
	case constants.TIME_TYPE__MON:
		y, m, _ := t.Date()
		start := time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
		end := time.Date(y, m+1, 1, 0, 0, 0, 0, time.Local).Add(-1 * time.Second)
		timeRange = append(timeRange, start, end)
	case constants.TIME_TYPE__SEASON:
		return SeasonRange(t)
	case constants.TIME_TYPE__YEAR:
		y, _, _ := t.Date()
		start := time.Date(y, 1, 1, 0, 0, 0, 0, time.Local)
		end := time.Date(y+1, 1, 1, 0, 0, 0, 0, time.Local).Add(-1 * time.Second)
		timeRange = append(timeRange, start, end)
	}
	return timeRange
}

func SeasonRange(t time.Time) []time.Time {
	season := GetSeason(t)
	timeRange := make([]time.Time, 0)
	switch season {
	case 0:
		start := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.Local)
		end := time.Date(time.Now().Year(), start.Month()+3, 1, 0, 0, 0, 0, time.Local).Add(-1 * time.Second)
		timeRange = append(timeRange, start, end)
	case 1:
		start := time.Date(time.Now().Year(), 4, 1, 0, 0, 0, 0, time.Local)
		end := time.Date(time.Now().Year(), start.Month()+3, 1, 0, 0, 0, 0, time.Local).Add(-1 * time.Second)
		timeRange = append(timeRange, start, end)
	case 2:
		start := time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.Local)
		end := time.Date(time.Now().Year(), start.Month()+3, 1, 0, 0, 0, 0, time.Local).Add(-1 * time.Second)
		timeRange = append(timeRange, start, end)
	case 3:
		start := time.Date(time.Now().Year(), 10, 1, 0, 0, 0, 0, time.Local)
		end := time.Date(time.Now().Year(), start.Month()+3, 1, 0, 0, 0, 0, time.Local).Add(-1 * time.Second)
		timeRange = append(timeRange, start, end)
	}
	return timeRange
}
